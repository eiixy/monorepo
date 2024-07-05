package data

import (
	"github.com/IBM/sarama"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	_ "github.com/eiixy/monorepo/internal/data/example/ent/runtime"
	"github.com/eiixy/monorepo/pkg/database"
	"github.com/eiixy/monorepo/pkg/kafka"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,
	NewEntDatabase,
	NewKafkaProducer,
)

type Data struct {
	client *ent.Client
	db     *ent.Database
}

func NewData(client *ent.Client, db *ent.Database, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	d := &Data{
		client: client,
		db:     db,
	}
	return d, func() {
		if err := d.client.Close(); err != nil {
			l.Error(err)
		}
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

func NewEntClient(cfg *conf.Config) (*ent.Client, error) {
	drv, err := database.NewEntDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewClient(ent.Driver(drv)), nil
}

func NewEntDatabase(cfg *conf.Config) (*ent.Database, error) {
	drv, err := database.NewEntDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewDatabase(ent.Driver(drv)), nil
}

func NewKafkaProducer(cfg *conf.Config) (sarama.SyncProducer, error) {
	return kafka.NewSyncProducerFromConfig(cfg.Data.Kafka)
}
