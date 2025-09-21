package plugin

import (
	"net/http"
	"slices"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tametsi/adnexos/internal/service"
)

func (p *plugin) groupJoinRoute(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	info, err := e.RequestInfo()
	if err != nil {
		return apis.NewInternalServerError("Failed to get Request Info.", err)
	}
	auth := info.Auth

	invite, err := p.app.FindRecordById("invites", id)
	if err != nil {
		return apis.NewNotFoundError("Invite not found.", nil)
	}

	if expires := invite.GetDateTime("expires").Time(); time.Now().After(expires) {
		return apis.NewApiError(http.StatusBadRequest, "Invite expired.", nil)
	}

	group, err := p.app.FindRecordById("groups", invite.GetString("group"))
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Something went wrong.", nil)
	}

	members := group.GetStringSlice("members")
	if slices.Contains(members, auth.Id) || group.GetString("owner") == auth.Id {
		return apis.NewBadRequestError("Already joined, du Witzbold", nil)
	}
	group.Set("members", append(members, auth.Id))

	err = p.app.Save(group)
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Saving seems to be a hard task...", nil)
	}

	return e.NoContent(http.StatusNoContent)
}

func (p *plugin) groupLeaveRoute(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	group, err := p.app.FindRecordById("groups", id)
	if err != nil {
		return apis.NewNotFoundError("Group not found.", err)
	}

	info, err := e.RequestInfo()
	if err != nil {
		return apis.NewInternalServerError("Failed to get Request Info.", err)
	}
	auth := info.Auth

	if group.GetString("owner") == auth.Id {
		return apis.NewBadRequestError("Set a new owner first.", nil)
	}

	members := group.GetStringSlice("members")
	if !slices.Contains(members, auth.Id) {
		return apis.NewNotFoundError("Group not found.", nil)
	}

	members = slices.DeleteFunc(members, func(x string) bool { return x == auth.Id })

	group.Set("members", members)
	err = p.app.Save(group)
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Failed to update group.", err)
	}

	return e.NoContent(http.StatusNoContent)
}

func (p *plugin) groupSettleRoute(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	info, err := e.RequestInfo()
	if err != nil {
		return apis.NewInternalServerError("Failed to get Request Info.", err)
	}
	auth := info.Auth

	note, _ := info.Body["note"].(string)
	if len(note) > 50 {
		return apis.NewBadRequestError("Invalid Note.", err)
	}

	group, err := p.app.FindRecordById("groups", id)
	if err != nil {
		return apis.NewNotFoundError("Group not found.", err)
	}

	members := group.GetStringSlice("members")
	if !slices.Contains(members, auth.Id) && group.GetString("owner") != auth.Id {
		return apis.NewNotFoundError("Group not found.", err)
	}

	// gather expenses
	recordExpenses, err := p.app.FindRecordsByFilter("expenses", "group = {:group} && isSettled = false", "", -1, 0,
		dbx.Params{"group": id})
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Failed to query expenses.", err)
	}

	// map expenses
	expenses := []service.Expense{}
	for _, e := range recordExpenses {
		expenses = append(expenses, service.Expense{
			Amount:  e.GetInt("amount"),
			Members: e.GetStringSlice("members"),
			Source:  e.GetString("source"),
		})
	}

	settle := service.NewSettle(append(members, group.GetString("owner")), expenses)
	payments := settle.CreatePayments()

	collection, err := p.app.FindCollectionByNameOrId("payments")
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Find payments collection.", err)
	}

	paymentRecords := []core.Record{}
	for _, payment := range payments {
		record := core.NewRecord(collection)
		record.Set("to", payment.To)
		record.Set("from", payment.From)
		record.Set("amount", payment.Amount)
		record.Set("note", note)

		err = p.app.Save(record)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError,
				"Creating at least one payment failed. Check already created payments!", err)
		}

		paymentRecords = append(paymentRecords, *record)
	}

	// mark expenses settled
	for _, e := range recordExpenses {
		e.Set("isSettled", true)

		err = p.app.Save(e)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Update expenses.", err)
		}
	}

	return e.JSON(http.StatusOK, paymentRecords)
}

