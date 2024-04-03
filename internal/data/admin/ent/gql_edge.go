// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (m *Menu) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryRoles().All(ctx)
	}
	return result, err
}

func (m *Menu) Parent(ctx context.Context) (*Menu, error) {
	result, err := m.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Menu) Children(ctx context.Context) (result []*Menu, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryChildren().All(ctx)
	}
	return result, err
}

func (m *Menu) Permissions(ctx context.Context) (result []*Permission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = m.NamedPermissions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = m.Edges.PermissionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = m.QueryPermissions().All(ctx)
	}
	return result, err
}

func (ol *OperationLog) User(ctx context.Context) (*User, error) {
	result, err := ol.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ol.QueryUser().Only(ctx)
	}
	return result, err
}

func (pe *Permission) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryRoles().All(ctx)
	}
	return result, err
}

func (pe *Permission) Menus(ctx context.Context) (result []*Menu, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedMenus(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.MenusOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryMenus().All(ctx)
	}
	return result, err
}

func (pe *Permission) Parent(ctx context.Context) (*Permission, error) {
	result, err := pe.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pe *Permission) Children(ctx context.Context) (result []*Permission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryChildren().All(ctx)
	}
	return result, err
}

func (r *Role) Menus(ctx context.Context) (result []*Menu, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedMenus(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.MenusOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryMenus().All(ctx)
	}
	return result, err
}

func (r *Role) Permissions(ctx context.Context) (result []*Permission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedPermissions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.PermissionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryPermissions().All(ctx)
	}
	return result, err
}

func (r *Role) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryUsers().All(ctx)
	}
	return result, err
}

func (u *User) Roles(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryRoles().All(ctx)
	}
	return result, err
}

func (u *User) OperationLogs(ctx context.Context) (result []*OperationLog, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedOperationLogs(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.OperationLogsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryOperationLogs().All(ctx)
	}
	return result, err
}
