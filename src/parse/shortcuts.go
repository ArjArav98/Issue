package parse

import (
	"strconv"
	"github.com/ArjArav98/Issue/src/api"
)

/*-------------------*/
/* EXPOSED FUNCTIONS */
/*-------------------*/

func ExpandSearchShortcuts (arguments []string) []string {
	var newArguments []string

	for _, argument := range arguments {
		if argumentIsShortcut(argument) {
			expansionFunction, _ := searchShortcutExpansionRegistry[argument]
			newArguments = append(newArguments, expansionFunction()...)
			continue
		}

		newArguments = append(newArguments, argument)
	}

	return newArguments
}

/*-----------------------------*/
/* PRIVATE EXPANSION FUNCTIONS */
/*-----------------------------*/

var searchShortcutExpansionRegistry map[string]func()[]string = map[string]func()[]string{
	"--my-issues": ArgsExpansionForCurrentUserIssues,
	"--my-open-issues": ArgsExpansionForCurrentUserOpenIssues,
}

func ArgsExpansionForCurrentUserIssues () []string {
	currentUser, err := api.GetCurrentUser()
	if err!=nil {
		return []string{}
	}

	return []string{
		"--assignee_id", strconv.FormatUint(currentUser.Id, 10),
	}
}

func ArgsExpansionForCurrentUserOpenIssues () []string {
	currentUser, err := api.GetCurrentUser()
	if err!=nil {
		return []string{}
	}

	return []string{
		"--assignee_id", strconv.FormatUint(currentUser.Id, 10),
		"--state", "opened",
	}
}

/*-------------------*/
/* PRIVATE FUNCTIONS */
/*-------------------*/

func argumentIsShortcut (argument string) bool {
	_, present := searchShortcutExpansionRegistry[argument]
	return present
}
