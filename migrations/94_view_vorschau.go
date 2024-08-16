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
			"id": "brhpewsk0h1e6u0",
			"created": "2024-02-20 18:07:52.869Z",
			"updated": "2024-02-20 18:07:52.869Z",
			"name": "v_vorschau",
			"type": "view",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "dlrzerc6",
					"name": "Band",
					"type": "relation",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"collectionId": "40iuvyhi5lwp4yt",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "jxnkp5ox",
					"name": "Seite",
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
					"id": "hhnzvcgw",
					"name": "Scans",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"mimeTypes": [
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
							"image/svg+xml"
						],
						"thumbs": [
							"100x0",
							"250x0",
							"500x0"
						],
						"maxSelect": 1000,
						"maxSize": 524288000,
						"protected": false
					}
				},
				{
					"system": false,
					"id": "z6cwexis",
					"name": "Geschaffen",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "moec6f4bvenx3u9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "0r3oqbgv",
					"name": "Geschrieben",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "moec6f4bvenx3u9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "zshocuos",
					"name": "Gezeichnet",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "moec6f4bvenx3u9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "glhk5aiw",
					"name": "Gestochen",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "moec6f4bvenx3u9",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "ozddh8wi",
					"name": "Objektnummer",
					"type": "number",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "rccnkogp",
					"name": "Urheberangabe",
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
					"id": "hgphcrcn",
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
					"id": "mf9ckzjy",
					"name": "Typ",
					"type": "select",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 26,
						"values": [
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
							"Widmung"
						]
					}
				},
				{
					"system": false,
					"id": "fnqzsuec",
					"name": "Titelangabe",
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
					"id": "vz6ly7sr",
					"name": "Incipit",
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
					"id": "ba3b63pe",
					"name": "Musenalm_ID",
					"type": "text",
					"required": true,
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
				"query": "SELECT id, Inhalte.Band, Inhalte.Seite, Inhalte.Scans, Inhalte.Geschaffen, Inhalte.Geschrieben, Inhalte.Gezeichnet, Inhalte.Gestochen, Inhalte.Objektnummer, Inhalte.Urheberangabe, Inhalte.Anmerkungen, Inhalte.Typ, Inhalte.Titelangabe, Inhalte.Incipit, Inhalte.Musenalm_ID\nFROM Inhalte\nWHERE Inhalte.Vorschau = TRUE"
			}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("brhpewsk0h1e6u0")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
