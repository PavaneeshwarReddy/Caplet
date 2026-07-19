package workspace

import (
	"caplet/internal/util"
	"caplet/internal/workspace"

	"github.com/spf13/cobra"
)

var wsDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a workspace",
	Long:  "Delete a workspace which is not activated",
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceName := args[0]

		workingDir, err := util.GetWorkingDir()
		if err != nil {
			return err
		}

		workspaceRepo := workspace.NewWorkspaceRepository(workingDir)
		workspaceService := workspace.NewService(workspaceRepo, nil)
		return workspaceService.DeleteWorkspace(workspaceName)
	},
}

func init() {
	WsCmd.AddCommand(wsDeleteCmd)
}
