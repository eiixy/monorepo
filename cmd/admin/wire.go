//go:build wireinject
// +build wireinject

package main

import (
	"github.com/eiixy/monorepo/internal/app/admin/biz"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/app/admin/server"
	"github.com/eiixy/monorepo/internal/app/admin/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(log.Logger, *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
