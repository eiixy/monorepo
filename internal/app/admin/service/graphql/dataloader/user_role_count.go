package dataloader

import (
	"context"
	"fmt"
	"github.com/eiixy/monorepo/internal/data/admin/ent"
	"github.com/eiixy/monorepo/internal/pkg/gql"
	"github.com/graph-gophers/dataloader"
	"strings"
)

type userRoleCountLoader struct {
	client *ent.Client
}

func (r userRoleCountLoader) Records(ctx context.Context, keys dataloader.Keys) (map[dataloader.Key]int, error) {
	placeholder := strings.TrimSuffix(strings.Repeat("?,", len(keys)), ",")
	var rows, err = r.client.QueryContext(ctx, fmt.Sprintf(`SELECT user_id, COUNT(role_id) FROM users LEFT JOIN user_roles ON users.id = user_roles.user_id WHERE user_id IN (%s) GROUP BY user_id`, placeholder), gql.ToAnySlice(keys)...)
	if err != nil {
		return nil, err
	}
	var result = make(map[dataloader.Key]int)
	for rows.Next() {
		var userId, count int
		err = rows.Scan(&userId, &count)
		if err != nil {
			return nil, err
		}
		result[gql.ToStringKey(userId)] = count
	}
	return result, nil
}
