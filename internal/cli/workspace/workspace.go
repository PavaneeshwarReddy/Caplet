package workspace

import "github.com/spf13/cobra"

var WsCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Manage workspaces",
	Long:  "Create, list, use, and delete workspaces with Caplet.",
}
