package api

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/ArjArav98/Issue/src/config"
)

/*------------*/
/* API CONFIG */
/*------------*/

const apiVersion string = "/api/v4"

var gitlabApiEndpoints map[string]string = map[string]string{
	"get-single-issue": "/projects/%v/issues/%v",
}

/*-------------------*/
/* PRIVATE FUNCTIONS */
/*-------------------*/

func hostUrl () (string, error) {
	config, err := config.Get()

	if err != nil {
		return "", err
	}

	return config.HostUrl, nil
}

/*-------------------*/
/* EXPOSED FUNCTIONS */
/*-------------------*/

func GenerateRequestUrl (endpointKey string, endpointSubstitutionParams ...interface{}) (string, error) {
	endpointUrl, urlPresent := gitlabApiEndpoints[endpointKey]

	if !urlPresent {
		return "", errors.New("the endpoint URL has not been configured")
	}

	host, err := hostUrl()

	if err != nil {
		return "", err
	}

	endpointUrl = fmt.Sprintf(endpointUrl, endpointSubstitutionParams...)
	return fmt.Sprintf("%v%v%v", host, apiVersion, endpointUrl), nil
}

func PerformGetRequest (url string) (*http.Response, error) {
	config, err := config.Get()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", config.BearerToken))
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New("this gitlab operation was not able to be performed")
	}

	return response, nil
}
