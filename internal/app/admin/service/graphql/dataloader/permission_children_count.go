package dataloader

import (
	"context"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/eiixy/monorepo/internal/data/example/ent/permission"
	"github.com/eiixy/monorepo/internal/pkg/gql"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
)

type permissionChildrenCountLoader struct {
	client *ent.Client
}

func (r permissionChildrenCountLoader) Records(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]int, error) {
	type item struct {
		ParentId int64 `json:"parent_id"`
		Count    int   `json:"count"`
	}
	var items []item
	err := r.client.Permission.Query().Where(permission.ParentIDIn(gql.ToInts(keys)...)).
		GroupBy(permission.FieldParentID).Aggregate(ent.Count()).Scan(ctx, &items)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(items, func(item item) (dataloader.Key, int) {
		return gql.ToStringKey(item.ParentId), item.Count
	}), nil
}
