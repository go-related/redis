package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A sample application",
}

func init() {
	setupLogging()
	createCommands()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Errorf("failed to setup commands.")
		os.Exit(1)
	}
}

func setupLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
