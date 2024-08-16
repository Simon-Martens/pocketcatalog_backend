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
			"id": "nv3w53x3vyq7nfb",
			"created": "2024-02-20 18:06:27.728Z",
			"updated": "2024-02-20 18:06:27.728Z",
			"name": "v_reihentitel_letters",
			"type": "view",
			"system": false,
			"schema": [],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {
				"query": "SELECT DISTINCT upper(substr(Titel, 1, 1)) as id \n  FROM Reihentitel\n  WHERE Reihentitel.Titel IS NOT \"\"\n  ORDER BY id COLLATE NOCASE ASC\n"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("nv3w53x3vyq7nfb")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
