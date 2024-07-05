package data

import (
	"github.com/IBM/sarama"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	_ "github.com/eiixy/monorepo/internal/data/example/ent/runtime"
	"github.com/eiixy/monorepo/pkg/database"
	"github.com/eiixy/monorepo/pkg/kafka"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
)

type Data struct {
	EntClient *ent.Client
	EntDB     *ent.Database
	Producer  sarama.SyncProducer
}

func NewData(cfg *conf.Config) (*Data, func(), error) {
	drv, err := database.NewEntDriver(cfg.Data.Database)
	//drv, err := database.NewEntDriverWithOtel(cfg.Data.Database)
	if err != nil {
		return nil, func() {}, err
	}
	client := ent.NewClient(ent.Driver(drv))
	producer, err := kafka.NewSyncProducerFromConfig(cfg.Data.Kafka)
	if err != nil {
		return nil, nil, err
	}
	return &Data{
			EntClient: client,
			EntDB:     ent.NewDatabase(ent.Driver(drv)),
			Producer:  producer,
		}, func() {
			_ = drv.Close()
		}, nil
}
