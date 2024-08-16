package helpers

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func Create_Relation_Bände_Reihen(db dbx.Builder) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Baende")
	if err != nil {
		return err
	}

	collection_reihe, err := dao.FindCollectionByNameOrId("Reihentitel")
	if err != nil {
		return err
	}

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Bevorzugter_Reihentitel",
			Type:     schema.FieldTypeRelation,
			Required: true,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
				MaxSelect:    types.Pointer(1),
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Alternativer_Reihentitel",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Franzoesischer_Reihentitel",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
				MaxSelect:    types.Pointer(1),
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Deutscher_Reihentitel",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
				MaxSelect:    types.Pointer(1),
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Alternatives_Titelblatt",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "TA_von",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "hat_TA",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_reihe.Id,
			},
		},
	)

	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_bevr ON Baende (Bevorzugter_Reihentitel)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_altr ON Baende (Alternativer_Reihentitel)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_fran ON Baende (Franzoesischer_Reihentitel)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_deut ON Baende (Deutscher_Reihentitel)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_altt ON Baende (Alternatives_Titelblatt)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_tavon ON Baende (TA_von)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_hatta ON Baende (hat_TA)")

	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}

func Create_Relation_Bände_Akteure(db dbx.Builder) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Baende")
	if err != nil {
		return err
	}

	collection_akteure, err := dao.FindCollectionByNameOrId("Akteure")
	if err != nil {
		return err
	}

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Herausgabe",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Verlag",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Druck",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Vertrieb",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_bherausgabe ON Baende (Herausgabe)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_bverlag ON Baende (Verlag)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_bdruck ON Baende (Druck)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_bvertrieb ON Baende (Vertrieb)")

	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}

func Create_Fields_Inhalte_Akteure(db dbx.Builder) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Inhalte")
	if err != nil {
		return err
	}

	collection_akteure, err := dao.FindCollectionByNameOrId("Akteure")
	if err != nil {
		return err
	}

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Geschaffen",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Geschrieben",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Gezeichnet",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Schema.AddField(
		&schema.SchemaField{
			Name:     "Gestochen",
			Type:     schema.FieldTypeRelation,
			Required: false,
			Options: &schema.RelationOptions{
				CollectionId: collection_akteure.Id,
			},
		},
	)

	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_geschaffen ON Inhalte (Geschaffen)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_geschrieben ON Inhalte (Geschrieben)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_gezeichnet ON Inhalte (Gezeichnet)")
	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_gestochen ON Inhalte (Gestochen)")

	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}
