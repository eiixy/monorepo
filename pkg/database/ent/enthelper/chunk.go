package enthelper

import (
	"context"
)

type chunkImp[Q, T any] interface {
	Offset(offset int) Q
	Limit(limit int) Q
	All(ctx context.Context) ([]*T, error)
	Count(ctx context.Context) (int, error)
}

// Chunk 分批处理数据
// example:
//
//	client, err := ent.Open("mysql", os.Getenv("ACCOUNT_DB_DSN"))
//	if err != nil {
//		panic(err)
//	}
//	ctx := context.Background()
//	err = enthelper.Chunk(ctx, client.Account.Query(), 100, func(batch int, items []*ent.Account) error {
//		for i, item := range items {
//			fmt.Println(batch, i, item.Email)
//		}
//		return nil
//	})
func Chunk[Q chunkImp[Q, T], T any](ctx context.Context, query Q, chunk int, fn func(batch int, items []*T) error) error {
	count, err := query.Count(ctx)
	if err != nil {
		return err
	}
	for i := 0; i < count; i = i + chunk {
		items, err := query.Offset(i).Limit(chunk).All(ctx)
		if err != nil {
			return err
		}
		err = fn(i/chunk, items)
		if err != nil {
			return err
		}
	}
	return nil
}
