// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/eiixy/monorepo/internal/data/admin/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eiixy/monorepo/internal/data/admin/ent/menu"
	"github.com/eiixy/monorepo/internal/data/admin/ent/operationlog"
	"github.com/eiixy/monorepo/internal/data/admin/ent/permission"
	"github.com/eiixy/monorepo/internal/data/admin/ent/role"
	"github.com/eiixy/monorepo/internal/data/admin/ent/user"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Menu is the client for interacting with the Menu builders.
	Menu *MenuClient
	// OperationLog is the client for interacting with the OperationLog builders.
	OperationLog *OperationLogClient
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Menu = NewMenuClient(c.config)
	c.OperationLog = NewOperationLogClient(c.config)
	c.Permission = NewPermissionClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Menu:         NewMenuClient(cfg),
		OperationLog: NewOperationLogClient(cfg),
		Permission:   NewPermissionClient(cfg),
		Role:         NewRoleClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:          ctx,
		config:       cfg,
		Menu:         NewMenuClient(cfg),
		OperationLog: NewOperationLogClient(cfg),
		Permission:   NewPermissionClient(cfg),
		Role:         NewRoleClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Menu.
//		Query().
//		Count(ctx)
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
	c.Menu.Use(hooks...)
	c.OperationLog.Use(hooks...)
	c.Permission.Use(hooks...)
	c.Role.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Menu.Intercept(interceptors...)
	c.OperationLog.Intercept(interceptors...)
	c.Permission.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MenuMutation:
		return c.Menu.mutate(ctx, m)
	case *OperationLogMutation:
		return c.OperationLog.mutate(ctx, m)
	case *PermissionMutation:
		return c.Permission.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MenuClient is a client for the Menu schema.
type MenuClient struct {
	config
}

