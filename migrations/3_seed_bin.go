package migrations

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		err := create_table_bin(db)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	}, func(db dbx.Builder) error { // revert op
		// add down queries...

		return nil
	})
}

func create_table_bin(db dbx.Builder) error {
	dao := daos.New(db)

	collection := &models.Collection{
		Name:       "Bin",
		Type:       models.CollectionTypeBase,
		ListRule:   types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		ViewRule:   types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
		DeleteRule: types.Pointer("@request.auth.id != '' && @request.auth.role = 'Admin'"),
		Schema: schema.NewSchema(
			&schema.SchemaField{
				System:      false,
				Name:        "oldid",
				Type:        schema.FieldTypeText,
				Required:    true,
				Presentable: true,
			},
			&schema.SchemaField{
				System:      false,
				Name:        "oldcollection",
				Type:        schema.FieldTypeText,
				Required:    true,
				Presentable: true,
				Options: &schema.JsonOptions{
					MaxSize: 4096000,
				},
			},
			&schema.SchemaField{
				System:      false,
				Name:        "oldjson",
				Type:        schema.FieldTypeJson,
				Required:    true,
				Presentable: false,
				Options: &schema.JsonOptions{
					MaxSize: 4096000,
				},
			},
			&schema.SchemaField{
				System:      false,
				Name:        "olddeps",
				Type:        schema.FieldTypeJson,
				Required:    false,
				Presentable: false,
				Options: &schema.JsonOptions{
					MaxSize: 4096000,
				},
			},
		),
		Indexes: types.JsonArray[string]{
			"CREATE INDEX idx_bincol ON Bin (oldcollection)",
			"CREATE INDEX idx_bindata ON Bin (oldjson)",
			"CREATE INDEX idx_binid ON Bin (oldid)",
		},
	}

	return dao.SaveCollection(collection)
}
