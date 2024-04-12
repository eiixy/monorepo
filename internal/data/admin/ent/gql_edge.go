// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (ol *OperationLog) User(ctx context.Context) (*User, error) {
	result, err := ol.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ol.QueryUser().Only(ctx)
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
