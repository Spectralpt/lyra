package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CommandInfo struct {
	Name        string        `json:"name"`
	Aliases     []string      `json:"aliases,omitempty"`
	Short       string        `json:"short,omitempty"`
	Long        string        `json:"long,omitempty"`
	Usage       string        `json:"usage,omitempty"`
	Flags       []FlagInfo    `json:"flags,omitempty"`
	SubCommands []CommandInfo `json:"sub_commands,omitempty"`
}

type FlagInfo struct {
	Name         string `json:"name"`
	Shorthand    string `json:"shorthand,omitempty"`
	DefaultValue string `json:"default_value,omitempty"`
	Usage        string `json:"usage"`
	Type         string `json:"type"`
	Required     bool   `json:"required"`
}

func collectCommands(cmd *cobra.Command) CommandInfo {
	if cmd.Hidden {
		return CommandInfo{}
	}

	info := CommandInfo{
		Name:        cmd.Name(),
		Aliases:     cmd.Aliases,
		Short:       cmd.Short,
		Long:        cmd.Long,
		Usage:       cmd.UseLine(),
		Flags:       collectFlags(cmd),
		SubCommands: []CommandInfo{},
	}

	for _, child := range cmd.Commands() {
		if child.Hidden {
			continue
		}
		if childInfo := collectCommands(child); childInfo.Name != "" {
			info.SubCommands = append(info.SubCommands, childInfo)
		}
	}

	return info
}

func collectFlags(cmd *cobra.Command) []FlagInfo {
	var flags []FlagInfo

	cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		if !flag.Hidden {
			flags = append(flags, createFlagInfo(flag))
		}
	})

	cmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		if !flag.Hidden {
			// Avoid duplicate flags
			if !cmd.LocalFlags().Lookup(flag.Name).Hidden {
				flags = append(flags, createFlagInfo(flag))
			}
		}
	})

	return flags
}

func createFlagInfo(flag *pflag.Flag) FlagInfo {
	required := false
	if flag.Annotations != nil {
		if _, ok := flag.Annotations[cobra.BashCompOneRequiredFlag]; ok {
			required = true
		}
	}

	return FlagInfo{
		Name:         flag.Name,
		Shorthand:    flag.Shorthand,
		Usage:        flag.Usage,
		DefaultValue: flag.DefValue,
		Type:         flag.Value.Type(),
		Required:     required,
	}
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list-commands",
		Short: "List all commands in JSON format",
		Run: func(cmd *cobra.Command, args []string) {
			root := cmd.Root()
			cmdInfo := collectCommands(root)
			jsonData, _ := json.MarshalIndent(cmdInfo, "", "  ")
			fmt.Println(string(jsonData))
		},
	})
}
