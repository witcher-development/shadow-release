package main

import (
	"net/http"
	sr "shadow_release"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func server() {
	e := echo.New()

	t := sr.New(sr.Config{1, "shadow"})

	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		t.Track(ctx.Path(), ctx.Request().Method, reqBody, resBody, ctx.QueryParam("synckey"))
	}))

	e.GET("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.POST("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1235"))
}

func main() {
	server()
}
