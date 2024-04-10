package repo

import (
	"fmt"
	"net/http"
	"shadow_release/tool"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


var SHADOW_URL = "http://localhost:1235"


func Server() {
	e := echo.New()

	t := tool.New(tool.Config{1, "original"})

	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		go func(base string, path string) {
			t.Track(ctx.Path(), reqBody, resBody)
			_, err := http.Get(strings.Join([]string{base, path}, ""))
			fmt.Println(err)
		}(SHADOW_URL, ctx.Path())
		// fmt.Println(ctx.Path())
		// fmt.Println("reqBody")
		// fmt.Println(reqBody)
		// fmt.Println("resBody")
		// fmt.Println(resBody)
	}))

	e.GET("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
