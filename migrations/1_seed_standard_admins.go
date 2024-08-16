// /////////////////////////////////////////////////
//
//	WARNING!!! FOR LOCAL TESTING PURPOSES ONLY! ///
//
// /////////////////////////////////////////////////
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
		err := update_user_fields(db)
		if err != nil {
			fmt.Println(err)
		}

		err = create_user(db)
		if err != nil {
			fmt.Println(err)
		}

		err = create_admin(db)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	}, func(db dbx.Builder) error { // revert op
		// add down queries...

		return nil
	})
}

func update_user_fields(db dbx.Builder) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	collection.Schema.AddField(&schema.SchemaField{
		System:      false,
		Id:          "users_role",
		Name:        "role",
		Type:        "select",
		Required:    true,
		Presentable: true,
		Unique:      false,
		Options: &schema.SelectOptions{
			MaxSelect: 1,
			Values: []string{
				"Admin",
				"Editor",
				"User",
			},
		},
	})

	collection.Schema.AddField(
		&schema.SchemaField{
			System:      false,
			Name:        "settings",
			Type:        schema.FieldTypeJson,
			Required:    false,
			Presentable: false,
			Options: &schema.JsonOptions{
				MaxSize: 4096000,
			},
		},
	)

	collection.Options["allowOAuth2Auth"] = false
	collection.Options["allowUsernameAuth"] = true
	collection.Options["allowEmailAuth"] = true

	collection.ListRule = types.Pointer("id = @request.auth.id || @request.auth.role = 'Admin'")
	collection.ViewRule = types.Pointer("id = @request.auth.id || @request.auth.role = 'Admin'")
	collection.CreateRule = types.Pointer("@request.auth.role = 'Admin'")
	collection.DeleteRule = types.Pointer("@request.auth.role = 'Admin'")
	collection.UpdateRule = types.Pointer("(id = @request.auth.id && role = @request.auth.role) || @request.auth.role = 'Admin'")
	collection.Options["manageRule"] = types.Pointer("(id = @request.auth.id && role = @request.auth.role) || @request.auth.role = 'Admin'")

	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}

func create_user(db dbx.Builder) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)
	record.SetUsername("Simon")
	record.Set("name", "Simon Martens")
	record.Set("role", "Admin")
	record.SetEmail("martens@tss-hd.de")
	record.SetPassword("qwerty123!!!")
	record.SetVerified(true)

	if err := dao.SaveRecord(record); err != nil {
		return err
	}

	return nil
}

func create_admin(db dbx.Builder) error {
	dao := daos.New(db)
	newAdmin := models.Admin{
		Avatar: 1,
		Email:  "martens@tss-hd.de",
	}
	newAdmin.SetPassword("admin123456!!")
	if err := dao.SaveAdmin(&newAdmin); err != nil {
		return err
	}

	return nil
}
