package db

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

var queries *Queries

func GetQueries() (context.Context, *Queries) {
	ctx := context.Background()

	if queries != nil {
		return ctx, queries
	}

	sqlite, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}

	if _, err := sqlite.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	queries = New(sqlite)
	return ctx, queries
}
