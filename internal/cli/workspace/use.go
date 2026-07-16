package workspace

import (
	"caplet/internal/config"
	"caplet/internal/util"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var wsUseCmd = &cobra.Command{
	Use:   "use <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Activate a workspace",
	Long:  "Sets the specified workspace as the active workspace",
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

		configPath := util.GetConfigPath(workingDir)
		_, err = os.Stat(configPath)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(configPath)
		if err != nil {
			return err
		}

		var cfg config.Config
		if err = json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		cfg.ActiveWorkspace = workspaceName
		data, err = json.MarshalIndent(cfg, "", " ")
		if err != nil {
			return err
		}

		if err = os.WriteFile(configPath, data, 0644); err != nil {
			return err
		}

		return nil

	},
}

func init() {
	WsCmd.AddCommand(wsUseCmd)
}
