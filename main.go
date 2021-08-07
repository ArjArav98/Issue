package main

import (
	"os"
	"io"
	"fmt"
	"strconv"
	"github.com/ArjArav98/Issue/src/api"
	"github.com/ArjArav98/Issue/src/types"
	"github.com/ArjArav98/Issue/src/format"
)

func main () {
	argsWithoutProg := os.Args[1:]

	switch argsWithoutProg[0] {
		case "show":
			switch argsWithoutProg[1] {
				default: showIssueWithComments(argsWithoutProg[1])
			}
		default: fmt.Println("Command not recognised.")
	}

}

/*-------------------*/
/* UTILITY FUNCTIONS */
/*-------------------*/

func printError (err error) {
	fmt.Printf("Fatal error: %v.", err)
}

/*---------------*/
/* API FUNCTIONS */
/*---------------*/

func getIssue (issueId int) (types.Issue, error) {
	repositoryId := "8540679"
	var issue types.Issue

	// We generate the request URL.
	url, err := api.GenerateRequestUrl("get-single-issue", repositoryId, issueId)
	if err != nil {
		return issue, err
	}

	// The GET request is performed with the URL.
	response, err := api.PerformGetRequest(url)
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

/*-----------------------*/
/* API COMMAND FUNCTIONS */
/*-----------------------*/

func showIssueWithComments (issueIdString string) {
	// We try to convert the passed parameter to an int issueId.
	issueId, err := strconv.ParseInt(issueIdString, 10, 32)
	if err != nil {
		printError(err)
	}

	issue, err := getIssue(int(issueId))
	if err != nil {
		printError(err)
	}

	fmt.Println(format.BeautifyIssue(issue))
}
