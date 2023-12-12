package routes

import (
	"github.com/elsierra/go-echo-zik/handlers"
	"github.com/labstack/echo/v4"
	
)

func Init(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.GET("/", handlers.HomeHandler)
	v1.POST("/search", handlers.PostHandler)
}
