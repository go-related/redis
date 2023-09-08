package main

import (
	"github.com/go-related/redis/service1"
	"github.com/go-related/redis/service2"
	"github.com/go-related/redis/settings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func createCommands() {
	var fooCmd = &cobra.Command{
		Use:   "service1",
		Short: "Run the service1 command",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := service1.InitService(settings.ApplicationConfiguration.Service1)
			if err != nil {
				logrus.Error(err)
			}
		},
	}

	var barCmd = &cobra.Command{
		Use:   "service2",
		Short: "Run the service2 command",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := service2.InitService(settings.ApplicationConfiguration.Service2)
			if err != nil {
				logrus.Error(err)
			}
		},
	}
	rootCmd.AddCommand(fooCmd)
	rootCmd.AddCommand(barCmd)
}
