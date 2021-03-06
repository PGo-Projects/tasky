package main

import (
	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/server"
	"github.com/spf13/cobra"
)

var (
	// TODO: Insert cobra command name here
	ServerCmd = &cobra.Command{
		Use: "<INSERT NAME HERE>",
		Run: server.MustRun,
	}
)

func init() {
	ServerCmd.PersistentFlags().StringVar(&config.Filename, "config", "",
		"config file (default is config.toml)")
	ServerCmd.PersistentFlags().BoolVar(&config.DevRun, "dev", false,
		"Run the server on a dev machine")
	cobra.OnInitialize(config.MustInit)
}

func main() {
	if err := ServerCmd.Execute(); err != nil {
		output.ErrorAndPanic(err)
	}
}
