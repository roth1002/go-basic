package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Post("/users", save)
	e.Run(standard.New(":1323"))
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue()["email"]
	fmt.Print(name, email)
	return c.HTML(http.StatusOK, "<b>Thank you!</b>")
}
