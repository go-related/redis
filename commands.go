package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func createCommands() {
	var fooCmd = &cobra.Command{
		Use:   "foo",
		Short: "Run the foo command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running foo command")
		},
	}

	var barCmd = &cobra.Command{
		Use:   "bar",
		Short: "Run the bar command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running bar command")
		},
	}
	rootCmd.AddCommand(fooCmd)
	rootCmd.AddCommand(barCmd)
}
