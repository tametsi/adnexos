package plugin

import (
	"net/http"
	"slices"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models/schema"
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
	if slices.Contains(members, auth.Id) {
		return apis.NewBadRequestError("Already joined, du Witzbold", nil)
	}
	group.Set("members", append(members, auth.Id))

	err = p.app.Dao().SaveRecord(group)
	if err != nil {
		return apis.NewApiError(http.StatusInternalServerError, "Saving seems to be a hard task...", nil)
	}

	return c.NoContent(http.StatusNoContent)
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

	for _, v := range newMembers {
		if newOwner == v {
			return apis.NewBadRequestError("Owner is not allowed to be a member.", nil)
		}

		if !slices.Contains(oldMembers, v) && v != oldOwner {
			// member has been added
			return apis.NewBadRequestError("Adding members is not allowed.", nil)
		}
	}

	return nil
}

func (p *plugin) onGroupsView(e *core.RecordViewEvent) error {
	if !hasFieldValue(e.HttpContext, "balance") {
		return nil
	}

	// temporarily add balance field to collection
	e.Record.Collection().Schema.AddField(&schema.SchemaField{
		Name: "balance",
		Type: "int",
	})

	auth := apis.RequestInfo(e.HttpContext).AuthRecord
	balance, err := p.calculateBalanceForGroup(e.Record.Id, auth.Id)
	if err != nil {
		return err
	}

	e.Record.Set("balance", balance)

	return nil
}

func (p *plugin) onGroupsList(e *core.RecordsListEvent) error {
	if !hasFieldValue(e.HttpContext, "balance") {
		return nil
	}

	// temporarily add balance field to collection
	e.Collection.Schema.AddField(&schema.SchemaField{
		Name: "balance",
		Type: "int",
	})

	auth := apis.RequestInfo(e.HttpContext).AuthRecord

	for _, record := range e.Records {
		balance, err := p.calculateBalanceForGroup(record.Id, auth.Id)
		if err != nil {
			return err
		}

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
