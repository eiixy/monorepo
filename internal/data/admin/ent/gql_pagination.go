// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/eiixy/monorepo/internal/data/admin/ent/operationlog"
	"github.com/eiixy/monorepo/internal/data/admin/ent/permission"
	"github.com/eiixy/monorepo/internal/data/admin/ent/role"
	"github.com/eiixy/monorepo/internal/data/admin/ent/user"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// OperationLogEdge is the edge representation of OperationLog.
type OperationLogEdge struct {
	Node   *OperationLog `json:"node"`
	Cursor Cursor        `json:"cursor"`
}

// OperationLogConnection is the connection containing edges to OperationLog.
type OperationLogConnection struct {
	Edges      []*OperationLogEdge `json:"edges"`
	PageInfo   PageInfo            `json:"pageInfo"`
	TotalCount int                 `json:"totalCount"`
}

func (c *OperationLogConnection) build(nodes []*OperationLog, pager *operationlogPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *OperationLog
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *OperationLog {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *OperationLog {
			return nodes[i]
		}
	}
	c.Edges = make([]*OperationLogEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &OperationLogEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// OperationLogPaginateOption enables pagination customization.
type OperationLogPaginateOption func(*operationlogPager) error

// WithOperationLogOrder configures pagination ordering.
func WithOperationLogOrder(order *OperationLogOrder) OperationLogPaginateOption {
	if order == nil {
		order = DefaultOperationLogOrder
	}
	o := *order
	return func(pager *operationlogPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultOperationLogOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithOperationLogFilter configures pagination filter.
func WithOperationLogFilter(filter func(*OperationLogQuery) (*OperationLogQuery, error)) OperationLogPaginateOption {
	return func(pager *operationlogPager) error {
		if filter == nil {
			return errors.New("OperationLogQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type operationlogPager struct {
	reverse bool
	order   *OperationLogOrder
	filter  func(*OperationLogQuery) (*OperationLogQuery, error)
}

func newOperationLogPager(opts []OperationLogPaginateOption, reverse bool) (*operationlogPager, error) {
	pager := &operationlogPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOperationLogOrder
	}
	return pager, nil
}

func (p *operationlogPager) applyFilter(query *OperationLogQuery) (*OperationLogQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *operationlogPager) toCursor(ol *OperationLog) Cursor {
	return p.order.Field.toCursor(ol)
}

func (p *operationlogPager) applyCursors(query *OperationLogQuery, after, before *Cursor) (*OperationLogQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultOperationLogOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *operationlogPager) applyOrder(query *OperationLogQuery) *OperationLogQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultOperationLogOrder.Field {
		query = query.Order(DefaultOperationLogOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *operationlogPager) orderExpr(query *OperationLogQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultOperationLogOrder.Field {
			b.Comma().Ident(DefaultOperationLogOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to OperationLog.
func (ol *OperationLogQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...OperationLogPaginateOption,
) (*OperationLogConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newOperationLogPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if ol, err = pager.applyFilter(ol); err != nil {
		return nil, err
	}
	conn := &OperationLogConnection{Edges: []*OperationLogEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = ol.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if ol, err = pager.applyCursors(ol, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		ol.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := ol.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	ol = pager.applyOrder(ol)
	nodes, err := ol.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// OperationLogOrderField defines the ordering field of OperationLog.
type OperationLogOrderField struct {
	// Value extracts the ordering value from the given OperationLog.
	Value    func(*OperationLog) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) operationlog.OrderOption
	toCursor func(*OperationLog) Cursor
}

// OperationLogOrder defines the ordering of OperationLog.
type OperationLogOrder struct {
	Direction OrderDirection          `json:"direction"`
	Field     *OperationLogOrderField `json:"field"`
}

// DefaultOperationLogOrder is the default ordering of OperationLog.
var DefaultOperationLogOrder = &OperationLogOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &OperationLogOrderField{
		Value: func(ol *OperationLog) (ent.Value, error) {
			return ol.ID, nil
		},
		column: operationlog.FieldID,
		toTerm: operationlog.ByID,
		toCursor: func(ol *OperationLog) Cursor {
			return Cursor{ID: ol.ID}
		},
	},
}

// ToEdge converts OperationLog into OperationLogEdge.
func (ol *OperationLog) ToEdge(order *OperationLogOrder) *OperationLogEdge {
	if order == nil {
		order = DefaultOperationLogOrder
	}
	return &OperationLogEdge{
		Node:   ol,
		Cursor: order.Field.toCursor(ol),
	}
}

// PermissionEdge is the edge representation of Permission.
type PermissionEdge struct {
	Node   *Permission `json:"node"`
	Cursor Cursor      `json:"cursor"`
}

// PermissionConnection is the connection containing edges to Permission.
type PermissionConnection struct {
	Edges      []*PermissionEdge `json:"edges"`
	PageInfo   PageInfo          `json:"pageInfo"`
	TotalCount int               `json:"totalCount"`
}

func (c *PermissionConnection) build(nodes []*Permission, pager *permissionPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Permission
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Permission {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Permission {
			return nodes[i]
		}
	}
	c.Edges = make([]*PermissionEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &PermissionEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// PermissionPaginateOption enables pagination customization.
type PermissionPaginateOption func(*permissionPager) error

// WithPermissionOrder configures pagination ordering.
func WithPermissionOrder(order *PermissionOrder) PermissionPaginateOption {
	if order == nil {
		order = DefaultPermissionOrder
	}
	o := *order
	return func(pager *permissionPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultPermissionOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithPermissionFilter configures pagination filter.
func WithPermissionFilter(filter func(*PermissionQuery) (*PermissionQuery, error)) PermissionPaginateOption {
	return func(pager *permissionPager) error {
		if filter == nil {
			return errors.New("PermissionQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type permissionPager struct {
	reverse bool
	order   *PermissionOrder
	filter  func(*PermissionQuery) (*PermissionQuery, error)
}

func newPermissionPager(opts []PermissionPaginateOption, reverse bool) (*permissionPager, error) {
	pager := &permissionPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultPermissionOrder
	}
	return pager, nil
}

func (p *permissionPager) applyFilter(query *PermissionQuery) (*PermissionQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *permissionPager) toCursor(pe *Permission) Cursor {
	return p.order.Field.toCursor(pe)
}

func (p *permissionPager) applyCursors(query *PermissionQuery, after, before *Cursor) (*PermissionQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultPermissionOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *permissionPager) applyOrder(query *PermissionQuery) *PermissionQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultPermissionOrder.Field {
		query = query.Order(DefaultPermissionOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *permissionPager) orderExpr(query *PermissionQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultPermissionOrder.Field {
			b.Comma().Ident(DefaultPermissionOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Permission.
func (pe *PermissionQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...PermissionPaginateOption,
) (*PermissionConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newPermissionPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if pe, err = pager.applyFilter(pe); err != nil {
		return nil, err
	}
	conn := &PermissionConnection{Edges: []*PermissionEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = pe.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if pe, err = pager.applyCursors(pe, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		pe.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := pe.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	pe = pager.applyOrder(pe)
	nodes, err := pe.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// PermissionOrderField defines the ordering field of Permission.
type PermissionOrderField struct {
	// Value extracts the ordering value from the given Permission.
	Value    func(*Permission) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) permission.OrderOption
	toCursor func(*Permission) Cursor
}

// PermissionOrder defines the ordering of Permission.
type PermissionOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *PermissionOrderField `json:"field"`
}

// DefaultPermissionOrder is the default ordering of Permission.
var DefaultPermissionOrder = &PermissionOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &PermissionOrderField{
		Value: func(pe *Permission) (ent.Value, error) {
			return pe.ID, nil
		},
		column: permission.FieldID,
		toTerm: permission.ByID,
		toCursor: func(pe *Permission) Cursor {
			return Cursor{ID: pe.ID}
		},
	},
}

// ToEdge converts Permission into PermissionEdge.
func (pe *Permission) ToEdge(order *PermissionOrder) *PermissionEdge {
	if order == nil {
		order = DefaultPermissionOrder
	}
	return &PermissionEdge{
		Node:   pe,
		Cursor: order.Field.toCursor(pe),
	}
}

// RoleEdge is the edge representation of Role.
type RoleEdge struct {
	Node   *Role  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// RoleConnection is the connection containing edges to Role.
type RoleConnection struct {
	Edges      []*RoleEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *RoleConnection) build(nodes []*Role, pager *rolePager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Role
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Role {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Role {
			return nodes[i]
		}
	}
	c.Edges = make([]*RoleEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &RoleEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// RolePaginateOption enables pagination customization.
type RolePaginateOption func(*rolePager) error

// WithRoleOrder configures pagination ordering.
func WithRoleOrder(order *RoleOrder) RolePaginateOption {
	if order == nil {
		order = DefaultRoleOrder
	}
	o := *order
	return func(pager *rolePager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultRoleOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithRoleFilter configures pagination filter.
func WithRoleFilter(filter func(*RoleQuery) (*RoleQuery, error)) RolePaginateOption {
	return func(pager *rolePager) error {
		if filter == nil {
			return errors.New("RoleQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type rolePager struct {
	reverse bool
	order   *RoleOrder
	filter  func(*RoleQuery) (*RoleQuery, error)
}

func newRolePager(opts []RolePaginateOption, reverse bool) (*rolePager, error) {
	pager := &rolePager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRoleOrder
	}
	return pager, nil
}

func (p *rolePager) applyFilter(query *RoleQuery) (*RoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *rolePager) toCursor(r *Role) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *rolePager) applyCursors(query *RoleQuery, after, before *Cursor) (*RoleQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultRoleOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *rolePager) applyOrder(query *RoleQuery) *RoleQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultRoleOrder.Field {
		query = query.Order(DefaultRoleOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *rolePager) orderExpr(query *RoleQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultRoleOrder.Field {
			b.Comma().Ident(DefaultRoleOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Role.
func (r *RoleQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...RolePaginateOption,
) (*RoleConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newRolePager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}
	conn := &RoleConnection{Edges: []*RoleEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = r.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if r, err = pager.applyCursors(r, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		r.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	r = pager.applyOrder(r)
	nodes, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// RoleOrderField defines the ordering field of Role.
type RoleOrderField struct {
	// Value extracts the ordering value from the given Role.
	Value    func(*Role) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) role.OrderOption
	toCursor func(*Role) Cursor
}

// RoleOrder defines the ordering of Role.
type RoleOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *RoleOrderField `json:"field"`
}

// DefaultRoleOrder is the default ordering of Role.
var DefaultRoleOrder = &RoleOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &RoleOrderField{
		Value: func(r *Role) (ent.Value, error) {
			return r.ID, nil
		},
		column: role.FieldID,
		toTerm: role.ByID,
		toCursor: func(r *Role) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Role into RoleEdge.
func (r *Role) ToEdge(order *RoleOrder) *RoleEdge {
	if order == nil {
		order = DefaultRoleOrder
	}
	return &RoleEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldEmail orders User by email.
	UserOrderFieldEmail = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.Email, nil
		},
		column: user.FieldEmail,
		toTerm: user.ByEmail,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.Email,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.column {
	case UserOrderFieldEmail.column:
		str = "EMAIL"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "EMAIL":
		*f = *UserOrderFieldEmail
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
