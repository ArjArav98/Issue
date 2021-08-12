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
	"get-issue-comments": "/projects/%v/issues/%v/notes?sort=asc&page=1&per_page=100",
}

/*-----------------------*/
/* EXPOSED API FUNCTIONS */
/*-----------------------*/

func GetIssue (issueId uint64) (types.Issue, error) {
	/* Get Necessary Information */

	var issue types.Issue

	repository, err := GetRepositoryInformation()
	if err != nil {
		return issue, err
	}

	/* Request Generation and Calling */

	url, err := generateRequestUrl("get-single-issue", repository.Id, issueId)
	if err != nil {
		return issue, err
	}

	response, err := performGetRequest(url)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return issue, err
	}

	/* JSON Unmarshalling and Return */

	err = issue.FromJson(body)
	if err != nil {
		return issue, err
	}

	return issue, nil
}

func GetComments (issueIid uint64, repositoryId uint64) ([]types.Comment, error) {
	var comments []types.Comment

	/* Request Generation and Calling */

	url, err := generateRequestUrl("get-issue-comments", repositoryId, issueIid)
	if err != nil {
		return comments, err
	}

	response, err := performGetRequest(url)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return comments, err
	}

	/* JSON Unmarshalling and Return */

	err = types.CommentsFromJson(body, &comments)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func GetRepositoryInformation () (types.Project, error) {
	/* Get Necessary Information */

	var project types.Project

	config, err := config.Get()
	if err != nil {
		return project, nil
	}

	/* Request Generation and Calling */

	url, err := generateRequestUrl("get-project-information",
					url.QueryEscape(config.RepositoryNamespace))
	if err != nil {
		return project, err
	}

	response, err := performGetRequest(url)
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return project, errors.New("the repository information could not be retrieved")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return project, err
	}

	/* JSON Unmarshalling and Return */

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
