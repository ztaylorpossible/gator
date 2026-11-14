package config

import (
	"os"
	"encoding/json"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath();
	if err != nil {
		return Config{}, err;
	}

	data, err := os.ReadFile(path);
	if err != nil {
		return Config{}, err;
	}

	var c Config;
	err = json.Unmarshal(data, &c);
	if err != nil {
		return Config{}, err;
	}

	return c, nil;
}

func (c Config) SetUser(username string) error {
	c.Current_user_name = username;

	err := write(c);
	if err != nil {
		return err;
	}

	return nil;
}

const configFileName = ".gatorconfig.json";

func write(c Config) error {
	data, err := json.Marshal(c);
	if err != nil {
		return err;
	}

	path, err := getConfigFilePath();
	if err != nil {
		return err;
	}

	err = os.WriteFile(path, data, 0666);
	if err != nil {
		return err;
	}

	return nil;
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir();
	if err != nil {
		return "", err
	}
	path += "/.gatorconfig.json";

	return path, nil;
}
