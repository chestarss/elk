// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/masseelch/elk/internal/simple/ent/migrate"

	"github.com/masseelch/elk/internal/simple/ent/category"
	"github.com/masseelch/elk/internal/simple/ent/owner"
	"github.com/masseelch/elk/internal/simple/ent/pet"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Owner is the client for interacting with the Owner builders.
	Owner *OwnerClient
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Category = NewCategoryClient(c.config)
	c.Owner = NewOwnerClient(c.config)
	c.Pet = NewPetClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Category: NewCategoryClient(cfg),
		Owner:    NewOwnerClient(cfg),
		Pet:      NewPetClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Category: NewCategoryClient(cfg),
		Owner:    NewOwnerClient(cfg),
		Pet:      NewPetClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Category.Use(hooks...)
	c.Owner.Use(hooks...)
	c.Pet.Use(hooks...)
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a create builder for Category.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id int) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CategoryClient) DeleteOneID(id int) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id int) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id int) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPets queries the pets edge of a Category.
func (c *CategoryClient) QueryPets(ca *Category) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, category.PetsTable, category.PetsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// OwnerClient is a client for the Owner schema.
type OwnerClient struct {
	config
}

// NewOwnerClient returns a client for the Owner from the given config.
func NewOwnerClient(c config) *OwnerClient {
	return &OwnerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `owner.Hooks(f(g(h())))`.
func (c *OwnerClient) Use(hooks ...Hook) {
	c.hooks.Owner = append(c.hooks.Owner, hooks...)
}

// Create returns a create builder for Owner.
func (c *OwnerClient) Create() *OwnerCreate {
	mutation := newOwnerMutation(c.config, OpCreate)
	return &OwnerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Owner entities.
func (c *OwnerClient) CreateBulk(builders ...*OwnerCreate) *OwnerCreateBulk {
	return &OwnerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Owner.
func (c *OwnerClient) Update() *OwnerUpdate {
	mutation := newOwnerMutation(c.config, OpUpdate)
	return &OwnerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OwnerClient) UpdateOne(o *Owner) *OwnerUpdateOne {
	mutation := newOwnerMutation(c.config, OpUpdateOne, withOwner(o))
	return &OwnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OwnerClient) UpdateOneID(id int) *OwnerUpdateOne {
	mutation := newOwnerMutation(c.config, OpUpdateOne, withOwnerID(id))
	return &OwnerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Owner.
func (c *OwnerClient) Delete() *OwnerDelete {
	mutation := newOwnerMutation(c.config, OpDelete)
	return &OwnerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OwnerClient) DeleteOne(o *Owner) *OwnerDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OwnerClient) DeleteOneID(id int) *OwnerDeleteOne {
	builder := c.Delete().Where(owner.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OwnerDeleteOne{builder}
}

// Query returns a query builder for Owner.
func (c *OwnerClient) Query() *OwnerQuery {
	return &OwnerQuery{
		config: c.config,
	}
}

// Get returns a Owner entity by its id.
func (c *OwnerClient) Get(ctx context.Context, id int) (*Owner, error) {
	return c.Query().Where(owner.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OwnerClient) GetX(ctx context.Context, id int) *Owner {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPets queries the pets edge of a Owner.
func (c *OwnerClient) QueryPets(o *Owner) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(owner.Table, owner.FieldID, id),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, owner.PetsTable, owner.PetsColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OwnerClient) Hooks() []Hook {
	return c.hooks.Owner
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pet.Hooks(f(g(h())))`.
func (c *PetClient) Use(hooks ...Hook) {
	c.hooks.Pet = append(c.hooks.Pet, hooks...)
}

// Create returns a create builder for Pet.
func (c *PetClient) Create() *PetCreate {
	mutation := newPetMutation(c.config, OpCreate)
	return &PetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pet entities.
func (c *PetClient) CreateBulk(builders ...*PetCreate) *PetCreateBulk {
	return &PetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	mutation := newPetMutation(c.config, OpUpdate)
	return &PetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPet(pe))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id string) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPetID(id))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	mutation := newPetMutation(c.config, OpDelete)
	return &PetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PetClient) DeleteOneID(id string) *PetDeleteOne {
	builder := c.Delete().Where(pet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PetDeleteOne{builder}
}

// Query returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{
		config: c.config,
	}
}

// Get returns a Pet entity by its id.
func (c *PetClient) Get(ctx context.Context, id string) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id string) *Pet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCategories queries the categories edge of a Pet.
func (c *PetClient) QueryCategories(pe *Pet) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, pet.CategoriesTable, pet.CategoriesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Pet.
func (c *PetClient) QueryOwner(pe *Pet) *OwnerQuery {
	query := &OwnerQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, id),
			sqlgraph.To(owner.Table, owner.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pet.OwnerTable, pet.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriends queries the friends edge of a Pet.
func (c *PetClient) QueryFriends(pe *Pet) *PetQuery {
	query := &PetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(pet.Table, pet.FieldID, id),
			sqlgraph.To(pet.Table, pet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, pet.FriendsTable, pet.FriendsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PetClient) Hooks() []Hook {
	return c.hooks.Pet
}
