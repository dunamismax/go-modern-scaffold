package main

import (
	"net/http"

	"github.com/dunamismax/go-modern-scaffold/internal/database"
	"github.com/dunamismax/go-modern-scaffold/internal/web/pages"
	"github.com/labstack/echo/v4"
)

// Handlers holds the database connection.
type Handlers struct {
	db database.DB
}

// NewHandlers creates a new Handlers instance.
func NewHandlers(db *database.DB) *Handlers {
	return &Handlers{db: *db}
}

// Register registers the handlers with the Echo instance.
func (h *Handlers) Register(e *echo.Echo) {
	e.GET("/", h.GetNotes)
	e.POST("/notes", h.CreateNote)
}

// GetNotes handles the request to get all notes.
func (h *Handlers) GetNotes(c echo.Context) error {
	notes, err := h.db.ListNotes(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get notes")
	}

	return pages.Index("Notes", notes).Render(c.Request().Context(), c.Response().Writer)
}

// CreateNote handles the request to create a new note.
func (h *Handlers) CreateNote(c echo.Context) error {
	content := c.FormValue("content")
	if content == "" {
		return c.String(http.StatusBadRequest, "content cannot be empty")
	}

	_, err := h.db.CreateNote(c.Request().Context(), content)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to create note")
	}

	notes, err := h.db.ListNotes(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to get notes")
	}

	return pages.Index("Notes", notes).Render(c.Request().Context(), c.Response().Writer)
}