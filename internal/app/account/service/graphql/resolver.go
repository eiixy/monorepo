package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/eiixy/monorepo/internal/app/account/biz"
	"github.com/eiixy/monorepo/internal/app/account/server/auth"
	"github.com/eiixy/monorepo/internal/data/account/ent"
	"github.com/go-kratos/kratos/v2/log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log            *log.Helper
	client         *ent.Client
	accountUseCase *biz.AccountUseCase
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger log.Logger, client *ent.Client, accountUseCase *biz.AccountUseCase) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			log:            log.NewHelper(log.With(logger, "module", "service/graphql")),
			client:         client,
			accountUseCase: accountUseCase,
		},
		Directives: DirectiveRoot{
			Disabled: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				return nil, ErrDisabled
			},
			Login: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
				if id := auth.GetUserId(ctx); id != 0 {
					return next(ctx)
				}
				return nil, ErrUnauthorized
			},
			HasPermission: func(ctx context.Context, obj interface{}, next graphql.Resolver, key string) (res interface{}, err error) {
				id := auth.GetUserId(ctx)
				if id == 0 {
					return nil, ErrUnauthorized
				}
				//keys, err := cache.LocalRemember(fmt.Sprintf("user:%d:permissions", id), time.Minute*2, func() ([]string, error) {
				//	return client.Permission.Query().Where(permission.HasRolesWith(role.HasUsersWith(user.ID(id)))).Select(permission.FieldKey).Strings(ctx)
				//})
				//if err != nil {
				//	return nil, err
				//}
				//if !slices.Contains(keys, key) {
				//	return res, ErrNoPermission
				//}
				return next(ctx)
			},
		},
	})
}
