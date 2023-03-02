// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chestarss/elk/internal/openapi/ent/owner"
	"github.com/chestarss/elk/internal/openapi/ent/pet"
	"github.com/chestarss/elk/internal/openapi/ent/predicate"
)

// OwnerQuery is the builder for querying Owner entities.
type OwnerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Owner
	// eager-loading edges.
	withPets *PetQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OwnerQuery builder.
func (oq *OwnerQuery) Where(ps ...predicate.Owner) *OwnerQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OwnerQuery) Limit(limit int) *OwnerQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OwnerQuery) Offset(offset int) *OwnerQuery {
	oq.offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OwnerQuery) Unique(unique bool) *OwnerQuery {
	oq.unique = &unique
	return oq
}

// Order adds an order step to the query.
func (oq *OwnerQuery) Order(o ...OrderFunc) *OwnerQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryPets chains the current query on the "pets" edge.
func (oq *OwnerQuery) QueryPets() *PetQuery {
	query := &PetQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(owner.Table, owner.FieldID, selector),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, owner.PetsTable, owner.PetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Owner entity from the query.
// Returns a *NotFoundError when no Owner was found.
func (oq *OwnerQuery) First(ctx context.Context) (*Owner, error) {
	nodes, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{owner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OwnerQuery) FirstX(ctx context.Context) *Owner {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Owner ID from the query.
// Returns a *NotFoundError when no Owner ID was found.
func (oq *OwnerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{owner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OwnerQuery) FirstIDX(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Owner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Owner entity is not found.
// Returns a *NotFoundError when no Owner entities are found.
func (oq *OwnerQuery) Only(ctx context.Context) (*Owner, error) {
	nodes, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{owner.Label}
	default:
		return nil, &NotSingularError{owner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OwnerQuery) OnlyX(ctx context.Context) *Owner {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Owner ID in the query.
// Returns a *NotSingularError when exactly one Owner ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oq *OwnerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = &NotSingularError{owner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OwnerQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Owners.
func (oq *OwnerQuery) All(ctx context.Context) ([]*Owner, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *OwnerQuery) AllX(ctx context.Context) []*Owner {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Owner IDs.
func (oq *OwnerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(owner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OwnerQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OwnerQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OwnerQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OwnerQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OwnerQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OwnerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OwnerQuery) Clone() *OwnerQuery {
	if oq == nil {
		return nil
	}
	return &OwnerQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		predicates: append([]predicate.Owner{}, oq.predicates...),
		withPets:   oq.withPets.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithPets tells the query-builder to eager-load the nodes that are connected to
// the "pets" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OwnerQuery) WithPets(opts ...func(*PetQuery)) *OwnerQuery {
	query := &PetQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withPets = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Owner.Query().
//		GroupBy(owner.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oq *OwnerQuery) GroupBy(field string, fields ...string) *OwnerGroupBy {
	group := &OwnerGroupBy{config: oq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Owner.Query().
//		Select(owner.FieldName).
//		Scan(ctx, &v)
//
func (oq *OwnerQuery) Select(fields ...string) *OwnerSelect {
	oq.fields = append(oq.fields, fields...)
	return &OwnerSelect{OwnerQuery: oq}
}

func (oq *OwnerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oq.fields {
		if !owner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OwnerQuery) sqlAll(ctx context.Context) ([]*Owner, error) {
	var (
		nodes       = []*Owner{}
		_spec       = oq.querySpec()
		loadedTypes = [1]bool{
			oq.withPets != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Owner{config: oq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oq.withPets; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Owner)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Pets = []*Pet{}
		}
		query.withFKs = true
		query.Where(predicate.Pet(func(s *sql.Selector) {
			s.Where(sql.InValues(owner.PetsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.owner_pets
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "owner_pets" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "owner_pets" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Pets = append(node.Edges.Pets, n)
		}
	}

	return nodes, nil
}

func (oq *OwnerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OwnerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oq *OwnerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   owner.Table,
			Columns: owner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: owner.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if unique := oq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, owner.FieldID)
		for i := range fields {
			if fields[i] != owner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OwnerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(owner.Table)
	columns := oq.fields
	if len(columns) == 0 {
		columns = owner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OwnerGroupBy is the group-by builder for Owner entities.
type OwnerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OwnerGroupBy) Aggregate(fns ...AggregateFunc) *OwnerGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scans the result into the given value.
func (ogb *OwnerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogb *OwnerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OwnerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogb *OwnerGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogb *OwnerGroupBy) StringX(ctx context.Context) string {
	v, err := ogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OwnerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogb *OwnerGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogb *OwnerGroupBy) IntX(ctx context.Context) int {
	v, err := ogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OwnerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogb *OwnerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogb *OwnerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OwnerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogb *OwnerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ogb *OwnerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogb *OwnerGroupBy) BoolX(ctx context.Context) bool {
	v, err := ogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogb *OwnerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ogb.fields {
		if !owner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OwnerGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql.Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
		for _, f := range ogb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ogb.fields...)...)
}

// OwnerSelect is the builder for selecting fields of Owner entities.
type OwnerSelect struct {
	*OwnerQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (os *OwnerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	os.sql = os.OwnerQuery.sqlQuery(ctx)
	return os.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (os *OwnerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := os.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OwnerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (os *OwnerSelect) StringsX(ctx context.Context) []string {
	v, err := os.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = os.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (os *OwnerSelect) StringX(ctx context.Context) string {
	v, err := os.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OwnerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (os *OwnerSelect) IntsX(ctx context.Context) []int {
	v, err := os.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = os.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (os *OwnerSelect) IntX(ctx context.Context) int {
	v, err := os.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OwnerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (os *OwnerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := os.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = os.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (os *OwnerSelect) Float64X(ctx context.Context) float64 {
	v, err := os.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OwnerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (os *OwnerSelect) BoolsX(ctx context.Context) []bool {
	v, err := os.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (os *OwnerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = os.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{owner.Label}
	default:
		err = fmt.Errorf("ent: OwnerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (os *OwnerSelect) BoolX(ctx context.Context) bool {
	v, err := os.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (os *OwnerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := os.sql.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
