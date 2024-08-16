package migrations

import (
	"fmt"

	"github.com/Simon-Martens/pocketcatalog_backend/migrations/helpers"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		lib, err := helpers.ReadXMLData("data/")
		if err != nil {
			fmt.Println(err)
		}

		*lib.Reihentitel = helpers.Sanitze_Reihentitel(lib.Reihentitel, lib.Relationen_Bände_Reihen)

		fmt.Println("Creating tables and fields...")
		fmt.Println("Creating Table Akteure...")
		// We create Akteure & Reihentitel & Orte first, since they are referenced by other tables,
		// and hold no references themselves.
		if err = helpers.Create_Table(db, helpers.AKTEURE_TABLE_MODEL); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Creating Table Orte...")
		if err = helpers.Create_Table(db, helpers.ORTE_TABLE_MODEL); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Creating Table Reihentitel...")
		if err = helpers.Create_Table(db, helpers.REIHENTITEL_TABLE_MODEL); err != nil {
			fmt.Println(err)
			return err
		}

		// Baende must be created before Inhalte, since Inhalte has a foreign key to Baende
		fmt.Println("Creating Table Bände...")
		if err = helpers.Create_Table_Baende(db); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Creating Table Inhalte...")
		if err = helpers.Create_Table_Inhalte(db); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Creating fields Bände-Reihen...")
		if err = helpers.Create_Relation_Bände_Reihen(db); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Creating fields Bände-Akteure...")
		if err = helpers.Create_Relation_Bände_Akteure(db); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Creating fields Inhalte-Akteure...")
		if err = helpers.Create_Fields_Inhalte_Akteure(db); err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Seeding tables...")
		// Same order as above
		helpers.Seed_Table_Akteure(db, lib.Akteure)
		helpers.Seed_Table_Orte(db, lib.Orte)
		helpers.Seed_Table_Reihentitel(db, lib.Reihentitel)
		helpers.Seed_Table_Bände(db, lib.Bände)
		helpers.Seed_Table_Inhalte(db, lib.Inhalte)
		helpers.Seed_Fields_Relation_Bände_Reihen(db, lib.Relationen_Bände_Reihen)
		helpers.Seed_Fields_Relation_Bände_Akteure(db, lib.Relationen_Bände_Akteure)
		helpers.Seed_Fields_Relation_Inhalte_Akteure(db, lib.Relationen_Inhalte_Akteure)

		return nil
	}, func(db dbx.Builder) error { // revert op
		dao := daos.New(db)
		err := dao.DeleteTable("Akteure")
		if err != nil {
			fmt.Println(err)
		}

		err = dao.DeleteTable("Reihentitel")
		if err != nil {
			fmt.Println(err)
		}

		err = dao.DeleteTable("Orte")
		if err != nil {
			fmt.Println(err)
		}

		err = dao.DeleteTable("Baende")
		if err != nil {
			fmt.Println(err)
		}

		err = dao.DeleteTable("Inhalte")
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
}
