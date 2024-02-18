//go:generate go run -mod=mod github.com/google/wire/cmd/wire

package main

import (
	"github.com/eiixy/monorepo/internal/app/example/cmd/migrate"
	"github.com/eiixy/monorepo/internal/app/example/conf"
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
	Name = "example"
	// Version is the version of the compiled software.
	Version = "latest"

	id, _ = os.Hostname()

	rootCmd = &cobra.Command{
		Use:     Name,
		Short:   "example",
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(migrate.Cmd)
	rootCmd.PersistentFlags().String("conf", "./configs/example.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		confPath, _ := cmd.Flags().GetString("conf")
		cfg := conf.Load(confPath)

		app, cleanup, err := initApp(newLogger(cfg.Log), cfg)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err := app.Run(); err != nil {
			panic(err)
		}
	}
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger klog.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
		kratos.Server(hs),
	)
}

func newLogger(conf config.Log) klog.Logger {
	return log.NewLoggerFromConfig(conf, Name,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)
}
