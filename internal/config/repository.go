package config

import (
	"encoding/json"
	"os"
)

type Repository interface {
	Get() (*Config, error)
	Save(cfg *Config) error
}

type ConfigRepository struct {
	config string
}

func NewConfigRepository(config string) Repository {
	return &ConfigRepository{config: config}
}

func (cr *ConfigRepository) Get() (*Config, error) {

	// check whether configuration file is present or not
	_, err := os.Stat(cr.config)
	if err != nil {
		return nil, err
	}

	// read the data from the file
	data, err := os.ReadFile(cr.config)
	if err != nil {
		return nil, err
	}

	// unmarshal the data to config model
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cr *ConfigRepository) Save(cfg *Config) error {

	// marshal the json to bytes
	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	// write new configuration to the file
	return os.WriteFile(cr.config, data, 0644)

}
