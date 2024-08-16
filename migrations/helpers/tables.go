package helpers

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func Create_Table(db dbx.Builder, collection *models.Collection) error {
	dao := daos.New(db)
	if err := dao.SaveCollection(collection); err != nil {
		return err
	}
	return nil
}

var AKTEURE_TABLE_MODEL = &models.Collection{
	Name:       "Akteure",
	Type:       models.CollectionTypeBase,
	ListRule:   types.Pointer(""),
	ViewRule:   types.Pointer(""),
	System:     true,
	CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	DeleteRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:        "Name",
			Type:        schema.FieldTypeText,
			Required:    true,
			Presentable: true,
		},
		&schema.SchemaField{
			Name:        "Lebensdaten",
			Type:        schema.FieldTypeText,
			Required:    false,
			Presentable: true,
		},
		&schema.SchemaField{
			Name:        "Koerperschaft",
			Type:        schema.FieldTypeBool,
			Presentable: false,
			Required:    false,
		},
		&schema.SchemaField{
			Name:     "Beruf",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Nachweis",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Pseudonyme",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Anmerkungen",
			Type:     schema.FieldTypeEditor,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "GND",
			Type:     schema.FieldTypeUrl,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Musenalm_ID",
			Type:     schema.FieldTypeText,
			Required: true,
		},
	),
	Indexes: types.JsonArray[string]{
		"CREATE INDEX idx_adef ON Akteure (Name,Lebensdaten)",
		"CREATE INDEX idx_aname ON Akteure (Name)",
		"CREATE INDEX idx_apseudonyme ON Akteure (Pseudonyme)",
		"CREATE INDEX idx_anachweis ON Akteure (Nachweis)",
		"CREATE INDEX idx_anm ON Akteure (Anmerkungen)",
		"CREATE UNIQUE INDEX idx_amus ON Akteure (Musenalm_ID)",
	},
}

var ORTE_TABLE_MODEL = &models.Collection{
	Name:       "Orte",
	Type:       models.CollectionTypeBase,
	ListRule:   types.Pointer(""),
	ViewRule:   types.Pointer(""),
	System:     true,
	CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	DeleteRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:        "Name",
			Type:        schema.FieldTypeText,
			Required:    true,
			Presentable: true,
		},
		&schema.SchemaField{
			Name:     "Fiktiv",
			Type:     schema.FieldTypeBool,
			Required: false,
		},
		&schema.SchemaField{
			Name:        "Geonames",
			Type:        schema.FieldTypeUrl,
			Required:    false,
			Presentable: false,
		},
		&schema.SchemaField{
			Name:        "Anmerkungen",
			Type:        schema.FieldTypeEditor,
			Presentable: false,
			Required:    false,
		},
		&schema.SchemaField{
			Name:     "Musenalm_ID",
			Type:     schema.FieldTypeText,
			Required: true,
		},
		&schema.SchemaField{
			Name:        "Genonames_Cache",
			Type:        schema.FieldTypeJson,
			Required:    false,
			Presentable: false,
			Options: &schema.JsonOptions{
				MaxSize: 2000000,
			},
		},
	),
	Indexes: types.JsonArray[string]{
		"CREATE UNIQUE INDEX idx_oname ON Orte (Name)",
		"CREATE INDEX idx_oanm ON Orte (Anmerkungen)",
		"CREATE INDEX idx_o_geoname ON Orte (Geonames)",
	},
}

var REIHENTITEL_TABLE_MODEL = &models.Collection{
	Name:       "Reihentitel",
	Type:       models.CollectionTypeBase,
	ListRule:   types.Pointer(""),
	ViewRule:   types.Pointer(""),
	System:     true,
	CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	DeleteRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:        "Titel",
			Type:        schema.FieldTypeText,
			Required:    true,
			Presentable: true,
		},
		&schema.SchemaField{
			Name:        "Nachweis",
			Type:        schema.FieldTypeText,
			Presentable: false,
			Required:    false,
		},
		&schema.SchemaField{
			Name:        "Anmerkungen",
			Type:        schema.FieldTypeEditor,
			Presentable: false,
			Required:    false,
		},
		&schema.SchemaField{
			Name:        "Musenalm_ID",
			Type:        schema.FieldTypeText,
			Presentable: false,
			Required:    true,
		},
	),
	Indexes: types.JsonArray[string]{
		"CREATE INDEX idx_rname ON Reihentitel (Titel)",
		"CREATE INDEX idx_rnachweis ON Reihentitel (Nachweis)",
		"CREATE INDEX idx_ranm ON Reihentitel (Anmerkungen)",
		"CREATE UNIQUE INDEX idx_rmus ON Reihentitel (Musenalm_ID)",
	},
}

