package main

import (
	"os"
	"fmt"
	"errors"
	"os/exec"
	"runtime"
)

func openUrlInBrowser (url string) err {
	var currentOs string = runtime.GOOS
	var openCmd *runtime.Cmd

	switch currentOs {
		case "darwin": openCmd = exec.Command("open")
		case "windows": openCmd = exec.Command("open")
		case "linux": openCmd = exec.Command("open")
		default: return errors.New("Sorry dude, this OS is not supported yet :/")
	}

	openCmd.Stdin = strings.NewReader(url)
	openCmd.Stdout = os.Stdout
	err := openCmd.Run()
	if err!=nil {
		return errors.New("Unable to open authentication URL in browser")
	}

	return nil
}
