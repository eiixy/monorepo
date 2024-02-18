package app

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"os"
)

type Schedule struct {
	cron *cron.Cron
	log  *log.Helper
}

func NewSchedule(logger log.Logger) *Schedule {
	cronLog := cronLogger{log.NewHelper(log.With(logger, "module", "schedule/cron"))}
	return &Schedule{
		log: log.NewHelper(log.With(logger, "module", "schedule")),
		cron: cron.New(cron.WithChain(
			cron.Recover(cronLog),
			cron.DelayIfStillRunning(cronLog),
		)),
	}
}

type cronLogger struct {
	log *log.Helper
}

func (l cronLogger) Info(msg string, keysAndValues ...interface{}) {
	keysAndValues = append([]any{"msg", msg}, keysAndValues...)
	l.log.Infow(keysAndValues...)
}

func (l cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append([]any{"msg", msg, "err", err}, keysAndValues...)
	l.log.Errorw(keysAndValues...)
}

func (s Schedule) Add(name string, spec string, job func() error) (cron.EntryID, error) {
	s.log.Infof("add schedule: %s", name)
	return s.cron.AddFunc(spec, func() {
		s.log.Infof("run schedule: %s", name)
		err := job()
		if err != nil {
			s.log.Errorw(
				"method", "schedule_err",
				"name", name,
				"err", err,
			)
		}
	})
}

func (s Schedule) Start() error {
	if os.Getenv("SCHEDULE_ENABLE") != "false" {
		s.cron.Run()
	}
	return nil
}

func (s Schedule) Stop() error {
	s.cron.Stop()
	return nil
}
