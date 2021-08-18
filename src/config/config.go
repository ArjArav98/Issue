package config

import (
	"os"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	BearerToken string
	HostUrl string
	RepositoryNamespace string
}

/*-------------------*/
/* EXPOSED VARIABLES */
/*-------------------*/

var DoesntExistError error = errors.New(`A config.json file could not be found`)

/*-------------------*/
/* EXPOSED FUNCTIONS */
/*-------------------*/

func Get () (Config, error) {
	var config Config
	fileName := "config.json"
	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		return config, errors.New("The config file for this repository could not be found")
	}

	if !json.Valid(contents) {
		return config, errors.New("The config file contains invalid JSON")
	}

	err = json.Unmarshal(contents, &config)

	if err != nil {
		return config, errors.New("The config file for this repository could not be parsed")
	}

	return config, nil
}

func DoesntExist () bool {
	if _, err := os.Stat("config.json"); os.IsExist(err) {
		return false
	}
	return true
}
