package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tejiriaustin/golang-llm-server/config"
	"github.com/tejiriaustin/golang-llm-server/server"
)

// rootCmd represents the base command when called without any subcommands
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the backend service",
	Run:   startServer,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startServer(c *cobra.Command, args []string) {

	conf := config.New().
		SetEnv("GOOGLE_API_KEY", config.MustGetEnv("GOOGLE_API_KEY")).
		SetEnv("PORT", config.MustGetEnv(("PORT")))

	ctx := context.Background()

	server.StartServer(ctx, conf)
}
