package endpoints

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func EndpointCollection(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api_ext/collections", func(c echo.Context) error {
			baseCollections, err := app.Dao().FindCollectionsByType(models.CollectionTypeBase)
			info := apis.RequestInfo(c)
			admin := info.Admin       // nil if not authenticated as admin
			record := info.AuthRecord // nil if not authenticated as regular auth record

			isGuest := admin == nil && record == nil

			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Internal Server error while getting records."})
			}

			if isGuest {
				return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Unauthorized."})
			}

			return c.JSON(http.StatusOK, baseCollections)
		} /* optional middlewares */)

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api_ext/users", func(c echo.Context) error {
			collection, err := app.Dao().FindCollectionByNameOrId("users")
			info := apis.RequestInfo(c)
			admin := info.Admin       // nil if not authenticated as admin
			record := info.AuthRecord // nil if not authenticated as regular auth record

			isGuest := admin == nil && record == nil

			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Internal Server error while getting users."})
			}
			if isGuest {
				return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Unauthorized."})
			}

			return c.JSON(http.StatusOK, collection)
		} /* optional middlewares */)

		return nil
	})
}