// NewMenuClient returns a client for the Menu from the given config.
func NewMenuClient(c config) *MenuClient {
	return &MenuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `menu.Hooks(f(g(h())))`.
func (c *MenuClient) Use(hooks ...Hook) {
	c.hooks.Menu = append(c.hooks.Menu, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `menu.Intercept(f(g(h())))`.
func (c *MenuClient) Intercept(interceptors ...Interceptor) {
	c.inters.Menu = append(c.inters.Menu, interceptors...)
}

// Create returns a builder for creating a Menu entity.
func (c *MenuClient) Create() *MenuCreate {
	mutation := newMenuMutation(c.config, OpCreate)
	return &MenuCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Menu entities.
func (c *MenuClient) CreateBulk(builders ...*MenuCreate) *MenuCreateBulk {
	return &MenuCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MenuClient) MapCreateBulk(slice any, setFunc func(*MenuCreate, int)) *MenuCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MenuCreateBulk{err: fmt.Errorf("calling to MenuClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MenuCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MenuCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Menu.
func (c *MenuClient) Update() *MenuUpdate {
	mutation := newMenuMutation(c.config, OpUpdate)
	return &MenuUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MenuClient) UpdateOne(m *Menu) *MenuUpdateOne {
	mutation := newMenuMutation(c.config, OpUpdateOne, withMenu(m))
	return &MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MenuClient) UpdateOneID(id int) *MenuUpdateOne {
	mutation := newMenuMutation(c.config, OpUpdateOne, withMenuID(id))
	return &MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Menu.
func (c *MenuClient) Delete() *MenuDelete {
	mutation := newMenuMutation(c.config, OpDelete)
	return &MenuDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MenuClient) DeleteOne(m *Menu) *MenuDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MenuClient) DeleteOneID(id int) *MenuDeleteOne {
	builder := c.Delete().Where(menu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MenuDeleteOne{builder}
}

// Query returns a query builder for Menu.
func (c *MenuClient) Query() *MenuQuery {
	return &MenuQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMenu},
		inters: c.Interceptors(),
	}
}

// Get returns a Menu entity by its id.
func (c *MenuClient) Get(ctx context.Context, id int) (*Menu, error) {
	return c.Query().Where(menu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MenuClient) GetX(ctx context.Context, id int) *Menu {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Menu.
func (c *MenuClient) QueryRoles(m *Menu) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, menu.RolesTable, menu.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Menu.
func (c *MenuClient) QueryParent(m *Menu) *MenuQuery {
	query := (&MenuClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, id),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, menu.ParentTable, menu.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Menu.
func (c *MenuClient) QueryChildren(m *Menu) *MenuQuery {
	query := (&MenuClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, id),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, menu.ChildrenTable, menu.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MenuClient) Hooks() []Hook {
	return c.hooks.Menu
}

// Interceptors returns the client interceptors.
func (c *MenuClient) Interceptors() []Interceptor {
	return c.inters.Menu
}

func (c *MenuClient) mutate(ctx context.Context, m *MenuMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MenuCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MenuUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MenuDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Menu mutation op: %q", m.Op())
	}
}

// OperationLogClient is a client for the OperationLog schema.
type OperationLogClient struct {
	config
}

// NewOperationLogClient returns a client for the OperationLog from the given config.
func NewOperationLogClient(c config) *OperationLogClient {
	return &OperationLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `operationlog.Hooks(f(g(h())))`.
func (c *OperationLogClient) Use(hooks ...Hook) {
	c.hooks.OperationLog = append(c.hooks.OperationLog, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `operationlog.Intercept(f(g(h())))`.
func (c *OperationLogClient) Intercept(interceptors ...Interceptor) {
	c.inters.OperationLog = append(c.inters.OperationLog, interceptors...)
}

// Create returns a builder for creating a OperationLog entity.
func (c *OperationLogClient) Create() *OperationLogCreate {
	mutation := newOperationLogMutation(c.config, OpCreate)
	return &OperationLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OperationLog entities.
func (c *OperationLogClient) CreateBulk(builders ...*OperationLogCreate) *OperationLogCreateBulk {
	return &OperationLogCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OperationLogClient) MapCreateBulk(slice any, setFunc func(*OperationLogCreate, int)) *OperationLogCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OperationLogCreateBulk{err: fmt.Errorf("calling to OperationLogClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OperationLogCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OperationLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OperationLog.
func (c *OperationLogClient) Update() *OperationLogUpdate {
	mutation := newOperationLogMutation(c.config, OpUpdate)
	return &OperationLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OperationLogClient) UpdateOne(ol *OperationLog) *OperationLogUpdateOne {
	mutation := newOperationLogMutation(c.config, OpUpdateOne, withOperationLog(ol))
	return &OperationLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OperationLogClient) UpdateOneID(id int) *OperationLogUpdateOne {
	mutation := newOperationLogMutation(c.config, OpUpdateOne, withOperationLogID(id))
	return &OperationLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OperationLog.
func (c *OperationLogClient) Delete() *OperationLogDelete {
	mutation := newOperationLogMutation(c.config, OpDelete)
	return &OperationLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OperationLogClient) DeleteOne(ol *OperationLog) *OperationLogDeleteOne {
	return c.DeleteOneID(ol.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OperationLogClient) DeleteOneID(id int) *OperationLogDeleteOne {
	builder := c.Delete().Where(operationlog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OperationLogDeleteOne{builder}
}

// Query returns a query builder for OperationLog.
func (c *OperationLogClient) Query() *OperationLogQuery {
	return &OperationLogQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOperationLog},
		inters: c.Interceptors(),
	}
}

// Get returns a OperationLog entity by its id.
func (c *OperationLogClient) Get(ctx context.Context, id int) (*OperationLog, error) {
	return c.Query().Where(operationlog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OperationLogClient) GetX(ctx context.Context, id int) *OperationLog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a OperationLog.
func (c *OperationLogClient) QueryUser(ol *OperationLog) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ol.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(operationlog.Table, operationlog.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, operationlog.UserTable, operationlog.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ol.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OperationLogClient) Hooks() []Hook {
	return c.hooks.OperationLog
}

// Interceptors returns the client interceptors.
func (c *OperationLogClient) Interceptors() []Interceptor {
	return c.inters.OperationLog
}

func (c *OperationLogClient) mutate(ctx context.Context, m *OperationLogMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OperationLogCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OperationLogUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OperationLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OperationLogDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OperationLog mutation op: %q", m.Op())
	}
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `permission.Intercept(f(g(h())))`.
func (c *PermissionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Permission = append(c.inters.Permission, interceptors...)
}

// Create returns a builder for creating a Permission entity.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PermissionClient) MapCreateBulk(slice any, setFunc func(*PermissionCreate, int)) *PermissionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PermissionCreateBulk{err: fmt.Errorf("calling to PermissionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PermissionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id int) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PermissionClient) DeleteOneID(id int) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePermission},
		inters: c.Interceptors(),
	}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id int) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id int) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Permission.
func (c *PermissionClient) QueryRoles(pe *Permission) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.RolesTable, permission.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Permission.
func (c *PermissionClient) QueryParent(pe *Permission) *PermissionQuery {
	query := (&PermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, permission.ParentTable, permission.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Permission.
func (c *PermissionClient) QueryChildren(pe *Permission) *PermissionQuery {
	query := (&PermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, permission.ChildrenTable, permission.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// Interceptors returns the client interceptors.
func (c *PermissionClient) Interceptors() []Interceptor {
	return c.inters.Permission
}

func (c *PermissionClient) mutate(ctx context.Context, m *PermissionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Permission mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoleClient) MapCreateBulk(slice any, setFunc func(*RoleCreate, int)) *RoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoleCreateBulk{err: fmt.Errorf("calling to RoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMenus queries the menus edge of a Role.
func (c *RoleClient) QueryMenus(r *Role) *MenuQuery {
	query := (&MenuClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.MenusTable, role.MenusPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermissions queries the permissions edge of a Role.
func (c *RoleClient) QueryPermissions(r *Role) *PermissionQuery {
	query := (&PermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.PermissionsTable, role.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Role.
func (c *RoleClient) QueryUsers(r *Role) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.UsersTable, role.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a User.
func (c *UserClient) QueryRoles(u *User) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.RolesTable, user.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOperationLogs queries the operation_logs edge of a User.
func (c *UserClient) QueryOperationLogs(u *User) *OperationLogQuery {
	query := (&OperationLogClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(operationlog.Table, operationlog.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.OperationLogsTable, user.OperationLogsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	inters := c.inters.User
	return append(inters[:len(inters):len(inters)], user.Interceptors[:]...)
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Menu, OperationLog, Permission, Role, User []ent.Hook
	}
	inters struct {
		Menu, OperationLog, Permission, Role, User []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}