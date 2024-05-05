package shadow_release

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"shadow_release/internal/db"
	"shadow_release/internal/views"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)


type Tool struct {
	Key int64 `json:"key"`
	Version int64 `json:"version"`
}

type TrackRequestPayload struct {
	Meta *Tool `json:"meta"`
	Path string
	Method string
	Reqbody []byte
	Resbody []byte
	Synckey string
}

func (s *Tool) Track(path string, method string, reqbody []byte, resbody []byte, syncKey string) {
	body := TrackRequestPayload{
		Meta: s,
		Path: path,
		Method: method,
		Reqbody: reqbody,
		Resbody: resbody,
		Synckey: syncKey,
	}
	body_json, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	_, err = http.Post(
		"http://localhost:3333/track",
		"application/json",
		bytes.NewBuffer(body_json),
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
	config_json, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	response, err := http.Post(
		"http://localhost:3333/init",
		"application/json",
		bytes.NewReader(config_json),
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var body db.Version
	if err = json.Unmarshal(data, &body); err != nil {
		panic(err)
	}

	s = &Tool{
		Key: body.App,
		Version: body.ID,
	}
	return s
}

func StartBackend() {
	e := echo.New()

	e.POST("/init", func(c echo.Context) error {
		var config Config
		if err := c.Bind(&config); err != nil {
			panic(err)
		}
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

		return c.JSON(http.StatusOK, version)
	})

	e.POST("/track", func (c echo.Context) error  {
		var record TrackRequestPayload
		err := json.NewDecoder(c.Request().Body).Decode(&record)
		if err != nil {
			panic(err)
		}
		fmt.Println("recieved", record)
		ctx, queries := db.GetQueries()
		db_record, err := queries.CreateRecord(ctx, db.CreateRecordParams{
			Version: record.Meta.Version,
			Path: record.Path,
			Method: record.Method,
			Reqbody: string(record.Reqbody),
			Resbody: string(record.Resbody),
			Synckey: string(record.Synckey),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(db_record)
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		ctx, queries := db.GetQueries()
		records, err := queries.GetRecords(ctx)
		if err != nil && err != sql.ErrNoRows {
			panic(err)
		}
		versions, err := queries.GetVersions(ctx)
		if err != nil && err != sql.ErrNoRows {
			panic(err)
		}
		return views.Page(records, versions).Render(context.Background(), c.Response().Writer)
	})
	e.Static("/assets", "internal/views/assets")
	// e.File("/", ")

	e.Logger.Fatal(e.Start(":3333"))
}

