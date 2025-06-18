package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Config contains db url and current username for home directory
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// SetUser sets the current user name in the configuration and persists the change to the configuration file.
func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

// getConfigFilePath returns the full path to the configuration file located in the user's home directory.
// It returns an error if the user's home directory cannot be determined.
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

// Read reads the configuration file, decodes its JSON content into a Config struct, and returns the Config instance.
// It returns an error if the file cannot be found, opened, or if the content fails to decode.
func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// write persists the given configuration to a JSON file in the user's home directory. Returns an error if unsuccessful.
func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
