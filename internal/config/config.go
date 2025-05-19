package config

import (
	"encoding/json"
	"os"
)

func getConfigFilePath() (string, error) {
	//Actual config file stored in user home directory, not project home
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path += string(os.PathSeparator) + configFileName
	return path, nil
}

// Need capitalized to export
func Read() (Config, error) {
	//Checking that filepath is valid
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	//Get contents from file
	contents, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	//Create new instance of config file
	cfg := Config{}
	//Must use & to change the actual contents in cfg
	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
