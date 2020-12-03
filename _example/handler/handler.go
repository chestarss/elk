// Code generated by entc, DO NOT EDIT.

package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/liip/sheriff"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/render"
	"github.com/sirupsen/logrus"

	"github.com/masseelch/elk/_example/schema"
)

// Shared handler.
type handler struct {
	*chi.Mux

	client    *ent.Client
	validator *validator.Validate
	logger    *logrus.Logger
}

// The OwnerHandler.
type OwnerHandler struct {
	*handler
}

// Create a new OwnerHandler
func NewOwnerHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *OwnerHandler {
	h := &OwnerHandler{
		&handler{
			Mux:       chi.NewRouter(),
			client:    c,
			validator: v,
			logger:    log,
		},
	}

	h.Post("/", h.Create)
	h.Get("/{id:\\d+}", h.Read)
	h.Patch("/{id:\\d+}", h.Update)

	h.Get("/", h.List)

	h.Get("/{id:\\d+}/pets", h.Pets)

	return h
}

// struct to bind the post body to.
type ownerCreateRequest struct {
	Name string `json:"name,omitempty"`
	Pets []int  `json:"pets"`
}

// This function creates a new Owner model and stores it in the database.
func (h OwnerHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get the post data.
	d := ownerCreateRequest{} // todo - allow form-url-encdoded/xml/protobuf data.
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		h.logger.WithError(err).Error("error decoding json")
		render.BadRequest(w, r, "invalid json string")
		return
	}

	// Validate the data.
	if err := h.validator.Struct(d); err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			h.logger.WithError(err).Error("error validating request data")
			render.InternalServerError(w, r, nil)
			return
		}

		h.logger.WithError(err).Info("validation failed")
		render.BadRequest(w, r, err)
		return
	}

	// Save the data.
	b := h.client.Owner.Create().
		SetName(d.Name).
		AddPetIDs(d.Pets...)

	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error saving Owner")
		render.InternalServerError(w, r, nil)
		return
	}

	// Read new entry.
	q := h.client.Owner.Query().Where(owner.ID(e.ID))

	e1, err := q.Only(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error reading Owner")
		render.InternalServerError(w, r, nil)
		return
	}

	// Serialize the data.
	j, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"owner:read"}}, e1)
	if err != nil {
		h.logger.WithError(err).WithField("Owner.id", e.ID).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("owner", e.ID).Info("owner rendered")
	render.OK(w, r, j)
}

// This function fetches the Owner model identified by a give url-parameter from
// database and returns it to the client.
func (h OwnerHandler) Read(w http.ResponseWriter, r *http.Request) {
	id, err := h.urlParamInt(w, r, "id")
	if err != nil {
		return
	}

	q := h.client.Owner.Query().Where(owner.ID(id))

	q.WithPets()

	e, err := q.Only(r.Context())

	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			h.logger.WithError(err).WithField("Owner.id", id).Debug("job not found")
			render.NotFound(w, r, err)
			return
		case *ent.NotSingularError:
			h.logger.WithError(err).WithField("Owner.id", id).Error("duplicate entry for id")
			render.InternalServerError(w, r, nil)
			return
		default:
			h.logger.WithError(err).WithField("Owner.id", id).Error("error fetching node from db")
			render.InternalServerError(w, r, nil)
			return
		}
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"owner:read", "pet:list"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Owner.id", id).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("owner", e.ID).Info("owner rendered")
	render.OK(w, r, d)
}

// struct to bind the post body to.
type ownerUpdateRequest struct {
	Name *string `json:"name"`
	Pets []int   `json:"pets"`
}

