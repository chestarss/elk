// Code generated by entc, DO NOT EDIT.

package http

import (
	"time"

	"github.com/google/uuid"
	badge "github.com/chestarss/elk/internal/pets/ent/badge"
	pet "github.com/chestarss/elk/internal/pets/ent/pet"
	playgroup "github.com/chestarss/elk/internal/pets/ent/playgroup"
	toy "github.com/chestarss/elk/internal/pets/ent/toy"
)

// Payload of a ent.Badge create request.
type BadgeCreateRequest struct {
	Color    *badge.Color    `json:"color"`
	Material *badge.Material `json:"material"`
	Wearer   *int            `json:"wearer"`
}

// Payload of a ent.Badge update request.
type BadgeUpdateRequest struct {
	Color    *badge.Color    `json:"color"`
	Material *badge.Material `json:"material"`
	Wearer   *int            `json:"wearer"`
}

// Payload of a ent.Pet create request.
type PetCreateRequest struct {
	Height     *int        `json:"height"`
	Weight     *float64    `json:"weight"`
	Castrated  *bool       `json:"castrated"`
	Name       *string     `json:"name"`
	Birthday   *time.Time  `json:"birthday"`
	Nicknames  *[]string   `json:"nicknames"`
	Sex        *pet.Sex    `json:"sex"`
	Chip       *uuid.UUID  `json:"chip"`
	Badge      *uint32     `json:"badge"`
	Protege    *int        `json:"protege"`
	Mentor     *int        `json:"mentor"`
	Spouse     *int        `json:"spouse"`
	Toys       []uuid.UUID `json:"toys"`
	Parent     *int        `json:"parent"`
	Children   []int       `json:"children"`
	PlayGroups []int       `json:"play_groups"`
	Friends    []int       `json:"friends"`
}

// Payload of a ent.Pet update request.
type PetUpdateRequest struct {
	Height     *int        `json:"height"`
	Weight     *float64    `json:"weight"`
	Castrated  *bool       `json:"castrated"`
	Name       *string     `json:"name"`
	Birthday   *time.Time  `json:"birthday"`
	Nicknames  *[]string   `json:"nicknames"`
	Chip       *uuid.UUID  `json:"chip"`
	Badge      *uint32     `json:"badge"`
	Protege    *int        `json:"protege"`
	Mentor     *int        `json:"mentor"`
	Spouse     *int        `json:"spouse"`
	Toys       []uuid.UUID `json:"toys"`
	Parent     *int        `json:"parent"`
	Children   []int       `json:"children"`
	PlayGroups []int       `json:"play_groups"`
	Friends    []int       `json:"friends"`
}

// Payload of a ent.PlayGroup create request.
type PlayGroupCreateRequest struct {
	Title        *string            `json:"title"`
	Description  *string            `json:"description"`
	Weekday      *playgroup.Weekday `json:"weekday"`
	Participants []int              `json:"participants"`
}

// Payload of a ent.PlayGroup update request.
type PlayGroupUpdateRequest struct {
	Title        *string            `json:"title"`
	Description  *string            `json:"description"`
	Weekday      *playgroup.Weekday `json:"weekday"`
	Participants []int              `json:"participants"`
}

// Payload of a ent.Toy create request.
type ToyCreateRequest struct {
	Color    *toy.Color    `json:"color"`
	Material *toy.Material `json:"material"`
	Title    *string       `json:"title"`
	Owner    *int          `json:"owner"`
}

// Payload of a ent.Toy update request.
type ToyUpdateRequest struct {
	Color    *toy.Color    `json:"color"`
	Material *toy.Material `json:"material"`
	Title    *string       `json:"title"`
	Owner    *int          `json:"owner"`
}
