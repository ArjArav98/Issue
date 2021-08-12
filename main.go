package main

import (
	"os"
	"fmt"
	"strconv"
	"errors"
	"github.com/ArjArav98/Issue/src/api"
	"github.com/ArjArav98/Issue/src/format"
)

func main () {
	args := os.Args[1:]

	if noFurtherArguments(args) || onlyOneFurtherArgument(args) {
		printError(errors.New("Command not recognised"))
		return
	}

	if args[0] == "show" {
		if onlyOneFurtherArgument(args[1:]) {
			if argumentNotNumeric(args[1]) {
				printError(errors.New("Command not recognised"))
				return
			}

			// If only one argument, it must be issueID.
			showIssueWithComments(args[1], true, true)
			return
		}

		if args[1] == "--no-comments" {
			showIssueWithComments(args[2], true, false)
		}
		if args[1] == "--only-comments" {
			showIssueWithComments(args[2], false, true)
		}

		printError(errors.New("Command not recognised"))
	} else {
		printError(errors.New("Command not recognised"))
	}
}

/*-----------------------*/
/* API COMMAND FUNCTIONS */
/*-----------------------*/

func showIssueWithComments (issueIdString string, showIssue bool, showComments bool) {
	// We try to convert the passed parameter to an int issueId.
	issueId, err := strconv.ParseInt(issueIdString, 10, 32)
	if err != nil {
		printError("Issue ID must be a valid number")
		return
	}

	// We call the API functions.
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

	// We format and display them.
	if showIssue {
		fmt.Println(format.BeautifyIssue(issue))
	}
	if showComments {
		fmt.Println(format.BeautifyComments(comments))
	}
}

func showAllIssues (extraSearchArgs []string) {
	
}

/*-----------------*/
/* CHECK FUNCTIONS */
/*-----------------*/

func noFurtherArguments (args []string) bool {
	return len(args) == 0
}

func onlyOneFurtherArgument (args []string) bool {
	return len(args) == 1
}

func argumentNotNumeric (argument string) bool {
	_, err := strconv.ParseFloat(argument, 10)
	return err != nil
}

/*-----------------*/
/* ERROR FUNCTIONS */
/*-----------------*/

func showHelpMenu () {
	fmt.Println("Usage: issue COMMAND [OPTION] [ISSUE_ID]")
}

func printError (err error) {
	fmt.Printf("ERROR: %v.\n", err)
}
