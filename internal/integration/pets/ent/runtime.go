// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/masseelch/elk/internal/integration/pets/ent/pet"
	"github.com/masseelch/elk/internal/integration/pets/ent/schema"
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
	// petDescChip is the schema descriptor for chip field.
	petDescChip := petFields[7].Descriptor()
	// pet.DefaultChip holds the default value on creation for the chip field.
	pet.DefaultChip = petDescChip.Default.(func() uuid.UUID)
}
