/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"lyra/cmd/ui"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lyra",
	Short: "A Fiware project creator",
	Long: `Lyra is fiware project creator that allows you to create:
	Fiware Context Broker
	IoT Agent`,
	Run: func(cmd *cobra.Command, args []string) {
		isInteractive, _ := cmd.Flags().GetBool("interactive")
		if isInteractive {
			ui.RunInteractive()
		} else {
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.Flags().BoolP("interactive", "i", false, "Run lyra in interactive mode")
}
