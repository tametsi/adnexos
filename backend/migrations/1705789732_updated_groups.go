package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("i543bq5kos5oc7k")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id = owner.id || (members.id ?= @request.auth.id && members:length > 0)")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("i543bq5kos5oc7k")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id = owner.id || members.id ?= @request.auth.id")

		return dao.SaveCollection(collection)
	})
}
