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
	}

	issue, err := api.GetIssue(int(issueId))
	if err != nil {
		printError(err)
	}

	fmt.Println(format.BeautifyIssue(issue))
}

/*-------------------*/
/* UTILITY FUNCTIONS */
/*-------------------*/

func printError (err error) {
	fmt.Printf("Fatal error: %v.", err)
}