// This function updates a given Owner model and saves the changes in the database.
func (h OwnerHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.urlParamInt(w, r, "id")
	if err != nil {
		return
	}

	// Get the post data.
	d := ownerUpdateRequest{} // todo - allow form-url-encoded/xml/protobuf data.
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		h.logger.WithError(err).Error("error decoding json")
		render.BadRequest(w, r, "invalid json string")
		return
	}

	// Validate the data.
	if err := h.validator.Struct(d); err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			h.logger.WithError(err).Error("error validating request data")
			render.InternalServerError(w, r, nil)
			return
		}

		h.logger.WithError(err).Info("validation failed")
		render.BadRequest(w, r, err)
		return
	}

	// Save the data.
	b := h.client.Owner.UpdateOneID(id)
	if d.Name != nil {
		b.SetName(d.Name)
	}
	if d.Pets != nil {
		b.AddPetIDs(d.Pets...)
	}

	// Save in database.
	e, err := b.Save(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error saving Owner")
		render.InternalServerError(w, r, nil)
		return
	}

	// Serialize the data.
	j, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"owner:read"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Owner.id", e.ID).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("owner", e.ID).Info("owner rendered")
	render.OK(w, r, j)
}

// This function queries for Owner models. Can be filtered by query parameters.
func (h OwnerHandler) List(w http.ResponseWriter, r *http.Request) {
	q := h.client.Owner.Query()

	if r.URL.Query().Get("order") == "" {
		q.Order(ent.Desc("name"))
	}

	// Pagination
	page, itemsPerPage, err := h.paginationInfo(w, r)
	if err != nil {
		return
	}

	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	// Use the query parameters to filter the query. todo - nested filter?
	if f := r.URL.Query().Get("name"); f != "" {
		q.Where(owner.Name(f))
	}

	es, err := q.All(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error querying database") // todo - better error
		render.InternalServerError(w, r, nil)
		return
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"owner:read", "pet:list"}}, es)
	if err != nil {
		h.logger.WithError(err).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("amount", len(es)).Info("owner rendered")
	render.OK(w, r, d)
}

func (h OwnerHandler) Pets(w http.ResponseWriter, r *http.Request) {
	id, err := h.urlParamInt(w, r, "id")
	if err != nil {
		return
	}

	q := h.client.Owner.Query().Where(owner.ID(id)).QueryPets()

	if r.URL.Query().Get("order") == "" {
		q.Order(ent.Asc("name"), ent.Desc("id"))
	}

	// Eager load edges.
	q

	// Pagination
	page, itemsPerPage, err := h.paginationInfo(w, r)
	if err != nil {
		return
	}

	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	// Use the query parameters to filter the query. todo - nested filter?
	if f := r.URL.Query().Get("name"); f != "" {
		q.Where(pet.Name(f))
	}

	if f := r.URL.Query().Get("age"); f != "" {
		i, err := strconv.Atoi(f)
		if err != nil {
			h.logger.WithError(err).WithField("age", f).Debug("could not parse query parameter")
			render.BadRequest(w, r, "'age' must be an integer")
			return
		}
		q.Where(pet.Age(i))
	}

	if f := r.URL.Query().Get("color"); f != "" {
		// todo
	}

	es, err := q.All(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error querying database") // todo - better error
		render.InternalServerError(w, r, nil)
		return
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"pet:list", "owner:list"}}, es)
	if err != nil {
		h.logger.WithError(err).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("amount", len(es)).Info("pet rendered")
	render.OK(w, r, d)

}

// The PetHandler.
type PetHandler struct {
	*handler
}

// Create a new PetHandler
func NewPetHandler(c *ent.Client, v *validator.Validate, log *logrus.Logger) *PetHandler {
	h := &PetHandler{
		&handler{
			Mux:       chi.NewRouter(),
			client:    c,
			validator: v,
			logger:    log,
		},
	}

	h.Post("/", h.Create)
	h.Get("/{id:\\d+}", h.Read)
	h.Patch("/{id:\\d+}", h.Update)

	h.Get("/", h.List)

	h.Get("/{id:\\d+}/owner", h.Owner)

	return h
}

// struct to bind the post body to.
type petCreateRequest struct {
	Name  string       `json:"name,omitempty"`
	Age   int          `json:"age,omitempty"`
	Color schema.Color `json:"color,omitempty"`
	Owner int          `json:"owner"`
}

