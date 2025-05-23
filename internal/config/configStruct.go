package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DBURL             string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error { //Need pointer to actualy change struct
	if user == "" {
		return errors.New("no user specified")
	}
	c.Current_user_name = user
	//this will write the new information to the file on disk
	return write(*c)
}

func write(cfg Config) error {
	//convert struct to json
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return errors.New("error marshaling JSON")
	}
	//get filepath that json will be written to
	path, err := getConfigFilePath()
	if err != nil {
		return errors.New("error getting filepath for writing")
	}
	//WriteFile opens and closes file automatically, 0644 will allow others to read file
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return errors.New("couldn't write file")
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
