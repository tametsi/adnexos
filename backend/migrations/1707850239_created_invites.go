package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "d4q6aeyn17sd9z9",
			"created": "2024-02-13 18:50:39.442Z",
			"updated": "2024-02-13 18:50:39.442Z",
			"name": "invites",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "zksplpg0",
					"name": "group",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "i543bq5kos5oc7k",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "b2thvscc",
					"name": "creator",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "u1exrssm",
					"name": "expires",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				}
			],
			"indexes": [],
			"listRule": "group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id",
			"viewRule": "group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id",
			"createRule": "@request.auth.id != \"\" && creator.id = @request.auth.id && (group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id) && expires > @now",
			"updateRule": null,
			"deleteRule": "creator.id = @request.auth.id || group.owner.id = @request.auth.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("d4q6aeyn17sd9z9")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
