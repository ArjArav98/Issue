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
				default:
					issueId, err := strconv.ParseInt(argsWithoutProg[1], 10, 32)
					if err != nil {
						printError(err)
					}

					response, err := showIssue(int(issueId))
					if err != nil {
						printError(err)
					}

					fmt.Println(response)
			}
		default: fmt.Println("Command not recognised.")
	}

}

func printError (err error) {
	fmt.Printf("Fatal error: %v.", err)
}

func showIssue (issueId int) (string, error) {
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
