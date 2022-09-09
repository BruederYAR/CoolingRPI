package input

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func ReadConfigurationFile(name string) (*Config, error) {
	if !FileExist(name) {
		return nil, errors.New("file not exist")
	}

	fileData, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var data *Config
	err = json.Unmarshal(fileData, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadFile(name string) (string, error) {
	if !FileExist(name) {
		return "", errors.New("file not exist")
	}

	fileData, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}

	return string(fileData), nil
}

func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return false
}
