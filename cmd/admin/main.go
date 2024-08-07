//go:generate go run -mod=mod github.com/google/wire/cmd/wire

package main

import (
	"github.com/eiixy/monorepo/internal/app/admin/cmd/migrate"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/app/admin/job"
	"github.com/eiixy/monorepo/internal/app/admin/schedule"
	"github.com/eiixy/monorepo/internal/pkg/config"
	"github.com/eiixy/monorepo/pkg/log"
	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/cobra"
	"os"
)

var (
	// Name is the name of the compiled software.
	Name = "admin"
	// Version is the version of the compiled software.
	Version = "latest"

	id, _ = os.Hostname()

	rootCmd = &cobra.Command{
		Use:     Name,
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(migrate.Cmd)
	rootCmd.PersistentFlags().String("conf", "./configs/admin.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		confPath, _ := cmd.Flags().GetString("conf")
		cfg, err := config.Load[conf.Config](confPath)
		if err != nil {
			panic(err)
		}
		logger := log.NewLoggerFromConfig(cfg.Log, Name,
			log.ServiceID(id),
			log.ServiceName(Name),
			log.ServiceVersion(Version),
		)
		app, cleanup, err := initApp(logger, cfg)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err = app.Run(); err != nil {
			panic(err)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger klog.Logger, hs *http.Server, job *job.Job, schedule *schedule.Schedule) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
		kratos.Server(hs, job, schedule),
	)
}
