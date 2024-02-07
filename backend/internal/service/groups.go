package service

import (
	"slices"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

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
