package main

import (
	"github.com/eiixy/monorepo/tools/y-deploy/cmd/image"
	"github.com/eiixy/monorepo/tools/y-deploy/cmd/k8s"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "y-deploy",
}

func init() {
	rootCmd.AddCommand(image.Cmd, k8s.Cmd)
	rootCmd.PersistentFlags().String("conf", "./y-deploy.yaml", "")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
