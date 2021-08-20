package parse

import (
	"errors"
	"strings"
	"net/url"
	"github.com/ArjArav98/Issue/src/defaults"
)

/*-------------------*/
/* PRIVATE VARIABLES */
/*-------------------*/

var SearchShortcutExpansionFunctions map[string]func()[]string = map[string]func()[]string{
	"--my-issues": defaults.ArgsForCurrentUserIssues,
	"--my-open-issues": defaults.ArgsForCurrentUserOpenIssues,
}

/*-------------------*/
/* EXPOSED FUNCTIONS */
/*-------------------*/

func CliArgumentsToQueryParams (arguments []string) (url.Values, error){
	arguments = expandSearchShortcuts(arguments)
	queryParams := url.Values{}

	if len(arguments)%2 != 0 {
		return queryParams, errors.New("A search parameter is missing a value")
	}

	/* We iterate over each key,value pair at a time.*/
	for iter:=0; iter<len(arguments); iter+=2 {
		key := arguments[iter]
		value := arguments[iter+1]

		if !strings.HasPrefix(key, "--") {
			return queryParams, errors.New("A search parameter is not prefixed with --.")
		}

		// We take off the --.
		queryParams.Add(key[2:], value)
	}

	return queryParams, nil
}

/*-------------------*/
/* PRIVATE FUNCTIONS */
/*-------------------*/

func expandSearchShortcuts (arguments []string) []string {
	var newArguments []string

	for _, argument := range arguments {
		if argumentIsShortcut(argument) {
			expansionFunction, _ := SearchShortcutExpansionFunctions[argument]
			newArguments = append(newArguments, expansionFunction()...)
			continue
		}

		newArguments = append(newArguments, argument)
	}

	return newArguments
}

func argumentIsShortcut (argument string) bool {
	_, present := SearchShortcutExpansionFunctions[argument]
	return present
}

func argumentIsSearchKey (argument string) bool {
	return strings.HasPrefix(argument, "--")
}

func argumentIsSearchValue (argument string) bool {
	return !strings.HasPrefix(argument, "--")
}
