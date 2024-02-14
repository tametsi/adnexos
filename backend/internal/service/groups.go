package service

import (
	"net/http"
	"slices"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func (p *plugin) groupJoinRoute(c echo.Context) error {
	id := c.PathParam("id")

	auth := apis.RequestInfo(c).AuthRecord

	invite, err := p.app.Dao().FindRecordById("invites", id)
	if err != nil {
		return apis.NewNotFoundError("Invite not found.", nil)
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
