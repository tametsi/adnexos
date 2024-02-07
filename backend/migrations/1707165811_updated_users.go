package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		// remove
		collection.Schema.RemoveField("7vxxjqye")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("id = @request.auth.id")

		// add
		del_theme := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7vxxjqye",
			"name": "theme",
			"type": "select",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"values": [
					"light",
					"dark"
				]
			}
		}`), del_theme)
		collection.Schema.AddField(del_theme)

		return dao.SaveCollection(collection)
	})
}
