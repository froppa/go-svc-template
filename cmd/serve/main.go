package main

import (
	"github.com/froppa/go-svc-template/internal/repo"
	"github.com/froppa/go-svc-template/internal/service"
	"github.com/froppa/go-svc-template/internal/transport"
	"github.com/froppa/go-svc-template/pkg/config"
	"github.com/froppa/go-svc-template/pkg/logger"
	"github.com/froppa/go-svc-template/pkg/server/httpfx"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var serveHTTPCmd = &cobra.Command{
	Use: "http",

	Short: "Serve HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config.RegisterFlags(cmd)
		app := fx.New(
			config.Module,
			httpfx.Module,
			logger.Module,
			repo.Module,
			service.Module,
			fx.Invoke(transport.NewHTTPHandler),
		)
		app.Run()
	},
}

func main() {
	rootCmd.AddCommand(serveHTTPCmd)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
