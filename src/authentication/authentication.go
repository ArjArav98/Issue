package main

import (
	"fmt"
	"errors"
	"os/exec"
	"runtime"
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

func main () {
	
}
