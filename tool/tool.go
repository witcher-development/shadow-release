package tool

import (
	"database/sql"
	"fmt"
	"shadow_release/db"

	_ "github.com/mattn/go-sqlite3"
)


type Tool struct {
	key int64
	version int64
}

func (s *Tool) Track(path string, body []byte) {
	ctx, queries := db.GetQueries()
	queries.CreateRecord(ctx, db.CreateRecordParams{
		Version: s.version,
		Path: path,
		Body: string(body),
	})
}

type Config struct {
	Key int64
	Version string
}

func New(config Config) (s *Tool) {
	ctx, queries := db.GetQueries()
	version, err := queries.GetVersion(ctx, config.Version)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if err == sql.ErrNoRows {
		version_new, err := queries.CreateVersion(ctx, db.CreateVersionParams{
			Name: config.Version,
			App: config.Key,
		})
		if err != nil {
			panic(err)
		}
		version = version_new
	}

	fmt.Println(version)

	s = &Tool{
		key: config.Key,
		version: version.ID,
	}
	return s
}