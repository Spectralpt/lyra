/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listLangsCmd represents the listLangs command
var listLangsCmd = &cobra.Command{
	Use:   "listLangs",
	Short: "Lists supported languages in json format",
	Long:  `This command lists all the supported languages in a json format, this is for internal use of the web version of this tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("{[c,cpp]}")
	},
}

func init() {
	rootCmd.AddCommand(listLangsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listLangsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listLangsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
