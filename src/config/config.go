package config

import (
	"os"
	"fmt"
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
/* EXPOSED FUNCTIONS */
/*-------------------*/

func Get () (Config, error) {
	var config Config
	fileName := "config.json"
	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		return config, errors.New(doesntExistError)
	}

	if !json.Valid(contents) {
		return config, errors.New("The config.json file contains invalid JSON")
	}

	err = json.Unmarshal(contents, &config)

	if err != nil {
		return config, errors.New("The config.json file for this repository could not be parsed")
	}

	if configDoesntContainAllRequiredData(config) {
		return config, errors.New("The config.json file does not contain all required data")
	}

	return config, nil
}

func CreateEmptyTemplateFile() error {
	contents := `{
	"BearerToken": "",
	"HostUrl": "",
	"RepositoryNamespace": ""
}`

	file, err := os.OpenFile("config.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return errors.New("The init command could not create the config.json file")
	}

	fmt.Fprintf(file, contents)
	return nil
}

/*-------------------*/
/* PRIVATE FUNCTIONS */
/*-------------------*/

func configDoesntContainAllRequiredData (config Config) bool {
	return  config.BearerToken == "" ||
		config.HostUrl == "" ||
		config.RepositoryNamespace == ""
}

/*-------------------*/
/* PRIVATE VARIABLES */
/*-------------------*/

var doesntExistError string = `A config.json file could not be found.

Create a config.json file at the root of your repository, with the format:-

{
 "BearerToken": <GITLAB_BEARER_TOKEN>,
 "HostUrl": <URL_TO_HOSTED_GITLAB_INSTANCE>,
 "RepositoryNamespace": <NAMESPACE>
}

- If you're using the cloud version of GitLab (the normal one), the HostUrl
  should be https://gitlab.com.
- If you're using a hosted version of GitLab, the HostUrl should be the URL
  to the hosted instance.
- The RepositoryNamespace must be of the format "<user>/<repo_name>" or 
  "<group>/<repo_name>" (ex; ArjArav98/Issue or Google/gmail)`
