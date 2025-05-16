package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path += string(os.PathSeparator) + configFileName
	return path, nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	contents, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}
	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
