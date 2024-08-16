package queries

import (
	"embed"
	"io/fs"
)

//go:embed Musenalm/*.sql
var musenalmq embed.FS

var MusenalmQueries = MustSubFS(musenalmq, "Musenalm")

func MustSubFS(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)

	if err != nil {
		panic("Could not create SubFS for " + dir)
	}

	return sub
}
