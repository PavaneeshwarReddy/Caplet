package workspace

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wsCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new workspace",
	Long:  "Creates a new workspace without activating it.",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		fmt.Println("Workspace name", name)
		return nil
	},
}

func init() {
	WsCmd.AddCommand(wsCreateCmd)
}
