// Code generated by entc, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/chestarss/elk/internal/simple/ent"
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
	NewCategoryHandler(c, l).MountRoutes(r)
	NewCollarHandler(c, l).MountRoutes(r)
	NewMediaHandler(c, l).MountRoutes(r)
	NewOwnerHandler(c, l).MountRoutes(r)
	NewPetHandler(c, l).MountRoutes(r)
}

// CategoryHandler handles http crud operations on ent.Category.
type CategoryHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewCategoryHandler(c *ent.Client, l *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		client: c,
		log:    l.With(zap.String("handler", "CategoryHandler")),
	}
}
func (h *CategoryHandler) MountCreateRoute(r chi.Router) *CategoryHandler {
	r.Post("/categories", h.Create)
	return h
}
func (h *CategoryHandler) MountReadRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}", h.Read)
	return h
}
func (h *CategoryHandler) MountUpdateRoute(r chi.Router) *CategoryHandler {
	r.Patch("/categories/{id}", h.Update)
	return h
}
func (h *CategoryHandler) MountDeleteRoute(r chi.Router) *CategoryHandler {
	r.Delete("/categories/{id}", h.Delete)
	return h
}
func (h *CategoryHandler) MountListRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories", h.List)
	return h
}
func (h *CategoryHandler) MountPetsRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}/pets", h.Pets)
	return h
}
func (h *CategoryHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountPetsRoute(r)
}

// CollarHandler handles http crud operations on ent.Collar.
type CollarHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewCollarHandler(c *ent.Client, l *zap.Logger) *CollarHandler {
	return &CollarHandler{
		client: c,
		log:    l.With(zap.String("handler", "CollarHandler")),
	}
}
func (h *CollarHandler) MountCreateRoute(r chi.Router) *CollarHandler {
	r.Post("/collars", h.Create)
	return h
}
func (h *CollarHandler) MountReadRoute(r chi.Router) *CollarHandler {
	r.Get("/collars/{id}", h.Read)
	return h
}
func (h *CollarHandler) MountUpdateRoute(r chi.Router) *CollarHandler {
	r.Patch("/collars/{id}", h.Update)
	return h
}
func (h *CollarHandler) MountDeleteRoute(r chi.Router) *CollarHandler {
	r.Delete("/collars/{id}", h.Delete)
	return h
}
func (h *CollarHandler) MountListRoute(r chi.Router) *CollarHandler {
	r.Get("/collars", h.List)
	return h
}
func (h *CollarHandler) MountPetRoute(r chi.Router) *CollarHandler {
	r.Get("/collars/{id}/pet", h.Pet)
	return h
}
func (h *CollarHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountPetRoute(r)
}

// MediaHandler handles http crud operations on ent.Media.
type MediaHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewMediaHandler(c *ent.Client, l *zap.Logger) *MediaHandler {
	return &MediaHandler{
		client: c,
		log:    l.With(zap.String("handler", "MediaHandler")),
	}
}
func (h *MediaHandler) MountCreateRoute(r chi.Router) *MediaHandler {
	r.Post("/media", h.Create)
	return h
}
func (h *MediaHandler) MountReadRoute(r chi.Router) *MediaHandler {
	r.Get("/media/{id}", h.Read)
	return h
}
func (h *MediaHandler) MountUpdateRoute(r chi.Router) *MediaHandler {
	r.Patch("/media/{id}", h.Update)
	return h
}
func (h *MediaHandler) MountDeleteRoute(r chi.Router) *MediaHandler {
	r.Delete("/media/{id}", h.Delete)
	return h
}
func (h *MediaHandler) MountListRoute(r chi.Router) *MediaHandler {
	r.Get("/media", h.List)
	return h
}
func (h *MediaHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r)
}

// OwnerHandler handles http crud operations on ent.Owner.
type OwnerHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewOwnerHandler(c *ent.Client, l *zap.Logger) *OwnerHandler {
	return &OwnerHandler{
		client: c,
		log:    l.With(zap.String("handler", "OwnerHandler")),
	}
}
func (h *OwnerHandler) MountCreateRoute(r chi.Router) *OwnerHandler {
	r.Post("/owners", h.Create)
	return h
}
func (h *OwnerHandler) MountReadRoute(r chi.Router) *OwnerHandler {
	r.Get("/owners/{id}", h.Read)
	return h
}
func (h *OwnerHandler) MountUpdateRoute(r chi.Router) *OwnerHandler {
	r.Patch("/owners/{id}", h.Update)
	return h
}
func (h *OwnerHandler) MountDeleteRoute(r chi.Router) *OwnerHandler {
	r.Delete("/owners/{id}", h.Delete)
	return h
}
func (h *OwnerHandler) MountListRoute(r chi.Router) *OwnerHandler {
	r.Get("/owners", h.List)
	return h
}
func (h *OwnerHandler) MountPetsRoute(r chi.Router) *OwnerHandler {
	r.Get("/owners/{id}/pets", h.Pets)
	return h
}
func (h *OwnerHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountPetsRoute(r)
}

// PetHandler handles http crud operations on ent.Pet.
type PetHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewPetHandler(c *ent.Client, l *zap.Logger) *PetHandler {
	return &PetHandler{
		client: c,
		log:    l.With(zap.String("handler", "PetHandler")),
	}
}
func (h *PetHandler) MountCreateRoute(r chi.Router) *PetHandler {
	r.Post("/pets", h.Create)
	return h
}
func (h *PetHandler) MountReadRoute(r chi.Router) *PetHandler {
	r.Get("/pets/{id}", h.Read)
	return h
}
func (h *PetHandler) MountUpdateRoute(r chi.Router) *PetHandler {
	r.Patch("/pets/{id}", h.Update)
	return h
}
func (h *PetHandler) MountDeleteRoute(r chi.Router) *PetHandler {
	r.Delete("/pets/{id}", h.Delete)
	return h
}
func (h *PetHandler) MountListRoute(r chi.Router) *PetHandler {
	r.Get("/pets", h.List)
	return h
}
func (h *PetHandler) MountCollarRoute(r chi.Router) *PetHandler {
	r.Get("/pets/{id}/collar", h.Collar)
	return h
}
func (h *PetHandler) MountCategoriesRoute(r chi.Router) *PetHandler {
	r.Get("/pets/{id}/categories", h.Categories)
	return h
}
func (h *PetHandler) MountOwnerRoute(r chi.Router) *PetHandler {
	r.Get("/pets/{id}/owner", h.Owner)
	return h
}
func (h *PetHandler) MountFriendsRoute(r chi.Router) *PetHandler {
	r.Get("/pets/{id}/friends", h.Friends)
	return h
}
func (h *PetHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r).MountCollarRoute(r).MountCategoriesRoute(r).MountOwnerRoute(r).MountFriendsRoute(r)
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
