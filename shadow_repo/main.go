package shadow_repo

import (
	"net/http"
	"shadow_release/tool"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func Server() {
	e := echo.New()

	t := tool.New(tool.Config{1, "shadow"})

	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		t.Track(ctx.Path(), reqBody, resBody)
	}))

	e.GET("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1235"))
}
