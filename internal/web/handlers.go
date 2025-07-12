package web

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dunamismax/go-modern-scaffold/internal/cache"
	"github.com/dunamismax/go-modern-scaffold/internal/db"
	"github.com/labstack/echo/v4"
)

const messagesCacheKey = "messages"

// Handlers holds the dependencies for the web handlers.
type Handlers struct {
	queries db.Querier
	cache   *cache.Cache
}

// NewHandlers creates a new Handlers instance.
func NewHandlers(queries db.Querier, cache *cache.Cache) *Handlers {
	return &Handlers{queries: queries, cache: cache}
}

// RenderIndex renders the main index page.
func (h *Handlers) RenderIndex(c echo.Context) error {
	// Try to get messages from cache first
	if cachedMessages, found := h.cache.Get(messagesCacheKey); found {
		if messages, ok := cachedMessages.([]db.Message); ok {
			slog.Info("cache hit for messages")
			return renderComponent(c, Index(messages))
		}
	}

	// If not in cache, get from DB
	slog.Info("cache miss for messages")
	messages, err := h.queries.GetMessages(context.Background())
	if err != nil {
		slog.Error("failed to get messages", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get messages")
	}

	// Set messages in cache
	h.cache.Set(messagesCacheKey, messages, 1)

	return renderComponent(c, Index(messages))
}

// CreateMessage handles the creation of a new message.
func (h *Handlers) CreateMessage(c echo.Context) error {
	body := c.FormValue("body")
	if body == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Message body cannot be empty")
	}

	err := h.queries.CreateMessage(context.Background(), body)
	if err != nil {
		slog.Error("failed to create message", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create message")
	}

	// Invalidate cache
	h.cache.Del(messagesCacheKey)
	slog.Info("cache invalidated for messages")

	messages, err := h.queries.GetMessages(context.Background())
	if err != nil {
		slog.Error("failed to get messages after creating new one", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get messages")
	}

	// This is where HTMX shines. We just render the component that needs updating.
	return renderComponent(c, MessageList(messages))
}

// renderComponent is a helper to render a templ component.
func renderComponent(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}
