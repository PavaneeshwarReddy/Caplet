package config

import (
	"caplet/internal/util"
	"encoding/json"
	"os"
)

func Init() error {

	// create a working director if not exists
	workingDir, err := util.GetWorkingDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(workingDir, 0755); err != nil {
		return err
	}

	// create required config file
	configPath := util.GetConfigPath(workingDir)
	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(configPath)
			if err != nil {
				return err
			}
			config := Config{
				ActiveWorkspace: "default",
			}
			data, err := json.MarshalIndent(config, "", " ")
			if err != nil {
				return err
			}
			err = os.WriteFile(configPath, data, 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Get() (*Config, error) {
	workingDir, err := util.GetWorkingDir()
	if err != nil {
		return nil, err
	}

	configPath := util.GetConfigPath(workingDir)

	_, err = os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
