package schedule

import (
	"context"
	"fmt"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/eiixy/monorepo/internal/pkg/app"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSchedule, NewDaily)

type Schedule struct {
	*app.Schedule
	log   *log.Helper
	ent   *ent.Client
	daily *Daily
}

func NewSchedule(logger log.Logger, data *data.Data, daily *Daily) *Schedule {
	s := app.NewSchedule(logger)
	return &Schedule{
		Schedule: s,
		ent:      data.EntClient,
		log:      log.NewHelper(log.With(logger, "module", "schedule")),
		daily:    daily,
	}
}

func (s Schedule) Start(ctx context.Context) error {
	_, _ = s.Add("example", "* * * * *", s.example)
	_, _ = s.Add("test", "* * * * *", test)
	_, _ = s.Add("daily", "* * * * *", func() error {
		return s.daily.Run(ctx)
	})
	return s.Schedule.Start()
}

func (s Schedule) Stop(ctx context.Context) error {
	return s.Schedule.Stop()
}

func (s Schedule) example() error {
	s.log.Info("todo")
	return nil
}

func test() error {
	fmt.Println("test")
	return nil
}