var BAENDE_TABLE_MODEL = &models.Collection{
	Name:       "Baende",
	Type:       models.CollectionTypeBase,
	ListRule:   types.Pointer(""),
	ViewRule:   types.Pointer(""),
	System:     true,
	CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	DeleteRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:        "Titelangabe",
			Type:        schema.FieldTypeText,
			Required:    false,
			Presentable: false,
		},
		&schema.SchemaField{
			Name:        "Kurztitel",
			Type:        schema.FieldTypeText,
			Required:    true,
			Presentable: true,
		},
		&schema.SchemaField{
			Name:     "Jahr",
			Type:     schema.FieldTypeNumber,
			Required: false,
			Options: &schema.TextOptions{
				Min: types.Pointer(0),
				Max: types.Pointer(9999),
			},
		},
		&schema.SchemaField{
			Name:     "Verantwortlichkeitsangabe",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Ortsangabe",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Ausgabebezeichnung",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Nachweis",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:        "Biblio_ID",
			Type:        schema.FieldTypeNumber,
			Required:    false,
			Presentable: false,
		},
		&schema.SchemaField{
			Name:     "Struktur",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Norm_DEPR",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Reihentitel_DEPR",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Anmerkungen",
			Type:     schema.FieldTypeEditor,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Status",
			Type:     schema.FieldTypeSelect,
			Required: false,
			Options: &schema.SelectOptions{
				Values:    []string{"Original vorhanden", "Reprint vorhanden", "Fremde Herkunft"},
				MaxSelect: 3,
			},
		},
		&schema.SchemaField{
			Name:     "Gesichtet",
			Type:     schema.FieldTypeBool,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Erfasst",
			Type:     schema.FieldTypeBool,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Musenalm_ID",
			Type:     schema.FieldTypeText,
			Required: true,
		},
	),
	Indexes: types.JsonArray[string]{
		"CREATE INDEX idx_bdef ON Baende (Kurztitel,Jahr)",
		"CREATE INDEX idx_banm ON Baende (Anmerkungen)",
		"CREATE INDEX idx_bort ON Baende (Ortsangabe)",
		"CREATE INDEX idx_btit ON Baende (Titelangabe)",
		"CREATE INDEX idx_bkur ON Baende (Kurztitel)",
		"CREATE INDEX idx_bver ON Baende (Verantwortlichkeitsangabe)",
		"CREATE UNIQUE INDEX idx_bmus ON Baende (Musenalm_ID)",
	},
}

