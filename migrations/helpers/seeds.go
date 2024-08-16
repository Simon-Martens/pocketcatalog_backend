package helpers

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func Seed_Table_Akteure(db dbx.Builder, akteure *Akteure) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Akteure")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(akteure.Akteure); i++ {
		record := models.NewRecord(collection)
		record.Set("Koerperschaft", akteure.Akteure[i].Körperschaft)
		record.Set("Name", NormalizeString(&(akteure.Akteure[i].Name)))
		record.Set("Nachweis", NormalizeString(&(akteure.Akteure[i].Nachweis)))
		record.Set("Lebensdaten", NormalizeString(&(akteure.Akteure[i].Lebensdaten)))
		record.Set("Beruf", NormalizeString(&(akteure.Akteure[i].Beruf)))
		record.Set("Pseudonyme", NormalizeString(&(akteure.Akteure[i].Pseudonyme)))
		record.Set("Anmerkungen", NormalizeString(&(akteure.Akteure[i].Anmerkungen)))
		record.Set("Musenalm_ID", "A-"+akteure.Akteure[i].ID)
		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(record); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func Seed_Table_Orte(db dbx.Builder, orte *Orte) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Orte")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(orte.Orte); i++ {
		record := models.NewRecord(collection)
		record.Set("Name", NormalizeString(&orte.Orte[i].Name))
		record.Set("Anmerkungen", NormalizeString(&orte.Orte[i].Anmerkungen))
		record.Set("Fiktiv", orte.Orte[i].Fiktiv)
		record.Set("Musenalm_ID", "O-"+orte.Orte[i].ID)
		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(record); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil

}

