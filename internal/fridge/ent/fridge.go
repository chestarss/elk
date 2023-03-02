// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/chestarss/elk/internal/fridge/ent/fridge"
)

// Fridge is the model entity for the Fridge schema.
type Fridge struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FridgeQuery when eager-loading is set.
	Edges FridgeEdges `json:"edges"`
}

// FridgeEdges holds the relations/edges for other nodes in the graph.
type FridgeEdges struct {
	// Compartments holds the value of the compartments edge.
	Compartments []*Compartment `json:"compartments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CompartmentsOrErr returns the Compartments value or an error if the edge
// was not loaded in eager-loading.
func (e FridgeEdges) CompartmentsOrErr() ([]*Compartment, error) {
	if e.loadedTypes[0] {
		return e.Compartments, nil
	}
	return nil, &NotLoadedError{edge: "compartments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Fridge) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fridge.FieldID:
			values[i] = new(sql.NullInt64)
		case fridge.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Fridge", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Fridge fields.
func (f *Fridge) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fridge.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case fridge.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				f.Title = value.String
			}
		}
	}
	return nil
}

// QueryCompartments queries the "compartments" edge of the Fridge entity.
func (f *Fridge) QueryCompartments() *CompartmentQuery {
	return (&FridgeClient{config: f.config}).QueryCompartments(f)
}

// Update returns a builder for updating this Fridge.
// Note that you need to call Fridge.Unwrap() before calling this method if this Fridge
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Fridge) Update() *FridgeUpdateOne {
	return (&FridgeClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Fridge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Fridge) Unwrap() *Fridge {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Fridge is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Fridge) String() string {
	var builder strings.Builder
	builder.WriteString("Fridge(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", title=")
	builder.WriteString(f.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Fridges is a parsable slice of Fridge.
type Fridges []*Fridge

func (f Fridges) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
