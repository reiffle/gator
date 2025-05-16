package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DbURL             string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error { //Need pointer to actualy change struct
	if user == "" {
		return errors.New("No user specified")
	}
	c.Current_user_name = user
	err := write(*c) //could reduce to "return write(*c)""
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return errors.New("Error marshaling JSON")
	}
	path, err := getConfigFilePath()
	if err != nil {
		return errors.New("Error getting filepath for writing")
	}
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return errors.New("Couldn't write file")
	}
	return nil
}

/* ALTERNATE WRITE FROM BOOT.DEV OFFICIAL SOLUTION

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
*/
