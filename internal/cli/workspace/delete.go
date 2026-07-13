package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wsDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a workspace",
	Long:  "Delete a workspace which is not activated",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Println("Workspace name", name)
		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsDeleteCmd)
}
