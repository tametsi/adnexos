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
			"id": "5bslbh0q4ky3wwx",
			"created": "2024-02-23 18:37:31.835Z",
			"updated": "2024-02-23 18:37:31.835Z",
			"name": "payments",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "yiwradbc",
					"name": "to",
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
					"id": "ujyhllyb",
					"name": "from",
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
					"id": "rkehpfep",
					"name": "amount",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": null,
						"noDecimal": true
					}
				}
			],
			"indexes": [],
			"listRule": "@request.auth.id = to.id || @request.auth.id = from.id",
			"viewRule": "@request.auth.id = to.id || @request.auth.id = from.id",
			"createRule": null,
			"updateRule": null,
			"deleteRule": "@request.auth.id = to.id",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("5bslbh0q4ky3wwx")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
