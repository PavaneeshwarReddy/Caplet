package workspace

import (
	"caplet/internal/util"
	"log"
	"os"

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

		entries, err := os.ReadDir(workingDir)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				log.Println(entry)
			}
		}

		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsListCmd)
}
