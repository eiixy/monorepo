package k8s

import (
	"github.com/eiixy/monorepo/tools/y-deploy/cmd/k8s/deployment"
	"github.com/eiixy/monorepo/tools/y-deploy/cmd/k8s/gen"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "k8s",
}

func init() {
	Cmd.AddCommand(deployment.Cmd, gen.Cmd)
}
