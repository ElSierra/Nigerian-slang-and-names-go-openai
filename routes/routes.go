package routes

import (
	"github.com/elsierra/go-echo-zik/handlers"
	"github.com/labstack/echo/v4"
	
)

func Init(e *echo.Echo, apiCfg *handlers.ApiConfig ) {
	v1 := e.Group("/v1")
	v1.GET("/", apiCfg.HomeHandler)
	v1.POST("/search", apiCfg.PostHandler)
	v1.POST("/reSearch", apiCfg.ReSearch)
}
