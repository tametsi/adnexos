package service

import (
	"slices"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func (p *plugin) onExpensesBeforeCreate(e *core.RecordCreateEvent) error {
	return p.validateExpense(e.Record)
}

// fires on groups update request to check if no new user have been added
func (p *plugin) onExpensesBeforeUpdate(e *core.RecordUpdateEvent) error {
	return p.validateExpense(e.Record)
}

func (p *plugin) validateExpense(r *models.Record) error {
	group, err := p.app.Dao().FindRecordById("groups", r.GetString("group"))
	if err != nil {
		return apis.NewBadRequestError("Failed to load group.", err)
	}

	groupMembers := append(group.GetStringSlice("members"), group.GetString("owner"))
	expenseMembers := r.GetStringSlice("members")

	for _, x := range expenseMembers {
		if !slices.Contains(groupMembers, x) {
			return apis.NewBadRequestError("Members must be part of the group.", nil)
		}
	}

	return nil
}
