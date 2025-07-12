package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/dunamismax/go-modern-scaffold/internal/cache"
	"github.com/dunamismax/go-modern-scaffold/internal/config"
	"github.com/dunamismax/go-modern-scaffold/internal/db"
	"github.com/dunamismax/go-modern-scaffold/internal/web"
	"github.com/go-playground/validator/v10"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CustomValidator holds the validator instance.
type CustomValidator struct {
	validator *validator.Validate
}

// Validate implements the echo.Validator interface.
func (v *CustomValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	// Setup structured logging
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Setup database connection
	dbConn, err := sql.Open("sqlite3", cfg.DBURL)
	if err != nil {
		log.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// Create a new sqlc querier
	queries := db.New(dbConn)

	// Create a new cache
	appCache, err := cache.New(&cfg.Cache)
	if err != nil {
		log.Error("failed to create cache", "error", err)
		os.Exit(1)
	}

	// Create Echo app
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Add middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Static files
	e.Static("/css", "./public/css")
	e.Static("/js", "./public/js")
	e.Static("/assets", "./public/assets")

	// Create web handlers
	webHandlers := web.NewHandlers(queries, appCache)

	// Register routes
	e.GET("/", webHandlers.RenderIndex)
	e.POST("/messages", webHandlers.CreateMessage)
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// Start server
	listenAddr := fmt.Sprintf(":%d", cfg.HTTPPort)
	log.Info("starting server", "address", listenAddr)
	if err := e.Start(listenAddr); err != nil {
		log.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
