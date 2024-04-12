// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"golang.org/x/exp/slices"
)

const (
	itemsField = "items"
	countField = "count"
)

func containsField(ctx context.Context, field string) bool {
	return slices.Contains(graphql.CollectAllFields(ctx), field)
}
func offsetLimit(page *int, size *int) (int, int) {
	var offset, limit = 0, 10
	if size != nil {
		limit = *size
	}
	if page != nil && *page > 1 {
		offset = (*page - 1) * limit
	}
	return offset, limit
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ol *OperationLogQuery) CustomCollectFields(ctx context.Context, path ...string) (*OperationLogQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ol, nil
	}
	if field := collectedField(ctx, path...); field != nil {
		if err := ol.collectField(ctx, graphql.GetOperationContext(ctx), *field, path); err != nil {
			return nil, err
		}
	}
	return ol, nil
}

// List executes the query and returns count and []*OperationLog.
func (ol *OperationLogQuery) List(ctx context.Context, page *int, size *int, orderBy *OperationLogOrder) (items []*OperationLog, count int, err error) {
	if orderBy != nil {
		o := Asc(orderBy.Field.column)
		if orderBy.Direction.String() == "DESC" {
			o = Desc(orderBy.Field.column)
		}
		ol.Order(o)
	}
	if containsField(ctx, countField) {
		count, err = ol.Count(ctx)
		if err != nil {
			return
		}
	}
	if containsField(ctx, itemsField) {
		ol, err = ol.CustomCollectFields(ctx, itemsField)
		if err != nil {
			return
		}
		offset, limit := offsetLimit(page, size)
		items, err = ol.Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return
		}
	}
	return
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PermissionQuery) CustomCollectFields(ctx context.Context, path ...string) (*PermissionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pe, nil
	}
	if field := collectedField(ctx, path...); field != nil {
		if err := pe.collectField(ctx, graphql.GetOperationContext(ctx), *field, path); err != nil {
			return nil, err
		}
	}
	return pe, nil
}

// List executes the query and returns count and []*Permission.
func (pe *PermissionQuery) List(ctx context.Context, page *int, size *int, orderBy *PermissionOrder) (items []*Permission, count int, err error) {
	if orderBy != nil {
		o := Asc(orderBy.Field.column)
		if orderBy.Direction.String() == "DESC" {
			o = Desc(orderBy.Field.column)
		}
		pe.Order(o)
	}
	if containsField(ctx, countField) {
		count, err = pe.Count(ctx)
		if err != nil {
			return
		}
	}
	if containsField(ctx, itemsField) {
		pe, err = pe.CustomCollectFields(ctx, itemsField)
		if err != nil {
			return
		}
		offset, limit := offsetLimit(page, size)
		items, err = pe.Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return
		}
	}
	return
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoleQuery) CustomCollectFields(ctx context.Context, path ...string) (*RoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if field := collectedField(ctx, path...); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, path); err != nil {
			return nil, err
		}
	}
	return r, nil
}

// List executes the query and returns count and []*Role.
func (r *RoleQuery) List(ctx context.Context, page *int, size *int, orderBy *RoleOrder) (items []*Role, count int, err error) {
	if orderBy != nil {
		o := Asc(orderBy.Field.column)
		if orderBy.Direction.String() == "DESC" {
			o = Desc(orderBy.Field.column)
		}
		r.Order(o)
	}
	if containsField(ctx, countField) {
		count, err = r.Count(ctx)
		if err != nil {
			return
		}
	}
	if containsField(ctx, itemsField) {
		r, err = r.CustomCollectFields(ctx, itemsField)
		if err != nil {
			return
		}
		offset, limit := offsetLimit(page, size)
		items, err = r.Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return
		}
	}
	return
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CustomCollectFields(ctx context.Context, path ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if field := collectedField(ctx, path...); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, path); err != nil {
			return nil, err
		}
	}
	return u, nil
}

// List executes the query and returns count and []*User.
func (u *UserQuery) List(ctx context.Context, page *int, size *int, orderBy *UserOrder) (items []*User, count int, err error) {
	if orderBy != nil {
		o := Asc(orderBy.Field.column)
		if orderBy.Direction.String() == "DESC" {
			o = Desc(orderBy.Field.column)
		}
		u.Order(o)
	}
	if containsField(ctx, countField) {
		count, err = u.Count(ctx)
		if err != nil {
			return
		}
	}
	if containsField(ctx, itemsField) {
		u, err = u.CustomCollectFields(ctx, itemsField)
		if err != nil {
			return
		}
		offset, limit := offsetLimit(page, size)
		items, err = u.Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return
		}
	}
	return
}
