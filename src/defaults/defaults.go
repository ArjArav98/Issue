package defaults

import (
	"strconv"
	"github.com/ArjArav98/Issue/src/api"
)

func ArgsForCurrentUserIssues () []string {
	currentUser, err := api.GetCurrentUser()
	if err!=nil {
		return []string{}
	}

	return []string{
		"--assignee_id", strconv.FormatUint(currentUser.Id, 10),
	}
}

func ArgsForCurrentUserOpenIssues () []string {
	currentUser, err := api.GetCurrentUser()
	if err!=nil {
		return []string{}
	}

	return []string{
		"--assignee_id", strconv.FormatUint(currentUser.Id, 10),
		"--state", "opened",
	}
}
