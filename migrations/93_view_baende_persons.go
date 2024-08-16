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
			"id": "fah6tsdj9y0238u",
			"created": "2024-02-20 18:04:59.996Z",
			"updated": "2024-02-20 18:04:59.996Z",
			"name": "v_b_akteure",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "bjxmyna2",
					"name": "b_id",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSize": 1
					}
				},
				{
					"system": false,
					"id": "uon65eax",
					"name": "Name",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "2pw1dp8p",
					"name": "Lebensdaten",
					"type": "text",
					"required": false,
					"presentable": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "iwbffdwp",
					"name": "Koerperschaft",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT b_id, Akteure.id as id, Akteure.Name, Akteure.Lebensdaten, Akteure.Koerperschaft \nFROM (\n    SELECT Baende.id AS b_id, value\n    FROM Baende, json_each(Baende.Herausgabe, '$')\n    WHERE json_array_length(Herausgabe) > 0 \n    GROUP BY value\n    UNION\n    SELECT Baende.id AS b_id, value\n    FROM Baende, json_each(Baende.Verlag, '$')\n    WHERE json_array_length(Verlag) > 0 \n    GROUP BY value\n    UNION\n    SELECT Baende.id AS b_id, value\n    FROM Baende, json_each(Baende.Vertrieb, '$')\n    WHERE json_array_length(Vertrieb) > 0 \n    GROUP BY value\n    UNION\n    SELECT Baende.id AS b_id, value\n    FROM Baende, json_each(Baende.Druck, '$')\n    WHERE json_array_length(Druck) > 0 \n    GROUP BY value\n)\nLEFT JOIN Akteure ON Akteure.id = value\nGROUP BY id\nORDER BY Akteure.Name COLLATE NOCASE ASC"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("fah6tsdj9y0238u")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
