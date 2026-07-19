package workspace

import (
	"caplet/internal/config"
	"caplet/internal/util"
	"caplet/internal/workspace"

	"github.com/spf13/cobra"
)

var wsUseCmd = &cobra.Command{
	Use:   "use <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Activate a workspace",
	Long:  "Sets the specified workspace as the active workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceName := args[0]

		workingDir, err := util.GetWorkingDir()
		if err != nil {
			return err
		}

		configPath := util.GetConfigPath(workingDir)

		workspaceRepo := workspace.NewWorkspaceRepository(workingDir)
		configRepo := config.NewConfigRepository(configPath)
		workspaceService := workspace.NewService(workspaceRepo, configRepo)
		return workspaceService.UseWorkspace(workspaceName)

	},
}

func init() {
	WsCmd.AddCommand(wsUseCmd)
}
