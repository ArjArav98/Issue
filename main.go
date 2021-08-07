package main

import (
	//"os"
	"io"
	"fmt"
	//"json"
	"net/http"
)

func main () {
	//argsWithoutProg = os.Args[1:]
	fmt.Printf(showIssue(1421))
}

func showIssue (issueId int) string {
	token := "zBX_EB6cpPgv3sJ7Fm6g"
	repositoryId := "8540679"

	url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%v/issues/%v", repositoryId, issueId)

	req, err := http.NewRequest("GET", url, nil)

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
