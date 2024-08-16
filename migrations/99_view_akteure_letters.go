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
			"created": "2024-03-20 18:06:27.728Z",
			"updated": "2024-03-20 18:06:27.728Z",
			"name": "v_akteure_letters",
			"type": "view",
			"system": false,
			"schema": [],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": "",
			"updateRule": "",
			"deleteRule": "",
			"options": {
				"query": "SELECT DISTINCT upper(substr(Name, 1, 1)) as id \n  FROM Akteure\n  WHERE Akteure.Name IS NOT \"\"\n AND Akteure.Koerperschaft IS NOT true\n ORDER BY id COLLATE NOCASE ASC\n"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("view_akteure_letters")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
