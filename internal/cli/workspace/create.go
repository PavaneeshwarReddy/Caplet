package workspace

import (
	"caplet/internal/util"
	"caplet/internal/workspace"

	"github.com/spf13/cobra"
)

var wsCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new workspace",
	Long:  "Creates a new workspace without activating it.",
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceName := args[0]

		workingDir, err := util.GetWorkingDir()
		if err != nil {
			return err
		}

		workspaceRepo := workspace.NewWorkspaceRepository(workingDir)
		workspaceService := workspace.NewService(workspaceRepo, nil)
		return workspaceService.CreateWorkspace(workspaceName)

	},
}

func init() {
	WsCmd.AddCommand(wsCreateCmd)
}
