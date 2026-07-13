package workspace

import (
	"github.com/spf13/cobra"
)

var wsUseCmd = &cobra.Command{
	Use:   "use <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Activate a workspace",
	Long:  "Sets the specified workspace as the active workspace",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsUseCmd)
}
