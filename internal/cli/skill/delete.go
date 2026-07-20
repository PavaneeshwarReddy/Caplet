package skill

import (
	"caplet/internal/config"
	"caplet/internal/skill"
	"caplet/internal/util"
	"caplet/internal/workspace"
	"path/filepath"

	"github.com/spf13/cobra"
)

var skillDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a skill",
	Long:  "Delete a skill",
	RunE: func(cmd *cobra.Command, args []string) error {
		skillName := args[0]

		workingDir, err := util.GetWorkingDir()
		if err != nil {
			return err
		}

		configDir := util.GetConfigPath(workingDir)

		configRepo := config.NewConfigRepository(configDir)
		workspaceRepo := workspace.NewWorkspaceRepository(workingDir)
		workspaceService := workspace.NewService(workspaceRepo, configRepo)

		activeWorkspace, err := workspaceService.GetActiveWorkspace()
		if err != nil {
			return err
		}

		skillWorkingDir := filepath.Join(workingDir, activeWorkspace)

		skillRepository := skill.NewSkillRepository(skillWorkingDir)
		skillService := skill.NewSkillService(skillRepository)
		skillService.DeleteSkill(skillName)

		return nil
	},
}

func init() {
	SkillCmd.AddCommand(skillDeleteCmd)
}
