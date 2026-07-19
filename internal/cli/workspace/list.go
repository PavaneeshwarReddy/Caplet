package workspace

import (
	"caplet/internal/util"
	"caplet/internal/workspace"

	"github.com/spf13/cobra"
)

var wsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces",
	Long:  "List all workspaces created in the system.",
	RunE: func(cmd *cobra.Command, args []string) error {
		workingDir, err := util.GetWorkingDir()
		if err != nil {
			return err
		}

		workspaceRepo := workspace.NewWorkspaceRepository(workingDir)
		worspsaceService := workspace.NewService(workspaceRepo, nil)
		return worspsaceService.ListWorkspaces()
	},
}

func init() {
	WsCmd.AddCommand(wsListCmd)
}
