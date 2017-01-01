package main

import (
	"github.com/labstack/echo"
	"github.com/wimpykid26/hermes/api"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/insert_user", api.Insert_user)
	e.GET("/insert_hub", api.Insert_hub)
	e.GET("/insert_event", api.Insert_event)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
