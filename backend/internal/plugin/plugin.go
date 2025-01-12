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

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// routing
		se.Router.POST("/api/collections/groups/join/{id}", p.groupJoinRoute).Bind(apis.RequireAuth())
		se.Router.POST("/api/collections/groups/{id}/leave", p.groupLeaveRoute).Bind(apis.RequireAuth())
		se.Router.POST("/api/collections/groups/{id}/settle", p.groupSettleRoute).Bind(apis.RequireAuth())

		return nil
	})

	// record operations
	app.OnRecordBeforeCreateRequest("expenses").Add(p.onExpensesBeforeCreate)
	app.OnRecordBeforeUpdateRequest("expenses").Add(p.onExpensesBeforeUpdate)

	app.OnRecordViewRequest("groups").Add(p.onGroupsView)
	app.OnRecordsListRequest("groups").Add(p.onGroupsList)
	app.OnRecordBeforeUpdateRequest("groups").Add(p.onGroupsBeforeUpdate)

	app.OnRecordBeforeCreateRequest("settings").Add(p.onSettingsCreate)

	// cron jobs
	app.Cron().MustAdd("remove-invites", "0 1 * * *", p.invitesRemove)

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
