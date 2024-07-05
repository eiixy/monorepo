package job

import (
	"context"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/pkg/app"
	"github.com/eiixy/monorepo/pkg/kafka"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewJob)

type Job struct {
	*app.Job
}

func NewJob(logger log.Logger, cfg *conf.Config, data *data.Data) *Job {
	works := []*app.Worker{
		app.NewWorker("example", &exampleJob{}),
		app.NewWorker("kafkaConsumer", &kafkaConsumer{
			cg:  kafka.NewConsumerGroupFromConfig(cfg.Data.Kafka, cfg.KafkaConsumerGroup),
			log: log.NewHelper(log.With(logger, "module", "job/kafka_consumer")),
		}),
	}
	j := app.NewJob(logger, works...)
	return &Job{j}
}

func (j Job) Start(ctx context.Context) error {
	return j.Job.Start(ctx)
}

func (j Job) Stop(ctx context.Context) error {
	return nil
}

type exampleJob struct {
}

func (e exampleJob) Run(ctx context.Context) error {
	panic("todo")
}
