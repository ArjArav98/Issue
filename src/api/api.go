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
	"get-issue-comments": "/projects/%v/issues/%v/notes",
	"get-current-user": "/user",
	"list-issues": "/projects/%v/issues",
}

/*-----------------------*/
/* EXPOSED API FUNCTIONS */
/*-----------------------*/

func GetIssue (issueId uint64) (types.Issue, error) {
	/*== @section ===========*/
	/*=======================*/

	var issue types.Issue
	queryParams := url.Values{}

	repository, err := GetRepositoryInformation()
	if err != nil {
		return issue, err
	}

	/*== @section ===========*/
	/*=======================*/

	url, err := generateRequestUrl("get-single-issue", queryParams,
					repository.Id, issueId)
	if err != nil {
		return issue, err
	}

	response, err := performGetRequest(url)
	if err != nil {
		return issue, err
	}
	defer response.Body.Close()

	/*== @section ===========*/
	/*=======================*/

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return issue, err
	}

	if response.StatusCode != 200 {
		return issue, errors.New(string(body))
	}

	/*== @section ===========*/
	/*=======================*/

	err = issue.FromJson(body)
	if err != nil {
		return issue, err
	}

	return issue, nil
}

func GetIssues (searchParams url.Values) ([]types.Issue, error) {
	/*== @section ===========*/
	/*=======================*/

	var issues []types.Issue

	repository, err := GetRepositoryInformation()
	if err != nil {
		return issues, err
	}

	/*== @section ===========*/
	/*=======================*/

	url, err := generateRequestUrl("list-issues", searchParams, repository.Id)
	if err != nil {
		return issues, err
	}

	response, err := performGetRequest(url)
	if err != nil {
		return issues, err
	}
	defer response.Body.Close()

	/*== @section ===========*/
	/*=======================*/

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return issues, err
	}

	if response.StatusCode != 200 {
		return issues, errors.New(string(body))
	}

	/*== @section ===========*/
	/*=======================*/

	err = types.IssuesFromJson(body, &issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}

func GetComments (issueIid uint64, repositoryId uint64) ([]types.Comment, error) {
	/*== @section ===========*/
	/*=======================*/

	var comments []types.Comment
	queryParams := url.Values{}

	queryParams.Add("sort", "asc")
	queryParams.Add("page", "1")
	queryParams.Add("per_page", "100")

	/*== @section ===========*/
	/*=======================*/

	url, err := generateRequestUrl("get-issue-comments", queryParams,
					repositoryId, issueIid)
	if err != nil {
		return comments, err
	}

	response, err := performGetRequest(url)
	if err != nil {
		return comments, err
	}
	defer response.Body.Close()

	/*== @section ===========*/
	/*=======================*/

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return comments, err
	}

	if response.StatusCode != 200 {
		return comments, errors.New(string(body))
	}

	/*== @section ===========*/
	/*=======================*/

	err = types.CommentsFromJson(body, &comments)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func GetRepositoryInformation () (types.Project, error) {
	/*== @section ===========*/
	/*=======================*/

	var project types.Project
	queryParams := url.Values{}

	config, err := config.Get()
	if err != nil {
		return project, nil
	}

	/*== @section ===========*/
	/*=======================*/

	url, err := generateRequestUrl("get-project-information", queryParams,
					url.QueryEscape(config.RepositoryNamespace))
	if err != nil {
		return project, err
	}

	response, err := performGetRequest(url)
	if err != nil {
		return project, err
	}
	defer response.Body.Close()

	/*== @section ===========*/
	/*=======================*/

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return project, err
	}

	if response.StatusCode != 200 {
		return project, errors.New(string(body))
	}

	/*== @section ===========*/
	/*=======================*/

	err = project.FromJson(body)
	if err != nil {
		return project, err
	}

	return project, nil
}

func GetCurrentUser () (types.User, error) {
	/*== @section ===========*/
	/*=======================*/

	var user types.User
	queryParams := url.Values{}

	/*== @section ===========*/
	/*=======================*/

	url, err := generateRequestUrl("get-current-user", queryParams)
	if err != nil {
		return user, err
	}

	response, err := performGetRequest(url)
	if err != nil {
		return user, err
	}
	defer response.Body.Close()

	/*== @section ===========*/
	/*=======================*/

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return user, err
	}

	if response.StatusCode != 200 {
		return user, errors.New(string(body))
	}

	/*== @section ===========*/
	/*=======================*/

	err = user.FromJson(body)
	if err != nil {
		return user, err
	}

	return user, nil
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

func generateRequestUrl (endpointKey string, queryParams url.Values, endpointSubstitutionParams ...interface{}) (string, error) {
	endpointUrl, urlPresent := gitlabApiEndpoints[endpointKey]

	if !urlPresent {
		return "", errors.New("The endpoint URL has not been configured")
	}

	host, err := hostUrl()

	if err != nil {
		return "", err
	}

	endpointUrl = fmt.Sprintf(endpointUrl, endpointSubstitutionParams...)
	return fmt.Sprintf("%v%v%v?%v", host, apiVersion, endpointUrl, queryParams.Encode()), nil
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
		return nil, errors.New("This gitlab operation was not able to be performed")
	}

	return response, nil
}
