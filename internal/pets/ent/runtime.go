// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/chestarss/elk/internal/pets/ent/pet"
	"github.com/chestarss/elk/internal/pets/ent/schema"
	"github.com/chestarss/elk/internal/pets/ent/toy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescHeight is the schema descriptor for height field.
	petDescHeight := petFields[0].Descriptor()
	// pet.HeightValidator is a validator for the "height" field. It is called by the builders before save.
	pet.HeightValidator = petDescHeight.Validators[0].(func(int) error)
	// petDescWeight is the schema descriptor for weight field.
	petDescWeight := petFields[1].Descriptor()
	// pet.WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	pet.WeightValidator = petDescWeight.Validators[0].(func(float64) error)
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[3].Descriptor()
	// pet.NameValidator is a validator for the "name" field. It is called by the builders before save.
	pet.NameValidator = petDescName.Validators[0].(func(string) error)
	// petDescChip is the schema descriptor for chip field.
	petDescChip := petFields[7].Descriptor()
	// pet.DefaultChip holds the default value on creation for the chip field.
	pet.DefaultChip = petDescChip.Default.(func() uuid.UUID)
	toyFields := schema.Toy{}.Fields()
	_ = toyFields
	// toyDescID is the schema descriptor for id field.
	toyDescID := toyFields[0].Descriptor()
	// toy.DefaultID holds the default value on creation for the id field.
	toy.DefaultID = toyDescID.Default.(func() uuid.UUID)
}
