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
			"id": "2c6eb226gn4jcs6",
			"created": "2024-02-20 18:06:09.235Z",
			"updated": "2024-02-20 18:06:09.235Z",
			"name": "v_b_jahre",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "trb9pvy0",
					"name": "Jahr",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 9999,
						"noDecimal": false
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
				"query": "SELECT Baende.id, Baende.Jahr\nFROM Baende\nWHERE Baende.Jahr > 0\nGROUP BY Baende.Jahr\nORDER BY Baende.Jahr;"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("2c6eb226gn4jcs6")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
