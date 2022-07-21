package cmd

import (
	"github.com/lstrgiang/ascenda/cmd/server"
	"github.com/spf13/cobra"
)

// New rootCommand

var usage = `server
	- server: start server

@ascenda
`

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "ascenda",
		Short:        "ascenda",
		Long:         usage,
		SilenceUsage: true,
	}

	cmd.AddCommand(server.NewCmd())
	return cmd
}
