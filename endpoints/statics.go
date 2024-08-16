package endpoints

import (
	"github.com/Simon-Martens/pocketcatalog_backend/ui"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func EndpointStatic(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(ui.DistDirFS, true))

		return nil
	})

}
