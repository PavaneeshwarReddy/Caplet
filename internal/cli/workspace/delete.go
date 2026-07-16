package workspace

import (
	"caplet/internal/util"
	"os"
	"path/filepath"

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

		entries, err := os.ReadDir(workingDir)
		if err != nil {
			return err
		}

		var workspaceFound bool
		for _, entry := range entries {
			if entry.IsDir() && entry.Name() == workspaceName {
				workspaceFound = true
				break
			}
		}

		if !workspaceFound {
			return ErrWorkspaceNotFound
		}

		if err := os.RemoveAll(filepath.Join(workingDir, workspaceName)); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsDeleteCmd)
}
