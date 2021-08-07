package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	BearerToken string
	HostUrl string
	RepositoryName string
}

func Get () (Config, error) {
	var config Config
	fileName := "config.json"
	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		return config, errors.New("the config file for this repository could not be found")
	}

	if !json.Valid(contents) {
		return config, errors.New("the config file contains invalid JSON")
	}

	err = json.Unmarshal(contents, &config)

	if err != nil {
		return config, errors.New("the config file for this repository could not be parsed")
	}

	return config, nil
}
