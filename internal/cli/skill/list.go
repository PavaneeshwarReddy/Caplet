package skill

import (
	"caplet/internal/config"
	"caplet/internal/skill"
	"caplet/internal/util"
	"caplet/internal/workspace"
	"path/filepath"

	"github.com/spf13/cobra"
)

var skillListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all skills",
	Long:  "Lists all skills in the activated workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
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
		skillService.ListSkills()

		return nil
	},
}

func init() {
	SkillCmd.AddCommand(skillListCmd)
}
