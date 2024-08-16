package main

import (
	"log"

	_ "github.com/Simon-Martens/pocketcatalog_backend/migrations"

	"github.com/Simon-Martens/pocketcatalog_backend/endpoints"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	// endpoints.EndpointFTS(app)
	endpoints.EndpointCollection(app)
	endpoints.EndpointStatic(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
