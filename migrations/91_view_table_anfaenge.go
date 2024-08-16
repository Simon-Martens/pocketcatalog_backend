package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		options := map[string]any{}
		json.Unmarshal([]byte(`{
			"query": "SELECT id, Anfang, Band\nFROM (\n    SELECT id, Band, Incipit AS Anfang \n    FROM Inhalte \n    WHERE Anfang != \"\" AND Anfang NOT LIKE \"#%\" AND Anfang NOT LIKE \"$%\"\n    UNION\n    SELECT id, Band, Titelangabe AS Anfang \n    FROM Inhalte \n    WHERE Anfang != \"\" AND Anfang NOT LIKE \"#%\" AND Anfang NOT LIKE \"$%\"\n)\nORDER BY Anfang COLLATE NOCASE ASC"
		}`), &options)

		collection := &models.Collection{
			Name:     "v_anfaenge",
			Type:     models.CollectionTypeView,
			ListRule: types.Pointer(""),
			ViewRule: types.Pointer(""),
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Name:        "Anfang",
					Type:        schema.FieldTypeJson,
					Required:    false,
					Presentable: false,
				},
				&schema.SchemaField{
					Name:        "Band",
					Type:        schema.FieldTypeJson,
					Required:    false,
					Presentable: false,
				}),
			Options: options,
		}

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error { // revert op
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("v_anfaenge")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

// func dc_inhalte(db dbx.Builder) error {
// 	dao := daos.New(db)
// 	collection, err := dao.FindCollectionByNameOrId("Inhalte")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return nil
// }
