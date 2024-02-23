package plugin

import (
	"slices"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// Register registers the custom business logic to the provided app instance.
func Register(app core.App) error {
	p := &plugin{app: app}

	// routing
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/collections/groups/join/:id", p.groupJoinRoute,
			apis.ActivityLogger(p.app), apis.RequireRecordAuth())

		return nil
	})

	// record operations
	app.OnRecordBeforeCreateRequest("expenses").Add(p.onExpensesBeforeCreate)
	app.OnRecordBeforeUpdateRequest("expenses").Add(p.onExpensesBeforeUpdate)

	app.OnRecordViewRequest("groups").Add(p.onGroupsView)
	app.OnRecordsListRequest("groups").Add(p.onGroupsList)
	app.OnRecordBeforeUpdateRequest("groups").Add(p.onGroupsBeforeUpdate)

	app.OnRecordBeforeCreateRequest("settings").Add(p.onSettingsCreate)

	return nil
}

type plugin struct {
	app core.App
}

// helpers

// checks whether `value` is present in a comma separated list in query param `fields`
func hasFieldValue(ctx echo.Context, value string) bool {
	fields := ctx.QueryParams().Get("fields")
	if fields == "" {
		return false
	}
	fieldValues := strings.Split(fields, ",")

	return slices.Contains(fieldValues, value)
}
