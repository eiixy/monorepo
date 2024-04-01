package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/app/admin/server/auth"
	"github.com/eiixy/monorepo/internal/app/admin/service/graphql/dataloader"
	"github.com/eiixy/monorepo/internal/data/admin/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, client *ent.Client, schema graphql.ExecutableSchema) *http.Server {
	srv := http.NewServer(cfg.Server.Http.HttpOptions(logger)...)
	// graphql
	gqlSrv := handler.NewDefaultServer(schema)
	loader := dataloader.NewDataLoader(client)
	srv.Handle("/admin/query", auth.Middleware(cfg.Key, client, dataloader.Middleware(loader, gqlSrv)))
	srv.HandleFunc("/admin/graphql-ui", playground.Handler("Admin", "/admin/query"))
	return srv
}