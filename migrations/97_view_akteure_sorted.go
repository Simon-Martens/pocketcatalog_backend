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
			"id": "089azl3d5nxv2im",
			"created": "2024-02-20 21:11:23.129Z",
			"updated": "2024-02-20 21:11:23.129Z",
			"name": "v_akteure_sorted",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "dgcip4rn",
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
					"id": "zz1o4lqk",
					"name": "Koerperschaft",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "ecumhixo",
					"name": "Beruf",
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
					"id": "gpv3x3lt",
					"name": "Anmerkungen",
					"type": "editor",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"convertUrls": false
					}
				},
				{
					"system": false,
					"id": "gdzrffpd",
					"name": "GND",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": null,
						"onlyDomains": null
					}
				},
				{
				"system": false,
				"id": "w7wtfx8z",
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
				"name": "Nachweis",
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
				"id": "vwrg5z37",
				"name": "Pseudonyme",
				"type": "text",
				"required": false,
				"presentable": false,
				"unique": false,
				"options": {
					"min": null,
					"max": null,
					"pattern": ""
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
				"query": "SELECT id, Akteure.Name, Akteure.Koerperschaft, Akteure.Beruf, Akteure.Anmerkungen, Akteure.GND, Akteure.Lebensdaten, Akteure.Nachweis, Akteure.Pseudonyme FROM Akteure ORDER BY Akteure.Name COLLATE NOCASE ASC"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("089azl3d5nxv2im")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
