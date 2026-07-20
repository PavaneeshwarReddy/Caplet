package cli

import (
	"caplet/internal/cli/skill"
	"caplet/internal/cli/workspace"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caplet",
	Short: "A CLI for managing skills",
	Long:  "Caplet helps you create, manage, and discover skills across environments.",
}

func init() {
	rootCmd.AddCommand(workspace.WsCmd)
	rootCmd.AddCommand(skill.SkillCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
