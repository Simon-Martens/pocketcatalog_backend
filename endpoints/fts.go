package endpoints

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/resolvers"
	"github.com/pocketbase/pocketbase/tools/list"
	"github.com/pocketbase/pocketbase/tools/search"
)

// Collection: the collection to be served. GEts an appended _fts for the corresponding FTS5 table. Right now there is inhalte_fts and baende_fts.
// Term: the search term to be matched against the FTS5 table. Everythinbg that comes after the
// Expand: the fields to be expanded. This is a comma separated list of fields that are to be expanded.
// Sort: the fields to be sorted by. This is a comma separated list of fields that are to be sorted by. Sadly it's not possible to sort by Joined Fields.
// Fields: the fields to be searched in. This is a comma separated list of fields that are to be matched against the search term. If empty, all fields are searched.
// Filter: WHERE expression to be applied to the RECORDS table.
type FTSRequest struct {
	Collection string `param:"collection" query:"collection" form:"collection" json:"collection" xml:"collection"`
	Term       string `param:"term" query:"term" form:"term" json:"term" xml:"term"`
	Expand     string `param:"expand" query:"expand" form:"expand" json:"expand" xml:"expand"`
	Sort       string `param:"sort" query:"sort" form:"sort" json:"sort" xml:"sort"`
	Fields     string `param:"fields" query:"fields" form:"fields" json:"fields" xml:"fields"`
	Filter     string `param:"filter" query:"filter" form:"filter" json:"filter" xml:"filter"`
}

// This adds an FTS Endpoint that accepts a search query (see above), then looks up if there is an FTS5 table available and, if so, MATCHes against the term.
// It then returns the original RECORDS from the original table
// TODO: It's only inefficiency as of right now is that it's external table backed FTS5,
// yet it searches mostly only a subset of all columns. FTS already returns the original
// columns; but we still have to go get the rest of the columns and the expanded Fields
// via another Select Query. It's nit bad though since there's an Index on ID.
func EndpointFTS(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.POST("/fts/:collection", func(c echo.Context) error {
			var ftsrequest FTSRequest
			err := c.Bind(&ftsrequest)
			if err != nil || ftsrequest.Term == "" || ftsrequest.Collection == "" {
				return apis.NewBadRequestError("The request is bad!", err)
			}

			ftsrequest.Collection = strings.TrimSpace(ftsrequest.Collection)
			ftsrequest.Term = strings.TrimSpace(ftsrequest.Term)
			ftsrequest.Sort = strings.TrimSpace(ftsrequest.Sort)
			ftsrequest.Expand = strings.TrimSpace(ftsrequest.Expand)
			ftsrequest.Fields = strings.TrimSpace(ftsrequest.Fields)

			collection, err := app.Dao().FindCollectionByNameOrId(ftsrequest.Collection)
			if err != nil {
				return apis.NewBadRequestError("The collection does not exist!", err)
			}

			e := []dbx.NullStringMap{}

			cname := ftsrequest.Collection + "_fts"

			if ftsrequest.Fields == "" {
				err = app.Dao().DB().
					NewQuery(
						"SELECT * FROM " +
							cname +
							" WHERE " +
							cname +
							" MATCH {:term} LIMIT 10000").
					Bind(dbx.Params{
						"collection": ftsrequest.Collection + "_fts",
						"term":       ftsrequest.Term,
					}).
					All(&e)
			} else {
				err = app.Dao().DB().
					NewQuery(
						"SELECT * FROM " +
							cname +
							" WHERE " +
							cname +
							" MATCH {:term} LIMIT 10000").
					Bind(dbx.Params{
						"collection": ftsrequest.Collection + "_fts",
						"term":       "{ " + ftsrequest.Fields + "} : " + ftsrequest.Term,
					}).
					All(&e)
			}

			var unpacked []string
			for _, s := range e {
				unpacked = append(unpacked, s["id"].String)
			}

			query := app.Dao().RecordQuery(collection).AndWhere(dbx.In(
				collection.Name+".id",
				list.ToInterfaceSlice(unpacked)...,
			))

			if ftsrequest.Filter != "" {

				resolver := resolvers.NewRecordFieldResolver(
					app.Dao(),
					collection, // the base collection
					nil,        // no request data
					false,      // allow searching hidden/protected fields like "email"
				)

				expr, err := search.FilterData(ftsrequest.Filter).BuildExpr(resolver)
				if err != nil || expr == nil {
					return apis.NewBadRequestError("invalid or empty filter expression", err)
				}
				query = query.AndWhere(expr)
			}

			if ftsrequest.Sort != "" {
				query = query.OrderBy(ftsrequest.Sort)
			}

			records := make([]*models.Record, 0, len(unpacked))

			if err := query.All(&records); err != nil {
				return apis.NewNotFoundError("The collection or term could not be found!:", err)
			}

			f := strings.Split(ftsrequest.Expand, ",")
			if len(f) > 0 {
				apis.EnrichRecords(c, app.Dao(), records, f...)
			}

			return c.JSON(http.StatusOK, records)
		})

		return nil
	})
}
