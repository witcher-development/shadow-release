package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"shadow-release/db"

	_ "github.com/mattn/go-sqlite3"
)


type Tool struct {
	key string
	version string
}

func (s *Tool) Track(requestId string, path string, body []byte) {
	// add to DB
}

type Config struct {
	Key string
	Version string
}

func New(config Config) (s *Tool) {
	// get or create app
	s = &Tool{
		key: config.Key,
		version: config.Version,
	}
	return s
}


// func (s *Tool) 

// func parseConfig() Config {
// 	configContent, err := os.ReadFile("./config.json")
// 	if err != nil {
// 		panic(err)
// 	}
//
//
// 	var config Config
//
// 	if err := json.Unmarshal(configContent, &config); err != nil {
// 		panic(err)
// 	}
//
// 	return config
// }
//
// func fetchRepo(url string) {
// 	cmd := exec.Command("git", "clone", url, "tmp/repo")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	cmd.Run()
// }

//go:embed db/schema.sql
var ddl string

func main() {
	ctx := context.Background()

	sqlite, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	defer sqlite.Close()

	if _, err := sqlite.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	queries := db.New(sqlite)

	_, err = queries.CreateApp(ctx, 1)
	if err != nil {
		fmt.Print(err)
	}

	app, err := queries.GetApp(ctx, 1)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(app)
}



