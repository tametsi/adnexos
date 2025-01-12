package plugin

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

// remove expired invitations
func (p *plugin) invitesRemove() {
	_, err := p.app.DB().
		NewQuery("DELETE FROM invites WHERE expires < {:now}").
		Bind(dbx.Params{
			"now": types.NowDateTime(),
		}).Execute()
	if err != nil {
		p.app.Logger().Error("Failed to remove expired invites", "error", err)
	}
}
