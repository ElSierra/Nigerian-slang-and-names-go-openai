// handlers/handler.go
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type msg struct {
	Message string `json:"message"`
}

func HomeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, msg{Message: "Hello World"})
}
