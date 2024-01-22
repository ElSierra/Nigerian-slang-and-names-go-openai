package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/elsierra/go-echo-zik/handlers"
	"github.com/elsierra/go-echo-zik/internal/database"
	"github.com/elsierra/go-echo-zik/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)



func main() {
	godotenv.Load()
	e := echo.New()
	config := middleware.RateLimiterConfig{
	
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 5, Burst: 5, ExpiresIn: 24 * time.Hour},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			fmt.Println(id)
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string,err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	
	
	

	API_KEY := os.Getenv("OPENAI_API_KEY")
	dbURL := os.Getenv("DB_URL")
	if API_KEY == "" || dbURL == "" {
		fmt.Println("OPENAI_API_KEY or DATABASE_URL is not set")
		os.Exit(1)
	}
	conn, err := sql.Open("postgres", dbURL)

	queries := database.New(conn)
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		os.Exit(1)
	}

	apiCfg := &handlers.ApiConfig{
		DB: queries,
	}


	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		os.Exit(1)
	}
	fmt.Println("apikey", API_KEY)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.RateLimiterWithConfig(config))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Init(e,apiCfg)
	e.Start(":8080")
}
