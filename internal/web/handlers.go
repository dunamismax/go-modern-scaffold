package web

import (
	"context"
	"log/slog"

	"github.com/a-h/templ"
	"github.com/dunamismax/go-modern-scaffold/internal/cache"
	"github.com/dunamismax/go-modern-scaffold/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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
func (h *Handlers) RenderIndex(c *fiber.Ctx) error {
	// Try to get messages from cache first
	if cachedMessages, found := h.cache.Get(messagesCacheKey); found {
		if messages, ok := cachedMessages.([]db.Message); ok {
			slog.Info("cache hit for messages")
			return adaptor.HTTPHandler(templ.Handler(Index(messages)))(c)
		}
	}

	// If not in cache, get from DB
	slog.Info("cache miss for messages")
	messages, err := h.queries.GetMessages(context.Background())
	if err != nil {
		slog.Error("failed to get messages", "error", err)
		return fiber.ErrInternalServerError
	}

	// Set messages in cache
	h.cache.Set(messagesCacheKey, messages, 1)

	return adaptor.HTTPHandler(templ.Handler(Index(messages)))(c)
}

// CreateMessage handles the creation of a new message.
func (h *Handlers) CreateMessage(c *fiber.Ctx) error {
	body := c.FormValue("body")
	if body == "" {
		return fiber.NewError(fiber.StatusBadRequest, "message body cannot be empty")
	}

	_, err := h.queries.CreateMessage(context.Background(), body)
	if err != nil {
		slog.Error("failed to create message", "error", err)
		return fiber.ErrInternalServerError
	}

	// Invalidate cache
	h.cache.Del(messagesCacheKey)
	slog.Info("cache invalidated for messages")

	messages, err := h.queries.GetMessages(context.Background())
	if err != nil {
		slog.Error("failed to get messages after creating new one", "error", err)
		return fiber.ErrInternalServerError
	}

	// This is where HTMX shines. We just render the component that needs updating.
	return adaptor.HTTPHandler(templ.Handler(MessageList(messages)))(c)
}