/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lyra/internal"
	"lyra/types"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")
		modules, _ := cmd.Flags().GetStringSlice("modules")
		git, _ := cmd.Flags().GetBool("git")
		license, _ := cmd.Flags().GetString("license")

		project := types.Project{
			Name:     name,
			Language: language,
			Modules:  modules,
			Git:      git,
			License:  license,
		}
		err := internal.ScaffoldProject(project)
		if err != nil {
			cmd.PrintErrf("Error creating project: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "myproject", "specify project name")
	createCmd.Flags().StringP("language", "l", "myproject", "specify project language")
	createCmd.Flags().StringSliceP("modules", "m", []string{}, "specify fiware modules to install")
	createCmd.Flags().BoolP("git", "g", false, "setup git repo")
	createCmd.Flags().StringSlice("license", []string{}, "specify what license you want your project to use")
}
