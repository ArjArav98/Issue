package main

import (
	"os"
	"fmt"
	"strconv"
	"errors"
	"github.com/ArjArav98/Issue/src/api"
	"github.com/ArjArav98/Issue/src/format"
	"github.com/ArjArav98/Issue/src/parse"
)

func main () {
	args := os.Args[1:]

	if noFurtherArguments(args) {
		printError(errors.New("Command needs at least one argument"))
		return
	}

	if args[0] == "show" {
		if noFurtherArguments(args[1:]) {
			printError(errors.New("Command needs at least one more argument"))
			return
		}

		if onlyOneFurtherArgument(args[1:]) {
			if argumentNotNumeric(args[1]) {
				printError(errors.New("The last argument must be a valid integer ID"))
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
		return
	}

	if args[0] == "list" {
		if noFurtherArguments(args[1:]) {
			showAllIssues(args[1:])
			return
		}

		showAllIssues(args[1:])
	} else {
		printError(errors.New("Command not recognised"))
	}
}

/*-----------------------*/
/* API COMMAND FUNCTIONS */
/*-----------------------*/

func showIssueWithComments (issueIdString string, showIssue bool, showComments bool) {
	/*== @section ===========*/
	/*=======================*/

	issueId, err := strconv.ParseInt(issueIdString, 10, 32)
	if err != nil {
		printError(errors.New("Issue ID must be a valid number"))
		return
	}

	/*== @section ===========*/
	/*=======================*/

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

	/*== @section ===========*/
	/*=======================*/

	if showIssue {
		fmt.Println(format.BeautifyIssue(issue))
	}
	if showComments {
		fmt.Println(format.BeautifyComments(comments))
	}
}

func showAllIssues (searchArgs []string) {
	/*== @section ===========*/
	/*=======================*/

	queryParams, err := parse.CliArgumentsToQueryParams(searchArgs)
	if err!=nil {
		printError(err)
		return
	}

	/*== @section ===========*/
	/*=======================*/

	issues, err := api.GetIssues(queryParams)
	if err!= nil {
		printError(err)
		return
	}

	/*== @section ===========*/
	/*=======================*/

	fmt.Println(issues)
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
