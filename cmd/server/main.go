package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/dunamismax/go-modern-scaffold/internal/cache"
	"github.com/dunamismax/go-modern-scaffold/internal/config"
	"github.com/dunamismax/go-modern-scaffold/internal/db"
	"github.com/dunamismax/go-modern-scaffold/internal/web"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
)

// XValidator provides a custom validator for Fiber.
type XValidator struct {
	validator *validator.Validate
}

// Validate performs validation on a struct.
func (v *XValidator) Validate(data interface{}) error {
	return v.validator.Struct(data)
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

	// Setup database connection pool
	pool, err := pgxpool.New(context.Background(), cfg.DBURL)
	if err != nil {
		log.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Error("failed to ping database", "error", err)
		os.Exit(1)
	}
	log.Info("successfully connected to the database")

	// Create a new sqlc querier
	queries := db.New(pool)

	// Create a new cache
	appCache, err := cache.New(&cfg.Cache)
	if err != nil {
		log.Error("failed to create cache", "error", err)
		os.Exit(1)
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "go-modern-scaffold",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	})

	// Set custom validator
	app.Validator = &XValidator{validator: validator.New()}

	// Add middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Static files
	app.Static("/css", "./public/css")
	app.Static("/js", "./public/js")
	app.Static("/assets", "./public/assets")

	// Create web handlers
	webHandlers := web.NewHandlers(queries, appCache)

	// Register routes
	app.Get("/", webHandlers.RenderIndex)
	app.Post("/messages", webHandlers.CreateMessage)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Start server
	listenAddr := fmt.Sprintf(":%d", cfg.HTTPPort)
	log.Info("starting server", "address", listenAddr)
	if err := app.Listen(listenAddr); err != nil {
		log.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}