package skill

import "github.com/spf13/cobra"

var SkillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Manage skills",
	Long:  "Create, list, use, and delete skills with Caplet.",
}
