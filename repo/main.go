package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	e := echo.New()
	e.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		fmt.Println(ctx.Path())
		fmt.Println("reqBody")
		fmt.Println(reqBody)
		fmt.Println("resBody")
		fmt.Println(resBody)
	}))

	e.GET("/hi", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
