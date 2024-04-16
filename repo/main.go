package main

import (
	"bytes"
	"fmt"
	"net/http"
	sr "shadow_release"
	"strings"

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
			t.Track(ctx.Path(), reqBody, resBody)

			method := ctx.Request().Method
			var err error 
			if method == "POST" {
				_, err = http.Post(
					strings.Join([]string{base, path}, ""),
					ctx.Request().Header["Content-Type"][0],
					bytes.NewReader(reqBody),
				)
			} else if method == "GET" {
				_, err = http.Get(strings.Join([]string{base, path}, ""))
			}

			fmt.Println(err)
		}(SHADOW_URL, ctx.Path())
		fmt.Println(ctx.Request().Method)
		fmt.Println("reqBody")
		fmt.Println(reqBody)
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
