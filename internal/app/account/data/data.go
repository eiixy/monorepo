package data

import (
	"github.com/eiixy/monorepo/internal/app/account/conf"
	"github.com/eiixy/monorepo/internal/data/account/ent"
	_ "github.com/eiixy/monorepo/internal/data/account/ent/runtime"
	"github.com/eiixy/monorepo/pkg/database"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEntClient,
)

func NewEntClient(cfg *conf.Config) (*ent.Client, error) {
	drv, err := database.NewEntDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewClient(ent.Driver(drv)), nil
}