var INHALTE_TABLE_MODEL = &models.Collection{
	Name:       "Inhalte",
	Type:       models.CollectionTypeBase,
	ListRule:   types.Pointer(""),
	ViewRule:   types.Pointer(""),
	System:     true,
	CreateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	UpdateRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	DeleteRule: types.Pointer("@request.auth.id != '' && (@request.auth.role = 'Admin' || @request.auth.role = 'Editor')"),
	Schema: schema.NewSchema(
		&schema.SchemaField{
			Name:        "Objektnummer",
			Type:        schema.FieldTypeNumber,
			Required:    true,
			Presentable: true,
		},
		&schema.SchemaField{
			Name:        "Titelangabe",
			Type:        schema.FieldTypeText,
			Required:    false,
			Presentable: false,
		},
		&schema.SchemaField{
			Name:        "Scans",
			Type:        schema.FieldTypeFile,
			Required:    false,
			Presentable: false,
			Options: &schema.FileOptions{
				MaxSelect: 1000,
				MaxSize:   524288000,
				MimeTypes: []string{
					"application/pdf",
					"image/png",
					"image/vnd.mozilla.apng",
					"image/jpeg",
					"image/jp2",
					"image/jpx",
					"image/jpm",
					"image/gif",
					"image/jxs",
					"image/jxl",
					"image/x-xpixmap",
					"image/vnd.adobe.photoshop",
					"image/webp",
					"image/tiff",
					"image/bmp",
					"image/x-icon",
					"image/vnd.djvu",
					"image/bpg",
					"image/vnd.dwg",
					"image/x-icns",
					"image/heic",
					"image/heic-sequence",
					"image/heif",
					"image/heif-sequence",
					"image/vnd.radiance",
					"image/x-xcf",
					"image/x-gimp-pat",
					"image/x-gimp-gbr",
					"image/avif",
					"image/jxr",
					"image/svg+xml",
				},
				Thumbs: []string{
					"100x0",
					"250x0",
					"500x0",
				},
				Protected: false,
			},
		},
		&schema.SchemaField{
			Name:        "Urheberangabe",
			Type:        schema.FieldTypeText,
			Required:    false,
			Presentable: false,
		},
		&schema.SchemaField{
			Name:     "Incipit",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Seite",
			Type:     schema.FieldTypeText,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Anmerkungen",
			Type:     schema.FieldTypeEditor,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Paginierung",
			Type:     schema.FieldTypeSelect,
			Required: false,
			Options: &schema.SelectOptions{
				Values: []string{
					"Römische Seitenzählung",
					"Arabische Seitenzählung",
					"Alphabetische Seitenzählung",
					"Sonstige Seitenzählung",
					"1. Arabische Seitenzählung",
					"2. Arabische Seitenzählung",
					"3. Arabische Seitenzählung",
					"4. Arabische Seitenzählung",
					"5. Arabische Seitenzählung",
					"6. Arabische Seitenzählung",
					"7. Arabische Seitenzählung",
					"8. Arabische Seitenzählung",
					"1. Römische Seitenzählung",
					"2. Römische Seitenzählung",
					"3. Römische Seitenzählung",
					"4. Römische Seitenzählung",
					"5. Römische Seitenzählung",
					"6. Römische Seitenzählung",
					"7. Römische Seitenzählung",
					"8. Römische Seitenzählung",
				},
				MaxSelect: 1,
			},
		},
		&schema.SchemaField{
			Name:     "Typ",
			Type:     schema.FieldTypeSelect,
			Required: false,
			Options: &schema.SelectOptions{
				Values: []string{
					"Corrigenda",
					"Diagramm",
					"Gedicht/Lied",
					"Graphik",
					"Graphik-Verzeichnis",
					"graph. Anleitung",
					"graph. Strickanleitung",
					"graph. Tanzanleitung",
					"Inhaltsverzeichnis",
					"Kalendarium",
					"Karte",
					"Musikbeigabe",
					"Musikbeigaben-Verzeichnis",
					"Motto",
					"Prosa",
					"Rätsel",
					"Sammlung",
					"Spiegel",
					"szen. Darstellung",
					"Tabelle",
					"Tafel",
					"Titel",
					"Text",
					"Trinkspruch",
					"Umschlag",
					"Widmung",
				},
				MaxSelect: 26,
			},
		},
		&schema.SchemaField{
			Name:     "Vorschau",
			Type:     schema.FieldTypeBool,
			Required: false,
		},
		&schema.SchemaField{
			Name:     "Musenalm_ID",
			Type:     schema.FieldTypeText,
			Required: true,
		},
	),
	Indexes: types.JsonArray[string]{
		"CREATE INDEX idx_iobjn ON Inhalte (Objektnummer)",
		"CREATE INDEX idx_ityp ON Inhalte (Typ)",
		"CREATE INDEX idx_iincp ON Inhalte (Incipit)",
		"CREATE INDEX idx_iurh ON Inhalte (Urheberangabe)",
		"CREATE INDEX idx_ititl ON Inhalte (Titelangabe)",
		"CREATE INDEX idx_ianm ON Inhalte (Anmerkungen)",
		"CREATE UNIQUE INDEX idx_imus ON Inhalte (Musenalm_ID)",
	},
}

func Create_Table_Baende(db dbx.Builder) error {
	dao := daos.New(db)

	collection_orte, err := dao.FindCollectionByNameOrId("Orte")
	if err != nil {
		return err
	}

	collection := BAENDE_TABLE_MODEL

	collection.Schema.AddField(&schema.SchemaField{
		Name:        "Erscheinungsorte",
		Type:        schema.FieldTypeRelation,
		Required:    false,
		Presentable: false,
		Options: &schema.RelationOptions{
			CollectionId: collection_orte.Id,
		},
	})

	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_beorte ON Baende (Erscheinungsorte)")
	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}

func Create_Table_Inhalte(db dbx.Builder) error {
	dao := daos.New(db)

	collection_baende, err := dao.FindCollectionByNameOrId("Baende")
	if err != nil {
		return err
	}

	collection := INHALTE_TABLE_MODEL
	collection.Schema.AddField(&schema.SchemaField{
		Name:        "Band",
		Type:        schema.FieldTypeRelation,
		Required:    true,
		Presentable: true,
		Options: &schema.RelationOptions{
			CollectionId: collection_baende.Id,
			MaxSelect:    types.Pointer(1),
		},
	})

	collection.Indexes = append(collection.Indexes, "CREATE INDEX idx_iband ON Inhalte (Band)")
	if err := dao.SaveCollection(collection); err != nil {
		return err
	}

	return nil
}
