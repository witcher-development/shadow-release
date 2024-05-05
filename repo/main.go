package main

import (
	"bytes"
	"net/http"
	sr "shadow_release"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


var SHADOW_URL = "http://localhost:1235"


func server() {
	e := echo.New()

	t := sr.New(sr.Config{1, "original"})

	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		go func(base string, path string) {
			syncKey := uuid.New().String()
			t.Track(ctx.Path(), ctx.Request().Method, reqBody, resBody, syncKey)

			method := ctx.Request().Method
			var err error 
			requestPath := strings.Join([]string{base, path, "?synckey=", syncKey}, "")
			if method == "POST" {
				_, err = http.Post(
					requestPath,
					ctx.Request().Header["Content-Type"][0],
					bytes.NewReader(reqBody),
				)
			} else if method == "GET" {
				_, err = http.Get(requestPath)
			}
			if err != nil {
				panic(err)
			}
		}(SHADOW_URL, ctx.Path())
		// fmt.Println(ctx.Request().Method)
		// fmt.Println("reqBody")
		// fmt.Println(reqBody)
		// fmt.Println("resBody")
		// fmt.Println(resBody)
	}))

	e.GET("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1234"))
}

func main() {
	server()
}
