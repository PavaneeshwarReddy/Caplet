package workspace

import (
	"caplet/internal/util"
	"os"
	"path/filepath"

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

		entries, err := os.ReadDir(workingDir)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.Name() == workspaceName {
				return ErrWorkspaceAlreadyExists
			}
		}

		if err := os.MkdirAll(filepath.Join(workingDir, workspaceName), 0755); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsCreateCmd)
}
