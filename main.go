package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"errors"
	"os/exec"
	"github.com/ArjArav98/Issue/src/api"
	"github.com/ArjArav98/Issue/src/format"
	"github.com/ArjArav98/Issue/src/parse"
	"github.com/ArjArav98/Issue/src/config"
)

func main () {
	args := os.Args[1:]

	if noFurtherArguments(args) {
		printError(errors.New("Command needs at least one argument"))
		showHelpPrompt()
		return
	}

	/*======*/
	/* INIT */
	/*======*/
	if args[0] == "init" {
		createEmptyConfigFile()
		return
	}

	/*=========*/
	/* VERSION */
	/*=========*/
	if args[0] == "version" {
		fmt.Println("0.1.0 (Beta)")
		return
	}

	/*======*/
	/* HELP */
	/*======*/
	if args[0] == "help" {
		showHelpMenu()
		return
	}

	/*======*/
	/* SHOW */
	/*======*/
	if args[0] == "show" {
		if noFurtherArguments(args[1:]) {
			printError(errors.New("Command needs at least one more argument"))
			showHelpPrompt()
			return
		}

		if onlyOneFurtherArgument(args[1:]) {
			if argumentNotNumeric(args[1]) {
				printError(errors.New("The last argument must be a valid integer ID"))
				showHelpPrompt()
				return
			}

			// If only one argument, it must be issueID.
			showIssueWithComments(args[1], true, true)
			return
		}

		if args[1] == "--no-comments" {
			showIssueWithComments(args[2], true, false)
			return
		}
		if args[1] == "--only-comments" {
			showIssueWithComments(args[2], false, true)
			return
		}

		showIssueWithComments(args[1], true, true)
		return
	}

	/*======*/
	/* LIST */
	/*======*/
	if args[0] == "list" {
		if noFurtherArguments(args[1:]) {
			showAllIssues(args[1:])
			return
		}

		showAllIssues(args[1:])
		return
	} else {
		printError(errors.New("Command not recognised"))
		showHelpPrompt()
	}
}



/*********************/
/* COMMAND FUNCTIONS */
/*********************/

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
	var output strings.Builder

	if showIssue {
		output.Write( []byte(format.BeautifyIssue(issue)) )
	}
	if showComments {
		output.Write( []byte(format.BeautifyComments(comments)) )
	}

	pipeInputToLess(output.String())
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
	pipeInputToLess(format.BeautifyIssueList(issues))
}

func createEmptyConfigFile () {
	err := config.CreateEmptyTemplateFile()
	if err!=nil {
		printError(err)
	}
}

func showHelpMenu () {
	fmt.Println(`USAGE: issue [COMMAND] [ARGS]

+----------+
| COMMANDS |
+----------+

list	: lists all issues
	  ARGS:    Optional OPTIONS and SEARCH PARAMS.
	  OPTIONS: --my-open-issues 
	  	   --my-issues 
	  SEARCH   
	  PARAMS:  --assignee_id (integer/Any/None)
	  	   --assignee_username (comma-separated-strings)
		   --created_after (datetime)
		   --created_before (datetime)
		   --updated_after (datetime)
		   --updated_before (datetime)
		   --labels (comma-separated-strings)
		   --search (string)
		   --order_by (created_at/updated_at/)
		   --state (opened/closed)
	  SAMPLE
	  CMDS:    issue list --my-open-issues --labels backend,doing
	  	   issue list --assignee_username sauron123 --assignee_username frodo99

show	: displays an issue in detail
	  ARGS:    Optional OPTIONS and required ISSUE ID.
	  OPTIONS: --no-comments
	  	   --only-comments
	  ISSUE
	  ID:      Any positive integer
	  SAMPLE
	  CMDS:    issue show 42
	  	   issue show --no-comments 666

init	: generates an empty config in current directory
version	: displays current version
help	: dispays the help menu
`)
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

func printError (err error) {
	fmt.Printf("ERROR: %v.\n", err)
}

func showHelpPrompt () {
	fmt.Println("Enter 'issue help' to see correct usage for commands.")
}

/*------------------*/
/* OUTPUT FUNCTIONS */
/*------------------*/

func pipeInputToLess (input string) {
	cmd := exec.Command("less")
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err!=nil {
		printError(err)
	}
}
