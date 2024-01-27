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
			"id": "g5rdaq8oufhn82z",
			"created": "2024-01-22 20:35:39.996Z",
			"updated": "2024-01-22 20:35:39.996Z",
			"name": "expenses",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "tcnrr1dv",
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
					"id": "zn11dfuj",
					"name": "amount",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": true
					}
				},
				{
					"system": false,
					"id": "1hmmrfct",
					"name": "title",
					"type": "text",
					"required": false,
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
					"id": "rh3dgscj",
					"name": "isSettled",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "lvfathxy",
					"name": "members",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "u69eyawg",
					"name": "source",
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
				}
			],
			"indexes": [],
			"listRule": "group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id",
			"viewRule": "group.members.id ?= @request.auth.id || group.owner.id = @request.auth.id",
			"createRule": "@request.auth.id != \"\"",
			"updateRule": "@request.auth.id = source.id",
			"deleteRule": "@request.auth.id = source.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("g5rdaq8oufhn82z")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
