package cmd

import (
	"github.com/kvitrvn/githeim/internal/app"
	"github.com/kvitrvn/githeim/internal/config"
	"github.com/spf13/cobra"
)

var port int
var configPath string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			return err
		}

		return app.Run(cmd.Context(), &cfg)
	},
}

func init() {
	serveCmd.Flags().IntVar(&port, "port", 8080, "Port to listen on")
	serveCmd.Flags().StringVar(&configPath, "config", ".", "Path to the config file")
	rootCmd.AddCommand(serveCmd)
}
