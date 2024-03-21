package service

import (
	"github.com/eiixy/monorepo/internal/app/account/service/graphql"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(graphql.NewSchema)
