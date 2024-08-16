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
			"id": "2sbj4zh54sm8mb9",
			"created": "2024-02-20 18:05:31.242Z",
			"updated": "2024-02-20 18:05:31.242Z",
			"name": "v_b_hrsgpseu",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "fmk1khum",
					"name": "Verantwortlichkeitsangabe",
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
					"id": "cetwun6p",
					"name": "Herausgabe",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "moec6f4bvenx3u9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT id, Baende.Verantwortlichkeitsangabe, Baende.Herausgabe\nFROM Baende\nWHERE Baende.Verantwortlichkeitsangabe IS NOT \"\"\nGROUP BY Baende.Verantwortlichkeitsangabe\nORDER BY Baende.Verantwortlichkeitsangabe COLLATE NOCASE ASC;"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("2sbj4zh54sm8mb9")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
