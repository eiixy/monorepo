package migrate

import (
	"context"
	"fmt"
	"github.com/eiixy/monorepo/internal/app/account/conf"
	"github.com/eiixy/monorepo/internal/app/account/data"
	"github.com/eiixy/monorepo/internal/data/account/ent"
	"github.com/eiixy/monorepo/internal/pkg/config"
	"github.com/spf13/cobra"
	"log"
)

var Cmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg conf.Config
		path, _ := cmd.Flags().GetString("conf")
		config.Load(path, &cfg)
		client, err := data.NewEntClient(&cfg)
		if err != nil {
			panic(err)
		}
		defer func(client *ent.Client) {
			err := client.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(client)
		// Run the auto migration tool.
		if err := client.Debug().Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	},
}
