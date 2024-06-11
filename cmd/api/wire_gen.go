// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/eiixy/monorepo/internal/app/api/conf"
	"github.com/eiixy/monorepo/internal/app/api/data"
	"github.com/eiixy/monorepo/internal/app/api/server"
	"github.com/eiixy/monorepo/internal/app/api/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func initApp(logger log.Logger, config *conf.Config) (*kratos.App, func(), error) {
	client, err := data.NewEntClient(config)
	if err != nil {
		return nil, nil, err
	}
	accountService := service.NewAccountService(client)
	httpServer := server.NewHTTPServer(config, logger, accountService)
	app := newApp(logger, httpServer)
	return app, func() {
	}, nil
}
