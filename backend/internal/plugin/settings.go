package plugin

import (
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// fires on settings create request to check if settings for user already exist
func (p *plugin) onSettingsCreate(e *core.RecordRequestEvent) error {
	userId := e.Record.GetString("user")
	if userId == "" {
		return e.Next() // nothing to do here, the req will fail later...
	}

	_, err := p.app.FindFirstRecordByFilter("settings", "user = {:user}", dbx.Params{"user": userId})
	if err == nil {
		// settings already exist => no creation ;)
		return apis.NewApiError(http.StatusBadRequest, "Settings already exist for user.", nil)
	}

	return e.Next()
}
