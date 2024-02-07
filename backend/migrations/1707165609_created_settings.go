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
			"id": "vqc8fq4hcqxxxe0",
			"created": "2024-02-05 20:40:09.241Z",
			"updated": "2024-02-05 20:40:09.241Z",
			"name": "settings",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "d0iqf3wb",
					"name": "user",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "0p0brqrz",
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
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id = user.id",
			"viewRule": "@request.auth.id = user.id",
			"createRule": "@request.auth.id = user.id",
			"updateRule": "@request.auth.id = user.id",
			"deleteRule": "@request.auth.id = user.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("vqc8fq4hcqxxxe0")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
