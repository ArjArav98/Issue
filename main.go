package main

import (
	//"os"
	"io"
	"fmt"
	//"json"
	"net/http"
	"github.com/ArjArav98/Issue/src/config"
	"github.com/ArjArav98/Issue/src/api"
)

func main () {
	//argsWithoutProg = os.Args[1:]
	fmt.Printf(showIssue(1421))
}

func showIssue (issueId int) string {
	configs, err := config.Get()

	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	token := configs.BearerToken
	repositoryId := "8540679"

	url, err := api.GenerateUrl("get-single-issue")

	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(url, repositoryId, issueId), nil)

	if err != nil {
		return "oops"
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "oops2"
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body)
}
