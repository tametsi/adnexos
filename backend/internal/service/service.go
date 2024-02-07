package service

import (
	"github.com/pocketbase/pocketbase/core"
)

// Register registers the custom business logic to the provided app instance.
func Register(app core.App) error {
	p := &plugin{app: app}

	// create settings request
	app.OnRecordBeforeCreateRequest("settings").Add(p.onSettingsCreateRequest)

	return nil
}

type plugin struct {
	app core.App
}