// This function creates a new Pet model and stores it in the database.
func (h PetHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get the post data.
	d := petCreateRequest{} // todo - allow form-url-encdoded/xml/protobuf data.
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		h.logger.WithError(err).Error("error decoding json")
		render.BadRequest(w, r, "invalid json string")
		return
	}

	// Validate the data.
	if err := h.validator.Struct(d); err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			h.logger.WithError(err).Error("error validating request data")
			render.InternalServerError(w, r, nil)
			return
		}

		h.logger.WithError(err).Info("validation failed")
		render.BadRequest(w, r, err)
		return
	}

	// Save the data.
	b := h.client.Pet.Create().
		SetName(d.Name).
		SetAge(d.Age).
		SetColor(d.Color).
		SetOwnerID(d.Owner)

	// Store in database.
	e, err := b.Save(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error saving Pet")
		render.InternalServerError(w, r, nil)
		return
	}

	// Read new entry.
	q := h.client.Pet.Query().Where(pet.ID(e.ID))

	e1, err := q.Only(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error reading Pet")
		render.InternalServerError(w, r, nil)
		return
	}

	// Serialize the data.
	j, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"pet:read"}}, e1)
	if err != nil {
		h.logger.WithError(err).WithField("Pet.id", e.ID).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("pet", e.ID).Info("pet rendered")
	render.OK(w, r, j)
}

// This function fetches the Pet model identified by a give url-parameter from
// database and returns it to the client.
func (h PetHandler) Read(w http.ResponseWriter, r *http.Request) {
	id, err := h.urlParamInt(w, r, "id")
	if err != nil {
		return
	}

	q := h.client.Pet.Query().Where(pet.ID(id))

	e, err := q.Only(r.Context())

	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			h.logger.WithError(err).WithField("Pet.id", id).Debug("job not found")
			render.NotFound(w, r, err)
			return
		case *ent.NotSingularError:
			h.logger.WithError(err).WithField("Pet.id", id).Error("duplicate entry for id")
			render.InternalServerError(w, r, nil)
			return
		default:
			h.logger.WithError(err).WithField("Pet.id", id).Error("error fetching node from db")
			render.InternalServerError(w, r, nil)
			return
		}
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"pet:read"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Pet.id", id).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("pet", e.ID).Info("pet rendered")
	render.OK(w, r, d)
}

// struct to bind the post body to.
type petUpdateRequest struct {
	Name  *string       `json:"name"`
	Age   int           `json:"age"`
	Color *schema.Color `json:"color"`
	Owner *int          `json:"owner"`
}

// This function updates a given Pet model and saves the changes in the database.
func (h PetHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := h.urlParamInt(w, r, "id")
	if err != nil {
		return
	}

	// Get the post data.
	d := petUpdateRequest{} // todo - allow form-url-encoded/xml/protobuf data.
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		h.logger.WithError(err).Error("error decoding json")
		render.BadRequest(w, r, "invalid json string")
		return
	}

	// Validate the data.
	if err := h.validator.Struct(d); err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			h.logger.WithError(err).Error("error validating request data")
			render.InternalServerError(w, r, nil)
			return
		}

		h.logger.WithError(err).Info("validation failed")
		render.BadRequest(w, r, err)
		return
	}

	// Save the data.
	b := h.client.Pet.UpdateOneID(id)
	if d.Name != nil {
		b.SetName(d.Name)
	}
	if d.Age != nil {
		b.SetAge(d.Age)
	}
	if d.Color != nil {
		b.SetColor(d.Color)
	}
	if d.Owner != nil {
		b.SetOwnerID(d.Owner)
	}

	// Save in database.
	e, err := b.Save(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error saving Pet")
		render.InternalServerError(w, r, nil)
		return
	}

	// Serialize the data.
	j, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"pet:read"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Pet.id", e.ID).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("pet", e.ID).Info("pet rendered")
	render.OK(w, r, j)
}

