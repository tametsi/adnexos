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

		collection, err := dao.FindCollectionByNameOrId("g5rdaq8oufhn82z")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("(group.members.id ?= @request.auth.id && isPrivate = false) || (isPrivate = true && (members.id ?= @request.auth.id || source.id = @request.auth.id)) || group.owner.id = @request.auth.id")

		collection.ViewRule = types.Pointer("(group.members.id ?= @request.auth.id && isPrivate = false) || (isPrivate = true && (members.id ?= @request.auth.id || source.id = @request.auth.id)) || group.owner.id = @request.auth.id")

		// add
		new_isPrivate := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "89nv20ga",
			"name": "isPrivate",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_isPrivate); err != nil {
			return err
		}
		collection.Schema.AddField(new_isPrivate)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("g5rdaq8oufhn82z")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id")

		collection.ViewRule = types.Pointer("group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id")

		// remove
		collection.Schema.RemoveField("89nv20ga")

		return dao.SaveCollection(collection)
	})
}
