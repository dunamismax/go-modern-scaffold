package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dunamismax/go-modern-scaffold/internal/config"
	"github.com/dunamismax/go-modern-scaffold/internal/database"
	customValidator "github.com/dunamismax/go-modern-scaffold/internal/validator"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found")
	}

	// Configuration
	cfg, err := config.New()
	if err != nil {
		slog.Error("configuration error", "error", err)
		os.Exit(1)
	}

	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Dependencies
	db, err := database.New(cfg.Database.DSN)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Echo instance
	e := echo.New()
	e.Validator = customValidator.NewCustomValidator(validator.New())

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Handlers
	handlers := NewHandlers(db)
	handlers.Register(e)

	// Static files
	e.Static("/css", "public/css")

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	slog.Info("server shutdown complete")
}