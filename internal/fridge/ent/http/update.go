// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"github.com/chestarss/elk/internal/fridge/ent"
	"github.com/chestarss/elk/internal/fridge/ent/compartment"
	"github.com/chestarss/elk/internal/fridge/ent/fridge"
	"github.com/chestarss/elk/internal/fridge/ent/item"
	"go.uber.org/zap"
)

// Update updates a given ent.Compartment and saves the changes to the database.
func (h CompartmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d CompartmentUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Compartment.UpdateOneID(id)
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Fridge != nil {
		b.SetFridgeID(*d.Fridge)

	}
	if d.Contents != nil {
		b.ClearContents().AddContentIDs(d.Contents...)
	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-compartment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Compartment.Query().Where(compartment.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-compartment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("compartment rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment1151786357View(e), w)
}

// Update updates a given ent.Fridge and saves the changes to the database.
func (h FridgeHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d FridgeUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Fridge.UpdateOneID(id)
	if d.Title != nil {
		b.SetTitle(*d.Title)
	}
	if d.Compartments != nil {
		b.ClearCompartments().AddCompartmentIDs(d.Compartments...)
	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-fridge", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Fridge.Query().Where(fridge.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-fridge", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("fridge rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewFridge2716213877View(e), w)
}

// Update updates a given ent.Item and saves the changes to the database.
func (h ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d ItemUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Item.UpdateOneID(id)
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Compartment != nil {
		b.SetCompartmentID(*d.Compartment)

	}
	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-update-item", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Item.Query().Where(item.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could-not-read-item", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("item rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewItem1509516544View(e), w)
}
