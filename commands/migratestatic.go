package commands

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/spf13/cobra"
)

type ImageInfo struct {
	Filename    string
	Title       string
	Description string
}

func readDescriptions(filePath string) (map[string]ImageInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	imageInfos := make(map[string]ImageInfo)
	var currentInfo ImageInfo

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			if currentInfo.Filename != "" {
				imageInfos[currentInfo.Filename] = currentInfo
			}
			currentInfo = ImageInfo{Filename: strings.TrimPrefix(line, "# ")}
		} else if strings.HasPrefix(line, "## ") {
			currentInfo.Title = strings.TrimPrefix(line, "## ")
		} else if strings.HasPrefix(line, "### ") {
			currentInfo.Description = strings.TrimPrefix(line, "### ")
		}
	}

	if currentInfo.Filename != "" {
		imageInfos[currentInfo.Filename] = currentInfo
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return imageInfos, nil
}

func MigrateStatics(app *pocketbase.PocketBase, defaultdir string) {
	app.RootCmd.AddCommand(&cobra.Command{
		Use:   "migratestatics [folder] [descriptions_file]",
		Short: "Migrate static assets from a folder (default: ./Static-Bilder) to the database",
		Run: func(cmd *cobra.Command, args []string) {
			path := defaultdir
			var descriptionPath string

			if len(args) >= 1 {
				path = args[0]
			}
			if len(args) >= 2 {
				descriptionPath = args[1]
			} else {
				// Look for beschreibungen.txt in the image directory
				descriptionPath = filepath.Join(path, "beschreibungen.txt")
				if _, err := os.Stat(descriptionPath); os.IsNotExist(err) {
					// If not found, look in the current directory
					descriptionPath = "beschreibungen.txt"
				}
			}

			imageInfos, err := readDescriptions(descriptionPath)
			if err != nil {
				app.Logger().Error("Failed to read descriptions file:", err)
				app.Logger().Info("Proceeding without descriptions")
				imageInfos = make(map[string]ImageInfo)
			}

			table, err := app.Dao().FindCollectionByNameOrId("Bilder")
			if err != nil {
				app.Logger().Error("Could not find Table Bilder! You need to execute table migrations first!")
				return
			}

			e := func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
				name := fileInfo.Name()
				if !fileInfo.IsDir() && (strings.HasSuffix(name, ".png")) {
					record := models.NewRecord(table)
					form := forms.NewRecordUpsert(app, record)

					info, exists := imageInfos[name]
					if exists {
						form.LoadData(map[string]any{
							"Titel":        info.Title,
							"Beschreibung": info.Description,
						})
					} else {
						// Use filename without extension as title and leave description empty
						titleWithoutExt := strings.TrimSuffix(name, filepath.Ext(name))
						form.LoadData(map[string]any{
							"Titel":        titleWithoutExt,
							"Beschreibung": "",
						})
						app.Logger().Info("No description found for file:", name, "- Using filename as title")
					}

					f, err := filesystem.NewFileFromPath(path)
					if err != nil {
						log.Println(err)
						return nil
					}
					form.AddFiles("Bilder", f)

					if err := form.Submit(); err != nil {
						log.Println(err)
						return nil
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
