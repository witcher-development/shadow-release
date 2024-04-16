package shadow_release

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"shadow_release/internal/db"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)


type Tool struct {
	key int64
	version int64
}

func (s *Tool) Track(path string, reqbody []byte, resbody []byte) {
	body := fmt.Sprintf(`{
		"reqbody": "%v",
		"resbody": "%v"
	}`, reqbody, resbody)
	// fmt.Println(body)
	_, err := http.Post(
		"http://localhost:3333/track",
		"application/json",
		bytes.NewReader([]byte(body)),
	)
	if err != nil {
		panic(err)
	}
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

	s = &Tool{
		key: config.Key,
		version: version.ID,
	}
	return s
}

func StartBackend() {
	e := echo.New()

	e.POST("/track", func (c echo.Context) error  {
		var json map[string]interface{} = map[string]interface{}{}
		if err := c.Bind(&json); err != nil {
			panic(err)
		}
		fmt.Println(json)
		ctx, queries := db.GetQueries()
		queries.CreateRecord(ctx, db.CreateRecordParams{
			Version: 1,
			Path: "",
			Reqbody: string(reqbody),
			Resbody: string(resbody),
		})
		return nil
	})

	e.Logger.Fatal(e.Start(":3333"))
}



