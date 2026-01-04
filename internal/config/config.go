package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_username"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUsername = userName
	return write(*c)
}

func Read() (Config, error) {
	var zero Config

	fullPath, err := getConfigFilePath()
	if err != nil {
		return zero, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return zero, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	err = decoder.Decode(&cfg)
	if err != nil {
		return zero, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	var zero string
	home, err := os.UserHomeDir()
	if err != nil {
		return zero, err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
