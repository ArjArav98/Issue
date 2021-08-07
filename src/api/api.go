package api

import (
	"fmt"
	"errors"
	"github.com/ArjArav98/Issue/src/config"
)

const apiVersion string = "/api/v4"

var endpoints map[string]string = map[string]string{
	"get-single-issue": "/projects/%v/issues/%v",
}

func hostUrl () (string, error) {
	config, err := config.Get()

	if err != nil {
		return "", err
	}

	return config.HostUrl, nil
}

func GenerateUrl (endpointKey string) (string, error) {
	endpointUrl, urlPresent := endpoints[endpointKey]

	if !urlPresent {
		return "", errors.New("the endpoint URL has not been configured")
	}

	host, err := hostUrl()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v%v%v", host, apiVersion, endpointUrl), nil
}
