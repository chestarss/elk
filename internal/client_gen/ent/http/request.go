// Code generated by entc, DO NOT EDIT.

package http

import (
	"github.com/google/uuid"
	collar "github.com/chestarss/elk/internal/client_gen/ent/collar"
)

// Payload of a ent.Category create request.
type CategoryCreateRequest struct {
	Name *string  `json:"name"`
	Pets []string `json:"pets"`
}

// Payload of a ent.Category update request.
type CategoryUpdateRequest struct {
	Name *string  `json:"name"`
	Pets []string `json:"pets"`
}

// Payload of a ent.Collar create request.
type CollarCreateRequest struct {
	Color *collar.Color `json:"color"`
	Pet   *string       `json:"pet"`
}

// Payload of a ent.Collar update request.
type CollarUpdateRequest struct {
	Color *collar.Color `json:"color"`
	Pet   *string       `json:"pet"`
}

// Payload of a ent.Owner create request.
type OwnerCreateRequest struct {
	Name *string  `json:"name"`
	Age  *int     `json:"age"`
	Pets []string `json:"pets"`
}

// Payload of a ent.Owner update request.
type OwnerUpdateRequest struct {
	Name *string  `json:"name"`
	Age  *int     `json:"age"`
	Pets []string `json:"pets"`
}

// Payload of a ent.Pet create request.
type PetCreateRequest struct {
	Name       *string    `json:"name"`
	Age        *int       `json:"age"`
	Collar     *int       `json:"collar"`
	Categories []uint64   `json:"categories"`
	Owner      *uuid.UUID `json:"owner"`
	Friends    []string   `json:"friends"`
}

// Payload of a ent.Pet update request.
type PetUpdateRequest struct {
	Name       *string    `json:"name"`
	Age        *int       `json:"age"`
	Collar     *int       `json:"collar"`
	Categories []uint64   `json:"categories"`
	Owner      *uuid.UUID `json:"owner"`
	Friends    []string   `json:"friends"`
}
