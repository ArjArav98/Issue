package main

import (
	"os"
	"io"
	"fmt"
	"strconv"
	"github.com/ArjArav98/Issue/src/api"
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

func getIssue (issueId int) (string, error) {
	repositoryId := "8540679"

	url, err := api.GenerateRequestUrl("get-single-issue", repositoryId, issueId)
	if err != nil {
		return "", err
	}

	response, err := api.PerformGetRequest(url)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	return string(body), nil
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

	response, err := getIssue(int(issueId))
	if err != nil {
		printError(err)
	}

	fmt.Println(response)
}
