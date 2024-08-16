package commands

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/spf13/cobra"
)

// Adding the possibility of migrating almanach-images by executiong pocketbase migrateimages [folder]
func MigrateImages(app *pocketbase.PocketBase, defaultdir string) {
	app.RootCmd.AddCommand(&cobra.Command{
		Use:   "migrateimages",
		Short: "Migrate images from a folder (default: ./Almanach-Bilder) to the database",
		Run: func(cmd *cobra.Command, args []string) {
			path := defaultdir
			if len(args) != 0 {
				path = args[0]
			}

			e := func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
				if !fileInfo.IsDir() {
					basesplit := strings.Split(fileInfo.Name(), "-")
					if len(basesplit) == 3 {
						extensionsplit := strings.Split(basesplit[2], ".")
						if len(extensionsplit) == 2 {
							commaseperatorsplit := strings.Split(extensionsplit[0], ",")
							id := commaseperatorsplit[0]
							record, err := app.Dao().FindFirstRecordByData("Inhalte", "Musenalm_ID", "I-"+id)
							if err != nil {
								log.Print("Fehler! Datei: " + path + " ")
								log.Print(err)
								log.Print("\n")
								return nil
							}
							f, err := filesystem.NewFileFromPath(path)
							if err != nil {
								log.Println(err)
								return nil
							}
							form := forms.NewRecordUpsert(app, record)
							form.AddFiles("Scans", f)
							// validate and submit (internally it calls app.Dao().SaveRecord(record) in a transaction)
							if err := form.Submit(); err != nil {
								log.Println(err)
								return nil
							}
						}
					}
				}
				return nil
			}

			if err := filepath.Walk(path, e); err != nil {
				log.Fatal(err)
			}
		},
	})
}
