package api

import (
	"io"
	"fmt"
	"errors"
	"net/url"
	"net/http"
	"github.com/ArjArav98/Issue/src/config"
	"github.com/ArjArav98/Issue/src/types"
)

/*------------*/
/* API CONFIG */
/*------------*/

const apiVersion string = "/api/v4"

var gitlabApiEndpoints map[string]string = map[string]string{
	"get-single-issue": "/projects/%v/issues/%v",
	"get-project-information": "/projects/%v",
}

/*-----------------------*/
/* EXPOSED API FUNCTIONS */
/*-----------------------*/

func GetIssue (issueId int) (types.Issue, error) {
	repositoryId := "8540679"
	var issue types.Issue

	// We generate the request URL.
	url, err := generateRequestUrl("get-single-issue", repositoryId, issueId)
	if err != nil {
		return issue, err
	}

	// The GET request is performed with the URL.
	response, err := performGetRequest(url)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return issue, err
	}

	// We convert the content into JSON
	// and then into an Issue struct type.
	err = issue.FromJson(body)
	if err != nil {
		return issue, err
	}

	return issue, nil
}

func GetRepositoryInformation () (types.Project, error) {
	var project types.Project

	// We get the project namespace.
	config, err := config.Get()
	if err != nil {
		return project, nil
	}

	// We generate the request URL.
	url, err := generateRequestUrl("get-project-information",
					url.QueryEscape(config.RepositoryNamespace))
	if err != nil {
		return project, err
	}

	// The GET request is performed with the URL.
	response, err := performGetRequest(url)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return project, err
	}

	// We convert the content into JSON
	// and then into an Project struct type.
	err = project.FromJson(body)
	if err != nil {
		return project, err
	}

	return project, nil
}

/*---------------------------*/
/* PRIVATE UTILITY FUNCTIONS */
/*---------------------------*/

func hostUrl () (string, error) {
	config, err := config.Get()

	if err != nil {
		return "", err
	}

	return config.HostUrl, nil
}

func generateRequestUrl (endpointKey string, endpointSubstitutionParams ...interface{}) (string, error) {
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

/*-------------------------------*/
/* PRIVATE API REQUEST FUNCTIONS */
/*-------------------------------*/

func performGetRequest (url string) (*http.Response, error) {
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
