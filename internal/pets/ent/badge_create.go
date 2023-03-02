// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chestarss/elk/internal/pets/ent/badge"
	"github.com/chestarss/elk/internal/pets/ent/pet"
)

// BadgeCreate is the builder for creating a Badge entity.
type BadgeCreate struct {
	config
	mutation *BadgeMutation
	hooks    []Hook
}

// SetColor sets the "color" field.
func (bc *BadgeCreate) SetColor(b badge.Color) *BadgeCreate {
	bc.mutation.SetColor(b)
	return bc
}

// SetMaterial sets the "material" field.
func (bc *BadgeCreate) SetMaterial(b badge.Material) *BadgeCreate {
	bc.mutation.SetMaterial(b)
	return bc
}

// SetID sets the "id" field.
func (bc *BadgeCreate) SetID(u uint32) *BadgeCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetWearerID sets the "wearer" edge to the Pet entity by ID.
func (bc *BadgeCreate) SetWearerID(id int) *BadgeCreate {
	bc.mutation.SetWearerID(id)
	return bc
}

// SetNillableWearerID sets the "wearer" edge to the Pet entity by ID if the given value is not nil.
func (bc *BadgeCreate) SetNillableWearerID(id *int) *BadgeCreate {
	if id != nil {
		bc = bc.SetWearerID(*id)
	}
	return bc
}

// SetWearer sets the "wearer" edge to the Pet entity.
func (bc *BadgeCreate) SetWearer(p *Pet) *BadgeCreate {
	return bc.SetWearerID(p.ID)
}

// Mutation returns the BadgeMutation object of the builder.
func (bc *BadgeCreate) Mutation() *BadgeMutation {
	return bc.mutation
}

// Save creates the Badge in the database.
func (bc *BadgeCreate) Save(ctx context.Context) (*Badge, error) {
	var (
		err  error
		node *Badge
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BadgeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BadgeCreate) SaveX(ctx context.Context) *Badge {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BadgeCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BadgeCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BadgeCreate) check() error {
	if _, ok := bc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "color"`)}
	}
	if v, ok := bc.mutation.Color(); ok {
		if err := badge.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "color": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Material(); !ok {
		return &ValidationError{Name: "material", err: errors.New(`ent: missing required field "material"`)}
	}
	if v, ok := bc.mutation.Material(); ok {
		if err := badge.MaterialValidator(v); err != nil {
			return &ValidationError{Name: "material", err: fmt.Errorf(`ent: validator failed for field "material": %w`, err)}
		}
	}
	return nil
}

func (bc *BadgeCreate) sqlSave(ctx context.Context) (*Badge, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (bc *BadgeCreate) createSpec() (*Badge, *sqlgraph.CreateSpec) {
	var (
		_node = &Badge{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: badge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: badge.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: badge.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := bc.mutation.Material(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: badge.FieldMaterial,
		})
		_node.Material = value
	}
	if nodes := bc.mutation.WearerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   badge.WearerTable,
			Columns: []string{badge.WearerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pet_badge = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BadgeCreateBulk is the builder for creating many Badge entities in bulk.
type BadgeCreateBulk struct {
	config
	builders []*BadgeCreate
}

// Save creates the Badge entities in the database.
func (bcb *BadgeCreateBulk) Save(ctx context.Context) ([]*Badge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Badge, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BadgeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BadgeCreateBulk) SaveX(ctx context.Context) []*Badge {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BadgeCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BadgeCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
