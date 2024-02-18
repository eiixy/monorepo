package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eiixy/monorepo/internal/app/example/conf"
	"github.com/eiixy/monorepo/internal/app/example/service/graphql/dataloader"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, client *ent.Client, schema graphql.ExecutableSchema) *http.Server {
	srv := http.NewServer(cfg.Server.Http.HttpOptions(logger)...)
	// graphql
	gqlSrv := handler.NewDefaultServer(schema)
	loader := dataloader.NewDataLoader(client)

	srv.Handle("/example/query", dataloader.Middleware(loader, gqlSrv))
	srv.HandleFunc("/example/graphql-ui", playground.Handler("Example", "/example/query"))
	return srv
}
