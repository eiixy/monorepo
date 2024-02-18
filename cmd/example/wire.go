//go:build wireinject
// +build wireinject

package main

import (
	"github.com/eiixy/monorepo/internal/app/example/conf"
	"github.com/eiixy/monorepo/internal/app/example/data"
	"github.com/eiixy/monorepo/internal/app/example/server"
	"github.com/eiixy/monorepo/internal/app/example/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(log.Logger, *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
