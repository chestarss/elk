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

// Fridge fetches the ent.fridge attached to the ent.Compartment
// identified by a given url-parameter from the database and renders it to the client.
func (h CompartmentHandler) Fridge(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Fridge"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the fridge attached to this compartment
	q := h.client.Compartment.Query().Where(compartment.ID(id)).QueryFridge()
	e, err := q.Only(r.Context())
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
	l.Info("fridge rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewFridge2716213877View(e), w)
}

// Contents fetches the ent.contents attached to the ent.Compartment
// identified by a given url-parameter from the database and renders it to the client.
func (h CompartmentHandler) Contents(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Contents"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the contents attached to this compartment
	q := h.client.Compartment.Query().Where(compartment.ID(id)).QueryContents()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching items from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("items rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewItem1509516544Views(es), w)
}

// Compartments fetches the ent.compartments attached to the ent.Fridge
// identified by a given url-parameter from the database and renders it to the client.
func (h FridgeHandler) Compartments(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Compartments"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the compartments attached to this fridge
	q := h.client.Fridge.Query().Where(fridge.ID(id)).QueryCompartments()
	page := 1
	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'page'", zap.String("page", d), zap.Error(err))
			BadRequest(w, "page must be an integer greater zero")
			return
		}
	}
	itemsPerPage := 30
	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			l.Info("error parsing query parameter 'itemsPerPage'", zap.String("itemsPerPage", d), zap.Error(err))
			BadRequest(w, "itemsPerPage must be an integer greater zero")
			return
		}
	}
	es, err := q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage).All(r.Context())
	if err != nil {
		l.Error("error fetching compartments from db", zap.Error(err))
		InternalServerError(w, nil)
		return
	}
	l.Info("compartments rendered", zap.Int("amount", len(es)))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment1151786357Views(es), w)
}

// Compartment fetches the ent.compartment attached to the ent.Item
// identified by a given url-parameter from the database and renders it to the client.
func (h ItemHandler) Compartment(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Compartment"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Create the query to fetch the compartment attached to this item
	q := h.client.Item.Query().Where(item.ID(id)).QueryCompartment()
	e, err := q.Only(r.Context())
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
	l.Info("compartment rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewCompartment1151786357View(e), w)
}
