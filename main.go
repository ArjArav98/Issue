package main

import (
	//"os"
	"io"
	"fmt"
	//"json"
	"github.com/ArjArav98/Issue/src/api"
)

func main () {
	//argsWithoutProg = os.Args[1:]
	fmt.Printf(showIssue(1421))
}

func showIssue (issueId int) string {
	repositoryId := "8540679"

	url, err := api.GenerateRequestUrl("get-single-issue", repositoryId, issueId)

	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	response, err := api.PerformGetRequest(url)

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	return string(body)
}
