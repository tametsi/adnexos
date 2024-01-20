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
			"id": "i543bq5kos5oc7k",
			"created": "2024-01-19 17:00:58.800Z",
			"updated": "2024-01-19 17:00:58.800Z",
			"name": "groups",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "rareix11",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "t7ryxei6",
					"name": "owner",
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
					"id": "keh9yf6q",
					"name": "members",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id = owner.id || members.id ?= @request.auth.id",
			"viewRule": "@request.auth.id = owner.id || members.id ?= @request.auth.id",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "@request.auth.id = owner.id",
			"deleteRule": "@request.auth.id = owner.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("i543bq5kos5oc7k")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
