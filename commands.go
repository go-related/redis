package main

import (
	"fmt"
	"github.com/go-related/redis/configurations"
	"github.com/spf13/cobra"
)

func createCommands() {
	var fooCmd = &cobra.Command{
		Use:   "service1",
		Short: "Run the service1 command",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("Running service %s", configurations.ApplicationConfiguration.Service1.Name)
		},
	}

	var barCmd = &cobra.Command{
		Use:   "service2",
		Short: "Run the service2 command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Running service %s", configurations.ApplicationConfiguration.Service2.Name)
		},
	}
	rootCmd.AddCommand(fooCmd)
	rootCmd.AddCommand(barCmd)
}
