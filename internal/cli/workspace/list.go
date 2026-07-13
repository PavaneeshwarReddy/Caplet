package workspace

import (
	"github.com/spf13/cobra"
)

var wsListCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(1),
	Short: "List workspaces",
	Long:  "List all workspaces created in the system.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsListCmd)
}
