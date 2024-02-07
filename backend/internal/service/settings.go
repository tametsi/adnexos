package service

import (
	"fmt"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/hook"
)

// fires on settings create request to check if settings for user already exists
func (p *plugin) onSettingsCreate(e *core.RecordCreateEvent) error {
	userId := e.Record.GetString("user")
	if userId == "" {
		return nil // nothing to do here, the req will fail later...
	}

	var setting models.Record
	err := p.app.Dao().DB().
		NewQuery(fmt.Sprintf("SELECT * FROM settings WHERE user=\"%s\"", userId)).
		One(&setting)
	if err != nil && setting.Id == "" {
		// no setting for current user found => create one
		return nil
	}

	// just return existing settings
	e.HttpContext.JSON(http.StatusOK, setting)

	return hook.StopPropagation
}