func Seed_Table_Reihentitel(db dbx.Builder, reihentitel *Reihentitel) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Reihentitel")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(reihentitel.Reihen); i++ {
		record := models.NewRecord(collection)
		if reihentitel.Reihen[i].Titel != "" {
			record.Set("Titel", NormalizeString(&reihentitel.Reihen[i].Titel))
		} else if reihentitel.Reihen[i].Sortiername != "" {
			record.Set("Titel", NormalizeString(&reihentitel.Reihen[i].Sortiername))
		}
		record.Set("Nachweis", NormalizeString(&reihentitel.Reihen[i].Nachweis))
		record.Set("Anmerkungen", NormalizeString(&reihentitel.Reihen[i].Anmerkungen))
		record.Set("Musenalm_ID", "R-"+reihentitel.Reihen[i].ID)
		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(record); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func Seed_Table_Bände(db dbx.Builder, bände *Bände) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Baende")

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(bände.Bände); i++ {
		record := models.NewRecord(collection)

		for _, v := range bände.Bände[i].Orte {
			ort, err := dao.FindFirstRecordByData("Orte", "Musenalm_ID", "O-"+v.Value)
			if err != nil {
				fmt.Println("Ort mit der ID " + v.Value + " nicht gefunden!")
				fmt.Println("Relation Band " + bände.Bände[i].ID + " - Ort " + v.Value + " nicht gesetzt!")
				fmt.Println(err)
				continue
			}
			before := record.GetStringSlice("Erscheinungsorte")
			record.Set("Erscheinungsorte", append(before, ort.Id))
		}

		record.Set("Titelangabe", NormalizeString(&bände.Bände[i].Titelangabe))
		// Gets replaced by Bevorzugter_Reihentitel + Jahr where applicable later on
		record.Set("Kurztitel", NormalizeString(&bände.Bände[i].ReihentitelALT))
		record.Set("Reihentitel_DEPR", NormalizeString(&bände.Bände[i].ReihentitelALT))
		record.Set("Nachweis", NormalizeString(&bände.Bände[i].Nachweis))
		record.Set("Anmerkungen", NormalizeString(&bände.Bände[i].Anmerkungen))
		record.Set("Jahr", bände.Bände[i].Jahr)
		record.Set("Verantwortlichkeitsangabe", NormalizeString(&bände.Bände[i].Verantwortlichkeitsangabe))
		record.Set("Ortsangabe", NormalizeString(&bände.Bände[i].Ortsangabe))
		record.Set("Biblio_ID", bände.Bände[i].BiblioID)
		record.Set("Struktur", NormalizeString(&bände.Bände[i].Struktur))
		record.Set("Norm_DEPR", NormalizeString(&bände.Bände[i].Norm))
		record.Set("Gesichtet", bände.Bände[i].Gesichtet)
		record.Set("Erfasst", bände.Bände[i].Erfasst)
		record.Set("Status", bände.Bände[i].Status.Value)
		record.Set("Musenalm_ID", "B-"+bände.Bände[i].ID)
		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(record); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func Seed_Table_Inhalte(db dbx.Builder, inhalte *Inhalte) error {
	dao := daos.New(db)
	collection, err := dao.FindCollectionByNameOrId("Inhalte")

	vorschau_items := map[string]bool{"124272": true, "127028": true, "117110": true, "106117": true, "72078": true, "34024": true, "44562": true, "115170": true}

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(inhalte.Inhalte); i++ {
		record := models.NewRecord(collection)

		// Sadly NOT possible ATM:
		// if _, err := os.Stat("Almanach-Bilder/alm-" + inhalte.Inhalte[i].Band); err == nil {
		// 	p := "Almanach-Bilder/alm-" + inhalte.Inhalte[i].Band + "/alm-" + inhalte.Inhalte[i].Band + "-" + inhalte.Inhalte[i].ID + ".jpg"
		// 	if _, err := os.Stat(p); err == nil {
		// 		fmt.Println("Sapperlot! " + p + " existiert!")
		// 		f, _ := filesystem.NewFileFromPath(p)
		// 		record.Set("Scans", f)
		// 	}
		// }

		band, err := dao.FindFirstRecordByData("Baende", "Musenalm_ID", "B-"+inhalte.Inhalte[i].Band)
		if err != nil {
			fmt.Println("Band mit der ID " + inhalte.Inhalte[i].Band + " nicht gefunden!")
			fmt.Println("Relation Band " + inhalte.Inhalte[i].Band + " - Inhalt " + inhalte.Inhalte[i].ID + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}

		record.Set("Titelangabe", NormalizeString(&inhalte.Inhalte[i].Titelangabe))
		record.Set("Urheberangabe", NormalizeString(&inhalte.Inhalte[i].Urheberangabe))
		record.Set("Objektnummer", &inhalte.Inhalte[i].Objektnummer)
		record.Set("Anmerkungen", NormalizeString(&inhalte.Inhalte[i].Anmerkungen))
		record.Set("Incipit", NormalizeString(&inhalte.Inhalte[i].Incipit))
		record.Set("Seite", NormalizeString(&inhalte.Inhalte[i].Seite))
		record.Set("Typ", inhalte.Inhalte[i].Typ.Value)
		record.Set("Band", band.Id)
		if len(inhalte.Inhalte[i].Paginierung) > 0 {
			value := []string{}
			switch inhalte.Inhalte[i].Paginierung {
			case "ar":
				value = append(value, "Arabische Seitenzählung")
			case "röm":
				value = append(value, "Römische Seitenzählung")
			case "ar1":
				value = append(value, "1. Arabische Seitenzählung")
			case "ar2":
				value = append(value, "2. Arabische Seitenzählung")
			case "ar3":
				value = append(value, "3. Arabische Seitenzählung")
			case "ar4":
				value = append(value, "4. Arabische Seitenzählung")
			case "ar5":
				value = append(value, "5. Arabische Seitenzählung")
			case "ar6":
				value = append(value, "6. Arabische Seitenzählung")
			case "ar7":
				value = append(value, "7. Arabische Seitenzählung")
			case "röm1":
				value = append(value, "1. Römische Seitenzählung")
			case "röm2":
				value = append(value, "2. Römische Seitenzählung")
			case "röm3":
				value = append(value, "3. Römische Seitenzählung")
			case "röm4":
				value = append(value, "4. Römische Seitenzählung")
			case "sonst":
				value = append(value, "Sonstige Seitenzählung")
			}

			record.Set("Paginierung", value)
		}

		_, ok := vorschau_items[inhalte.Inhalte[i].ID]
		if ok {
			record.Set("Vorschau", true)
		}

		record.Set("Musenalm_ID", "I-"+inhalte.Inhalte[i].ID)
		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(record); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func Seed_Fields_Relation_Bände_Reihen(db dbx.Builder, relationen *Relationen_Bände_Reihen) error {
	dao := daos.New(db)

	for i := 0; i < len(relationen.Relationen); i++ {
		band, err := dao.FindFirstRecordByData("Baende", "Musenalm_ID", "B-"+relationen.Relationen[i].Band)
		if err != nil {
			fmt.Println("Band mit der ID " + relationen.Relationen[i].Band + " nicht gefunden!")
			fmt.Println("Relation Band " + relationen.Relationen[i].Band + " - Reihe " + relationen.Relationen[i].Reihe + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}
		reihe, err := dao.FindFirstRecordByData("Reihentitel", "Musenalm_ID", "R-"+relationen.Relationen[i].Reihe)
		if err != nil {
			fmt.Println("Reihe mit der ID " + relationen.Relationen[i].Reihe + " nicht gefunden!")
			fmt.Println("Relation Band " + relationen.Relationen[i].Band + " - Reihe " + relationen.Relationen[i].Reihe + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}
		switch relationen.Relationen[i].Relation {
		case "1":
			jahr := ""
			if band.GetString("Jahr") != "" && band.GetString("Jahr") != "0" {
				jahr = " " + band.GetString("Jahr")
			}
			band.Set("Kurztitel", reihe.GetString("Titel")+jahr)
			band.Set("Bevorzugter_Reihentitel", reihe.Id)
		case "2":
			before := band.GetStringSlice("Alternativer_Reihentitel")
			band.Set("Alternativer_Reihentitel", append(before, reihe.Id))
		case "3":
			band.Set("Franzoesischer_Reihentitel", reihe.Id)
		case "4":
			band.Set("Deutscher_Reihentitel", reihe.Id)
		case "5":
			before := band.GetStringSlice("Alternatives_Titelblatt")
			band.Set("Alternatives_Titelblatt", append(before, reihe.Id))
		case "6":
			before := band.GetStringSlice("TA_von")
			band.Set("TA_von", append(before, reihe.Id))
		case "7":
			before := band.GetStringSlice("hat_TA")
			band.Set("hat_TA", append(before, reihe.Id))
		}

		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(band); err != nil {
			fmt.Println(err)
			continue
		}
	}
	return nil
}

func Seed_Fields_Relation_Bände_Akteure(db dbx.Builder, relationen *Relationen_Bände_Akteure) error {
	dao := daos.New(db)

	for i := 0; i < len(relationen.Relationen); i++ {
		band, err := dao.FindFirstRecordByData("Baende", "Musenalm_ID", "B-"+relationen.Relationen[i].Band)
		if err != nil {
			fmt.Println("Band mit der ID " + relationen.Relationen[i].Band + " nicht gefunden!")
			fmt.Println("Relation Band " + relationen.Relationen[i].Band + " - Akteur " + relationen.Relationen[i].Akteur + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}

		akteur, err := dao.FindFirstRecordByData("Akteure", "Musenalm_ID", "A-"+relationen.Relationen[i].Akteur)
		if err != nil {
			fmt.Println("Akteur mit der ID " + relationen.Relationen[i].Akteur + " nicht gefunden!")
			fmt.Println("Relation Band " + relationen.Relationen[i].Band + " - Akteur " + relationen.Relationen[i].Akteur + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}

		switch relationen.Relationen[i].Relation {
		case "8":
			before := band.GetStringSlice("Vertrieb")
			band.Set("Vertrieb", append(before, akteur.Id))
		case "7":
			before := band.GetStringSlice("Druck")
			band.Set("Druck", append(before, akteur.Id))
		case "6":
			before := band.GetStringSlice("Verlag")
			band.Set("Verlag", append(before, akteur.Id))
		case "5":
			before := band.GetStringSlice("Herausgabe")
			band.Set("Herausgabe", append(before, akteur.Id))
		}

		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(band); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}

func Seed_Fields_Relation_Inhalte_Akteure(db dbx.Builder, relationen *Relationen_Inhalte_Akteure) error {
	dao := daos.New(db)

	for i := 0; i < len(relationen.Relationen); i++ {
		inhalt, err := dao.FindFirstRecordByData("Inhalte", "Musenalm_ID", "I-"+relationen.Relationen[i].Band)
		if err != nil {
			fmt.Println("Inhalt mit der ID " + relationen.Relationen[i].Band + " nicht gefunden!")
			fmt.Println("Relation Inhalt " + relationen.Relationen[i].Band + " - Akteur " + relationen.Relationen[i].Akteur + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}

		akteur, err := dao.FindFirstRecordByData("Akteure", "Musenalm_ID", "A-"+relationen.Relationen[i].Akteur)
		if err != nil {
			fmt.Println("Akteur mit der ID " + relationen.Relationen[i].Akteur + " nicht gefunden!")
			fmt.Println("Relation Inhalt " + relationen.Relationen[i].Band + " - Akteur " + relationen.Relationen[i].Akteur + " nicht gesetzt!")
			fmt.Println(err)
			continue
		}

		before := inhalt.GetStringSlice("Geschaffen")
		inhalt.Set("Geschaffen", append(before, akteur.Id))

		// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
		if err := dao.SaveRecord(inhalt); err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil
}
