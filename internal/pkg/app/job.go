package app

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/time/rate"
	"os"
	"runtime"
	"time"
)

type Job struct {
	jobs  map[string]JobHandler
	works []*Worker
	log   *log.Helper
}

func NewJob(logger log.Logger, works ...*Worker) *Job {
	return &Job{
		log:   log.NewHelper(log.With(logger, "module", "job")),
		works: works,
	}
}

type Worker struct {
	name    string
	job     JobHandler
	limiter *rate.Limiter
}

type JobHandler interface {
	Run(ctx context.Context) error
}

func NewWorker(name string, job JobHandler) *Worker {
	limiter := rate.NewLimiter(rate.Every(time.Hour), 5)
	w := &Worker{name: name, job: job, limiter: limiter}
	return w
}

func (j Job) Start(ctx context.Context) error {
	for _, item := range j.works {
		go j.exec(ctx, item)
	}
	select {}
}

func (j Job) run(ctx context.Context, work *Worker) {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			j.log.Error(err, "panic", "stack", "...\n"+string(buf))
		}
	}()
	err := work.job.Run(ctx)
	if err != nil {
		j.log.Error(err)
	}
}

func (j Job) exec(ctx context.Context, work *Worker) {
	for {
		if work.limiter.Allow() {
			if os.Getenv("JOB_ENABLE") != "false" {
				fmt.Println("start job:", work.name)
				j.run(ctx, work)
			}
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}
