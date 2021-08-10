package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/ArjArav98/Issue/src/api"
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

/*-----------------------*/
/* API COMMAND FUNCTIONS */
/*-----------------------*/

func showIssueWithComments (issueIdString string) {
	// We try to convert the passed parameter to an int issueId.
	issueId, err := strconv.ParseInt(issueIdString, 10, 32)
	if err != nil {
		printError(err)
		return
	}

	issue, err := api.GetIssue(uint64(issueId))
	if err != nil {
		printError(err)
		return
	}

	comments, err := api.GetComments(issue.Iid, issue.ProjectId)
	if err != nil {
		printError(err)
		return
	}

	fmt.Println(format.BeautifyIssue(issue))
	fmt.Println(comments)
}

/*-------------------*/
/* UTILITY FUNCTIONS */
/*-------------------*/

func printError (err error) {
	fmt.Printf("Fatal error: %v.\n", err)
}
