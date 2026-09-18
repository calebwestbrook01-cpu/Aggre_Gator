package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	URL      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	result := filepath.Join(path, configFileName)
	return result, nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	final := Config{}
	err = json.Unmarshal(file, &final)
	if err != nil {
		return Config{}, err
	}
	return final, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	final, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, final, 0664)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	result := write(*c)
	return result
}