// fires on groups update request to check if no new user have been added
func (p *plugin) onGroupsBeforeUpdate(e *core.RecordRequestEvent) error {
	oldRecord, err := p.app.FindRecordById(e.Collection.Id, e.Record.Id)
	if err != nil {
		return err
	}

	oldOwner := oldRecord.GetString("owner")
	newOwner := e.Record.GetString("owner")
	oldMembers := oldRecord.GetStringSlice("members")
	newMembers := e.Record.GetStringSlice("members")

	info, err := e.RequestInfo()
	if err != nil {
		return apis.NewInternalServerError("Failed to get Request Info.", err)
	}
	isAdmin := info.HasSuperuserAuth()

	for _, v := range newMembers {
		if newOwner == v {
			return apis.NewBadRequestError("Owner is not allowed to be a member.", nil)
		}

		if !slices.Contains(oldMembers, v) && v != oldOwner && !isAdmin {
			// member has been added
			return apis.NewBadRequestError("Adding members is not allowed.", nil)
		}
	}

	return e.Next()
}

func (p *plugin) onGroupsView(e *core.RecordRequestEvent) error {
	hasBalance := hasFieldValue(e.Request.URL, "balance")
	hasBalanceForMembers := hasFieldValue(e.Request.URL, "members.balance")
	hasCosts := hasFieldValue(e.Request.URL, "costs")
	if !hasBalance && !hasCosts && !hasBalanceForMembers {
		return e.Next()
	}

	info, err := e.RequestInfo()
	if err != nil {
		return apis.NewInternalServerError("Failed to get Request Info.", err)
	}
	auth := info.Auth
	e.Record.WithCustomData(true)

	if hasBalance {
		balance, err := p.calculateBalanceForGroup(e.Record.Id, auth.Id)
		if err != nil {
			return err
		}

		e.Record.Set("balance", balance)
	}

	if hasCosts {
		costs, err := p.calculateCostsForGroup(e.Record.Id, auth.Id)
		if err != nil {
			return err
		}

		e.Record.Set("costs", costs)
	}

	if hasBalanceForMembers {
		// WARNING: massivly botched code ahead!
		// However, we need to modify expanded data which is not possible without even more chaos...
		// FIXME one day and come up with something better (:

		memberIDs := e.Record.GetStringSlice("members")
		membersBalance := map[string]int{}

		for _, mID := range append(memberIDs, e.Record.GetString("owner")) {
			balance, err := p.calculateBalanceForGroup(e.Record.Id, mID)
			if err != nil {
				return err
			}

			membersBalance[mID] = balance
		}

		e.Record.Set("membersBalance", membersBalance)
	}

	return e.Next()
}

func (p *plugin) onGroupsList(e *core.RecordsListRequestEvent) error {
	if !hasFieldValue(e.Request.URL, "balance") {
		return e.Next()
	}

	info, err := e.RequestInfo()
	if err != nil {
		return apis.NewInternalServerError("Failed to get Request Info.", err)
	}
	auth := info.Auth

	for _, record := range e.Records {
		balance, err := p.calculateBalanceForGroup(record.Id, auth.Id)
		if err != nil {
			return err
		}

		record.WithCustomData(true)
		record.Set("balance", balance)
	}

	return e.Next()
}

func (p *plugin) calculateBalanceForGroup(groupId, authId string) (int, error) {
	balance := 0

	var balancePlus int
	// + balance for all expenses with source == me
	err := p.app.DB().
		NewQuery("SELECT COALESCE(SUM(amount), 0) FROM expenses WHERE `group` = {:group} AND source = {:auth} AND isSettled = false").
		Bind(dbx.Params{
			"group": groupId,
			"auth":  authId,
		}).Row(&balancePlus)
	if err != nil {
		return 0, apis.NewApiError(http.StatusInternalServerError, "Failed to query DB to calculate balance (1).", err)
	}
	balance += balancePlus

	// - balance for all expenses with: members contains me
	expenses, err := p.app.FindRecordsByFilter("expenses", "group = {:group} && members.id ?= {:auth} && isSettled = false", "", -1, 0,
		dbx.Params{"group": groupId, "auth": authId})
	if err != nil {
		return 0, apis.NewApiError(http.StatusInternalServerError, "Failed to query DB to calculate balance (2).", err)
	}

	for _, expense := range expenses {
		balance -= expense.GetInt("amount") / len(expense.GetStringSlice("members"))
	}

	return balance, nil
}

func (p *plugin) calculateCostsForGroup(groupId, authId string) (int, error) {
	costs := 0

	expenses, err := p.app.FindRecordsByFilter("expenses", "group = {:group} && members.id ?= {:auth} && isSettled = false", "", -1, 0,
		dbx.Params{"group": groupId, "auth": authId})
	if err != nil {
		return 0, apis.NewApiError(http.StatusInternalServerError, "Failed to query DB to calculate costs.", err)
	}

	for _, expense := range expenses {
		costs += expense.GetInt("amount") / len(expense.GetStringSlice("members"))
	}

	return costs, nil
}
