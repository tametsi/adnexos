package plugin

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/types"
)

// remove expired invitations
func (p *plugin) invitesRemove() {
	p.app.Dao().DB().
		NewQuery("DELETE FROM invites WHERE expires < {:now}").
		Bind(dbx.Params{
			"now": types.NowDateTime(),
		}).Execute()
}
