package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Message(c echo.Context) error {
	return c.Render(http.StatusOK, "a.html", map[string]interface{}{})
}
