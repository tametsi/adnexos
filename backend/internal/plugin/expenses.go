package plugin

import (
	"slices"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func (p *plugin) onExpensesBeforeCreate(e *core.RecordRequestEvent) error {
	err := p.validateExpense(e.Record)
	if err != nil {
		return err
	}

	return e.Next()
}

// fires on groups update request to check if no new user have been added
func (p *plugin) onExpensesBeforeUpdate(e *core.RecordRequestEvent) error {
	err := p.validateExpense(e.Record)
	if err != nil {
		return err
	}

	return e.Next()
}

func (p *plugin) validateExpense(r *core.Record) error {
	group, err := p.app.FindRecordById("groups", r.GetString("group"))
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
