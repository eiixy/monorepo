// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eiixy/monorepo/internal/data/account/ent/menu"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedRoles map[string][]*Role
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			values[i] = new(sql.NullInt64)
		case menu.FieldName, menu.FieldPath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the Menu entity.
func (m *Menu) QueryRoles() *RoleQuery {
	return NewMenuClient(m.config).QueryRoles(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(m.Path)
	builder.WriteByte(')')
	return builder.String()
}

// NamedRoles returns the Roles named value or an error if the edge was not
// loaded in eager-loading with this name.
func (m *Menu) NamedRoles(name string) ([]*Role, error) {
	if m.Edges.namedRoles == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := m.Edges.namedRoles[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (m *Menu) appendNamedRoles(name string, edges ...*Role) {
	if m.Edges.namedRoles == nil {
		m.Edges.namedRoles = make(map[string][]*Role)
	}
	if len(edges) == 0 {
		m.Edges.namedRoles[name] = []*Role{}
	} else {
		m.Edges.namedRoles[name] = append(m.Edges.namedRoles[name], edges...)
	}
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
