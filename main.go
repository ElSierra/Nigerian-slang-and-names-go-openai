package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elsierra/go-echo-zik/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	e := echo.New()

	
	API_KEY := os.Getenv("OPENAI_API_KEY")
	fmt.Println("apikey", API_KEY)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Init(e)
	e.Start(":8080")
}
