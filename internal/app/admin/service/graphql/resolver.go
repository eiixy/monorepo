package graphql

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/eiixy/monorepo/internal/app/admin/biz"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/app/admin/server/auth"
	"github.com/eiixy/monorepo/internal/data/example/ent/permission"
	"github.com/eiixy/monorepo/internal/data/example/ent/role"
	"github.com/eiixy/monorepo/internal/data/example/ent/user"
	"github.com/eiixy/monorepo/pkg/cache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/exp/slices"
	"image/color"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log            *log.Helper
	data           *data.Data
	accountUseCase *biz.AccountUseCase
	captcha        *base64Captcha.Captcha
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger log.Logger, data *data.Data, accountUseCase *biz.AccountUseCase) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			log:            log.NewHelper(log.With(logger, "module", "service/graphql")),
			data:           data,
			accountUseCase: accountUseCase,
			captcha:        base64Captcha.NewCaptcha(base64Captcha.NewDriverString(40, 140, 0, 0, 4, "1234567890abcdefghijklmnopqrktuvwxyz", &color.RGBA{}, base64Captcha.DefaultEmbeddedFonts, nil), base64Captcha.DefaultMemStore),
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
				u := auth.GetUser(ctx)
				if u == nil {
					return nil, ErrUnauthorized
				}
				if !u.IsAdmin {
					keys, err := cache.LocalRemember(fmt.Sprintf("user:%d:permissions", u.ID), time.Minute*2, func() ([]string, error) {
						return data.EntDB.Permission(ctx).Query().Where(permission.HasRolesWith(role.HasUsersWith(user.ID(u.ID)))).Select(permission.FieldKey).Strings(ctx)
					})
					if err != nil {
						return nil, err
					}
					if !slices.Contains(keys, key) {
						return res, ErrNoPermission
					}
				}
				return next(ctx)
			},
		},
	})
}
