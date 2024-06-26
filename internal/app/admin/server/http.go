package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/app/admin/server/auth"
	"github.com/eiixy/monorepo/internal/app/admin/service/graphql/dataloader"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, data *data.Data, schema graphql.ExecutableSchema) *http.Server {
	srv := http.NewServer(cfg.Server.Http.HttpOptions(logger)...)
	// graphql
	gqlSrv := handler.NewDefaultServer(schema)
	loader := dataloader.NewDataLoader(data.EntClient)
	srv.Handle("/query", auth.Middleware(cfg.Key, data.EntClient, dataloader.Middleware(loader, gqlSrv)))
	srv.HandleFunc("/graphql-ui", playground.Handler("Admin", "/api/admin/query"))
	return srv
}
