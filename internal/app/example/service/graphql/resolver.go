package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/go-kratos/kratos/v2/log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log    *log.Helper
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger log.Logger, client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			log:    log.NewHelper(log.With(logger, "module", "service/graphql")),
			client: client,
		},
	})
}
