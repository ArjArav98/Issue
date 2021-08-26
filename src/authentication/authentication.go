package main

import (
	"fmt"
	"errors"
	"os/exec"
	"net/url"
	"runtime"
	"strconv"
	"math/rand"
)

/*-------------------*/
/* PRIVATE FUNCTIONS */
/*-------------------*/
func getClientId () string {
	return "ISSUE_OAUTH_CLIENT_ID"
}

func getClientSecret () string {
	return "ISSUE_OAUTH_CLIENT_SECRET"
}

func getRedirectUri () string {
	return "http://localhost:8090/authorized"
}

func openUrlInBrowser (url string) error {
	var currentOs string = runtime.GOOS
	var err error

	switch currentOs {
		case "darwin": err = exec.Command("open", url).Start()
		case "windows": err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "linux": err = exec.Command("xdg-open", url).Start()
		default: return errors.New("Sorry dude, this OS is not supported yet :/")
	}

	if err!=nil {
		return errors.New(fmt.Sprintf("Unable to open authentication URL in browser: %v", err))
	}

	return nil
}

func generateOauthAuthorizationUrl () (string, int) {
	hostUrl := "https://gitlab.com/oauth/authorize"
	state := rand.Int()
	queryParams := url.Values{}

	queryParams.Add("client_id", getClientId())
	queryParams.Add("redirect_uri", getRedirectUri())
	queryParams.Add("response_type", "code")
	queryParams.Add("state", strconv.Itoa(state))
	queryParams.Add("scope", "read_user,read_api,read_repository")

	return fmt.Sprintf("%v?%v", hostUrl, queryParams.Encode()), state
}

func main () {
}
