// Code generated by entc, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/chestarss/elk/internal/uuid/ent"
	"go.uber.org/zap"
)

// NewHandler returns a ready to use handler with all generated endpoints mounted.
func NewHandler(c *ent.Client, l *zap.Logger) chi.Router {
	r := chi.NewRouter()
	MountRoutes(c, l, r)
	return r
}

// MountRoutes mounts all generated routes on the given router.
func MountRoutes(c *ent.Client, l *zap.Logger, r chi.Router) {
	NewUserHandler(c, l).MountRoutes(r)
}

// UserHandler handles http crud operations on ent.User.
type UserHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewUserHandler(c *ent.Client, l *zap.Logger) *UserHandler {
	return &UserHandler{
		client: c,
		log:    l.With(zap.String("handler", "UserHandler")),
	}
}
func (h *UserHandler) MountCreateRoute(r chi.Router) *UserHandler {
	r.Post("/users", h.Create)
	return h
}
func (h *UserHandler) MountReadRoute(r chi.Router) *UserHandler {
	r.Get("/users/{id}", h.Read)
	return h
}
func (h *UserHandler) MountUpdateRoute(r chi.Router) *UserHandler {
	r.Patch("/users/{id}", h.Update)
	return h
}
func (h *UserHandler) MountDeleteRoute(r chi.Router) *UserHandler {
	r.Delete("/users/{id}", h.Delete)
	return h
}
func (h *UserHandler) MountListRoute(r chi.Router) *UserHandler {
	r.Get("/users", h.List)
	return h
}
func (h *UserHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r)
}

func stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func zapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}
