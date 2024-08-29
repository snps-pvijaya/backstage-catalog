package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", helloWorld)
	e.Start(":9000")
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