// This function queries for Pet models. Can be filtered by query parameters.
func (h PetHandler) List(w http.ResponseWriter, r *http.Request) {
	q := h.client.Pet.Query()

	// Eager load edges.
	q.WithOwner()

	// Pagination
	page, itemsPerPage, err := h.paginationInfo(w, r)
	if err != nil {
		return
	}

	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	// Use the query parameters to filter the query. todo - nested filter?
	if f := r.URL.Query().Get("name"); f != "" {
		q.Where(pet.Name(f))
	}

	if f := r.URL.Query().Get("age"); f != "" {
		i, err := strconv.Atoi(f)
		if err != nil {
			h.logger.WithError(err).WithField("age", f).Debug("could not parse query parameter")
			render.BadRequest(w, r, "'age' must be an integer")
			return
		}
		q.Where(pet.Age(i))
	}

	if f := r.URL.Query().Get("color"); f != "" {
		// todo
	}

	es, err := q.All(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("error querying database") // todo - better error
		render.InternalServerError(w, r, nil)
		return
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"pet:list"}}, es)
	if err != nil {
		h.logger.WithError(err).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("amount", len(es)).Info("pet rendered")
	render.OK(w, r, d)
}

func (h PetHandler) Owner(w http.ResponseWriter, r *http.Request) {
	id, err := h.urlParamInt(w, r, "id")
	if err != nil {
		return
	}

	q := h.client.Pet.Query().Where(pet.ID(id)).QueryOwner()

	q.WithPets()

	e, err := q.Only(r.Context())

	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			h.logger.WithError(err).WithField("Owner.id", id).Debug("job not found")
			render.NotFound(w, r, err)
			return
		case *ent.NotSingularError:
			h.logger.WithError(err).WithField("Owner.id", id).Error("duplicate entry for id")
			render.InternalServerError(w, r, nil)
			return
		default:
			h.logger.WithError(err).WithField("Owner.id", id).Error("error fetching node from db")
			render.InternalServerError(w, r, nil)
			return
		}
	}

	d, err := sheriff.Marshal(&sheriff.Options{Groups: []string{"owner:read", "pet:list"}}, e)
	if err != nil {
		h.logger.WithError(err).WithField("Owner.id", id).Error("serialization error")
		render.InternalServerError(w, r, nil)
		return
	}

	h.logger.WithField("owner", e.ID).Info("owner rendered")
	render.OK(w, r, d)

}

func (h handler) urlParamString(w http.ResponseWriter, r *http.Request, param string) (id string, err error) {
	id = chi.URLParam(r, param)
	if id == "" {
		err = errors.New("empty url param")
		h.logger.WithField("param", param).Info("empty url param")
		render.BadRequest(w, r, param+" cannot be ''")
	}

	return
}
func (h handler) urlParamInt(w http.ResponseWriter, r *http.Request, param string) (id int, err error) {
	p := chi.URLParam(r, param)
	if p == "" {
		err = errors.New("empty url param")
		h.logger.WithField("param", param).Info("empty url param")
		render.BadRequest(w, r, param+" cannot be ''")
		return
	}

	id, err = strconv.Atoi(p)
	if err != nil {
		h.logger.WithField(param, p).Info("error parsing url parameter")
		render.BadRequest(w, r, param+" must be a positive integer greater zero")
		return
	}

	return
}

func (h handler) urlParamTime(w http.ResponseWriter, r *http.Request, param string) (date time.Time, err error) {
	p := chi.URLParam(r, param)
	if p == "" {
		h.logger.WithField("param", param).Info("empty url param")
		render.BadRequest(w, r, param+" cannot be ''")
		return
	}

	date, err = time.Parse("2006-01-02", p)
	if err != nil {
		h.logger.WithField(param, p).Info("error parsing url parameter")
		render.BadRequest(w, r, param+" must be a valid date in yyyy-mm-dd format")
		return
	}

	return
}

func (h handler) paginationInfo(w http.ResponseWriter, r *http.Request) (page int, itemsPerPage int, err error) {
	page = 1
	itemsPerPage = 30

	if d := r.URL.Query().Get("itemsPerPage"); d != "" {
		itemsPerPage, err = strconv.Atoi(d)
		if err != nil {
			h.logger.WithField("itemsPerPage", d).Info("error parsing query parameter 'itemsPerPage'")
			render.BadRequest(w, r, "itemsPerPage must be a positive integer greater zero")
			return
		}
	}

	if d := r.URL.Query().Get("page"); d != "" {
		page, err = strconv.Atoi(d)
		if err != nil {
			h.logger.WithField("page", d).Info("error parsing query parameter 'page'")
			render.BadRequest(w, r, "page must be a positive integer greater zero")
			return
		}
	}

	return
}
