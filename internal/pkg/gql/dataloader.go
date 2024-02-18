package gql

import (
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

type Loader[T any] interface {
	Records(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]T, error)
}

func Load[T any](r Loader[T]) func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		records, err := r.Records(ctx, keys)
		if err != nil {
			return []*dataloader.Result{{Data: nil, Error: err}}
		}
		results := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			if v, ok := records[key]; ok {
				results[i] = &dataloader.Result{Data: v, Error: nil}
			} else {
				var t T
				results[i] = &dataloader.Result{Data: t, Error: nil}
			}
		}
		return results
	}
}

func ToInts(keys dataloader.Keys) []int {
	return lo.Map(keys, func(item dataloader.Key, index int) int {
		return cast.ToInt(item.String())
	})
}

func ToAnySlice(keys dataloader.Keys) []any {
	return lo.Map(keys, func(item dataloader.Key, index int) any {
		return cast.ToInt(item.String())
	})
}

func ToStringKey(id any) dataloader.Key {
	return dataloader.StringKey(cast.ToString(id))
}
