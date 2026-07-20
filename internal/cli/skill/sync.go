package skill

import (
	"caplet/internal/skill"
	"caplet/internal/util"

	"github.com/spf13/cobra"
)

var skillSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs all skills",
	Long:  "Syncs all skills and validates them",
	RunE: func(cmd *cobra.Command, args []string) error {

		workingDir, err := util.GetWorkingDir()
		if err != nil {
			return err
		}

		skillRepository := skill.NewSkillRepository(workingDir)
		skillService := skill.NewSkillService(skillRepository)
		skillService.SyncSkills()

		return nil
	},
}

func init() {
	SkillCmd.AddCommand(skillSyncCmd)
}
