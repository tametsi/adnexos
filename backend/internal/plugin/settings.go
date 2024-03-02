package plugin

import (
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

// fires on settings create request to check if settings for user already exist
func (p *plugin) onSettingsCreate(e *core.RecordCreateEvent) error {
	userId := e.Record.GetString("user")
	if userId == "" {
		return nil // nothing to do here, the req will fail later...
	}

	setting := models.Record{}
	err := p.app.Dao().DB().
		NewQuery("SELECT * FROM settings WHERE user = {:user}").
		Bind(dbx.Params{"user": userId}).
		One(&setting)
	if err == nil {
		// settings already exist => no creation ;)
		return apis.NewApiError(http.StatusBadRequest, "settings already exist for user", nil)
	}

	return nil
}
