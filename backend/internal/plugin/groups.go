package plugin

import (
	"net/http"
	"slices"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/tametsi/adnexos/internal/service"
)

func (p *plugin) groupJoinRoute(c echo.Context) error {
	id := c.PathParam("id")

	auth := apis.RequestInfo(c).AuthRecord

	invite, err := p.app.Dao().FindRecordById("invites", id)
	if err != nil {
		return apis.NewNotFoundError("Invite not found.", nil)
	}

	if expires := invite.GetDateTime("expires").Time(); time.Now().After(expires) {
		return apis.NewApiError(http.StatusBadRequest, "Invite expired.", nil)
	}

	group, err := p.app.Dao().FindRecordById("groups", invite.GetString("group"))
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Something went wrong.", nil)
	}

	members := group.GetStringSlice("members")
	if slices.Contains(members, auth.Id) || group.GetString("owner") == auth.Id {
		return apis.NewBadRequestError("Already joined, du Witzbold", nil)
	}
	group.Set("members", append(members, auth.Id))

	err = p.app.Dao().SaveRecord(group)
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Saving seems to be a hard task...", nil)
	}

	return c.NoContent(http.StatusNoContent)
}

func (p *plugin) groupLeaveRoute(c echo.Context) error {
	id := c.PathParam("id")

	group, err := p.app.Dao().FindRecordById("groups", id)
	if err != nil {
		return apis.NewNotFoundError("Group not found.", err)
	}

	auth := apis.RequestInfo(c).AuthRecord

	if group.GetString("owner") == auth.Id {
		return apis.NewBadRequestError("Set a new owner first.", nil)
	}

	members := group.GetStringSlice("members")
	if !slices.Contains(members, auth.Id) {
		return apis.NewNotFoundError("Group not found.", nil)
	}

	members = slices.DeleteFunc(members, func(x string) bool { return x == auth.Id })

	group.Set("members", members)
	err = p.app.Dao().SaveRecord(group)
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Failed to update group.", err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (p *plugin) groupSettleRoute(c echo.Context) error {
	id := c.PathParam("id")

	auth := apis.RequestInfo(c).AuthRecord

	group, err := p.app.Dao().FindRecordById("groups", id)
	if err != nil {
		return apis.NewNotFoundError("Group not found.", err)
	}

	members := group.GetStringSlice("members")
	if !slices.Contains(members, auth.Id) && group.GetString("owner") != auth.Id {
		return apis.NewNotFoundError("Group not found.", err)
	}

	// gather expenses
	recordExpenses, err := p.app.Dao().FindRecordsByFilter("expenses", "group = {:group} && isSettled = false", "", -1, 0,
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

	collection, err := p.app.Dao().FindCollectionByNameOrId("payments")
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Find payments collection.", err)
	}

	paymentRecords := []models.Record{}
	for _, payment := range payments {
		record := models.NewRecord(collection)
		record.Set("to", payment.To)
		record.Set("from", payment.From)
		record.Set("amount", payment.Amount)

		err = p.app.Dao().SaveRecord(record)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError,
				"Creating at least one payment failed. Check already created payments!", err)
		}

		paymentRecords = append(paymentRecords, *record)
	}

	// mark expenses settled
	for _, e := range recordExpenses {
		e.Set("isSettled", true)

		err = p.app.Dao().SaveRecord(e)
		if err != nil {
			return apis.NewApiError(http.StatusInternalServerError, "Update expenses.", err)
		}
	}

	return c.JSON(http.StatusOK, paymentRecords)
}

// fires on groups update request to check if no new user have been added
func (p *plugin) onGroupsBeforeUpdate(e *core.RecordUpdateEvent) error {
	oldRecord, err := p.app.Dao().FindRecordById(e.Collection.Id, e.Record.Id)
	if err != nil {
		return err
	}

	oldOwner := oldRecord.GetString("owner")
	newOwner := e.Record.GetString("owner")
	oldMembers := oldRecord.GetStringSlice("members")
	newMembers := e.Record.GetStringSlice("members")

	admin, _ := e.HttpContext.Get(apis.ContextAdminKey).(*models.Admin)

	for _, v := range newMembers {
		if newOwner == v {
			return apis.NewBadRequestError("Owner is not allowed to be a member.", nil)
		}

		if !slices.Contains(oldMembers, v) && v != oldOwner && admin == nil {
			// member has been added
			return apis.NewBadRequestError("Adding members is not allowed.", nil)
		}
	}

	return nil
}

func (p *plugin) onGroupsView(e *core.RecordViewEvent) error {
	hasBalance := hasFieldValue(e.HttpContext, "balance")
	hasBalanceForMembers := hasFieldValue(e.HttpContext, "members.balance")
	hasCosts := hasFieldValue(e.HttpContext, "costs")
	if !hasBalance && !hasCosts && !hasBalanceForMembers {
		return nil
	}

	auth := apis.RequestInfo(e.HttpContext).AuthRecord
	e.Record.WithUnknownData(true)

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

		for _, mID := range memberIDs {
			balance, err := p.calculateBalanceForGroup(e.Record.Id, mID)
			if err != nil {
				return err
			}

			membersBalance[mID] = balance
		}

		e.Record.Set("membersBalance", membersBalance)
	}

	return nil
}

func (p *plugin) onGroupsList(e *core.RecordsListEvent) error {
	if !hasFieldValue(e.HttpContext, "balance") {
		return nil
	}

	auth := apis.RequestInfo(e.HttpContext).AuthRecord

	for _, record := range e.Records {
		balance, err := p.calculateBalanceForGroup(record.Id, auth.Id)
		if err != nil {
			return err
		}

		record.WithUnknownData(true)
		record.Set("balance", balance)
	}

	return nil
}

func (p *plugin) calculateBalanceForGroup(groupId, authId string) (int, error) {
	balance := 0

	var balancePlus int
	// + balance for all expenses with source == me
	err := p.app.Dao().DB().
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
	expenses, err := p.app.Dao().FindRecordsByFilter("expenses", "group = {:group} && members.id ?= {:auth} && isSettled = false", "", -1, 0,
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

	expenses, err := p.app.Dao().FindRecordsByFilter("expenses", "group = {:group} && members.id ?= {:auth} && isSettled = false", "", -1, 0,
		dbx.Params{"group": groupId, "auth": authId})
	if err != nil {
		return 0, apis.NewApiError(http.StatusInternalServerError, "Failed to query DB to calculate costs.", err)
	}

	for _, expense := range expenses {
		costs += expense.GetInt("amount") / len(expense.GetStringSlice("members"))
	}

	return costs, nil
}
