package server

import (
	"github.com/lstrgiang/ascenda/internal/app/server"
	"github.com/lstrgiang/ascenda/internal/app/server/config"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cfg := config.DefaultConfig()
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			server := server.NewServer(cfg)
			server.RegisterHandler()
			server.ProcessData()
			server.Run()
		},
	}
	cmd.Flags().StringVar(&cfg.SupplierFilePath, "supplier", cfg.SupplierFilePath, "Path to supplier file")
	cmd.Flags().StringVar(&cfg.Path, "path", cfg.Path, "API path")
	cmd.Flags().StringVar(&cfg.Host, "host", cfg.Host, "Server host")
	cmd.Flags().IntVar(&cfg.Port, "port", cfg.Port, "Server port")
	return cmd

}
