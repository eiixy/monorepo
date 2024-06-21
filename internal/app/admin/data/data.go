package data

import (
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	_ "github.com/eiixy/monorepo/internal/data/example/ent/runtime"
	"github.com/eiixy/monorepo/pkg/database"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
)

type Data struct {
	EntClient *ent.Client
	EntDB     *ent.Database
}

func NewData(cfg *conf.Config) (*Data, func(), error) {
	drv, err := database.NewEntDriver(cfg.Data.Database)
	//drv, err := database.NewEntDriverWithOtel(cfg.Data.Database)
	if err != nil {
		return nil, func() {}, err
	}
	client := ent.NewClient(ent.Driver(drv))
	return &Data{
			EntClient: client,
			EntDB:     ent.NewDatabase(ent.Driver(drv)),
		}, func() {
			_ = drv.Close()
		}, nil
}
