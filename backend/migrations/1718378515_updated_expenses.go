package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("g5rdaq8oufhn82z")
		if err != nil {
			return err
		}

		collection.UpdateRule = types.Pointer("(@request.auth.id = group.owner.id ||  @request.auth.id = source.id) && (group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id)")
		collection.DeleteRule = types.Pointer("@request.auth.id = source.id || @request.auth.id = group.owner.id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("g5rdaq8oufhn82z")
		if err != nil {
			return err
		}

		collection.UpdateRule = types.Pointer("@request.auth.id = source.id && (group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id)")
		collection.DeleteRule = types.Pointer("@request.auth.id = source.id")

		return dao.SaveCollection(collection)
	})
}
