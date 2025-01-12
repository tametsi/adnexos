package plugin

import (
	"net/url"
	"slices"
	"strings"

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

		return se.Next()
	})

	// record operations
	app.OnRecordCreateRequest("expenses").BindFunc(p.onExpensesBeforeCreate)
	app.OnRecordUpdateRequest("expenses").BindFunc(p.onExpensesBeforeUpdate)

	app.OnRecordViewRequest("groups").BindFunc(p.onGroupsView)
	app.OnRecordsListRequest("groups").BindFunc(p.onGroupsList)
	app.OnRecordUpdateRequest("groups").BindFunc(p.onGroupsBeforeUpdate)

	app.OnRecordCreateRequest("settings").BindFunc(p.onSettingsCreate)

	// cron jobs
	app.Cron().MustAdd("remove-invites", "0 1 * * *", p.invitesRemove)

	return nil
}

type plugin struct {
	app core.App
}

// helpers

// checks whether `value` is present in a comma separated list in query param `fields`
func hasFieldValue(url *url.URL, value string) bool {
	fields := url.Query().Get("fields")
	if fields == "" {
		return false
	}
	fieldValues := strings.Split(fields, ",")

	return slices.Contains(fieldValues, value)
}
