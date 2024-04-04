package repo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"shadow_release/tool"
)


func Server() {
	e := echo.New()

	t := tool.New(tool.Config{1, "original"})

	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		t.Track(ctx.Path(), resBody)
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
