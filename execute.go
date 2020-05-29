package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

func executeCommand(command ...string) (err error) {

	//Fromat command
	cmd := exec.Command(command[0], command[1:]...)

	//Assign error listner
	var stderr io.ReadCloser
	stderr, err = cmd.StderrPipe()
	if err != nil {
		return
	}

	//Assign output listner
	var stdout io.ReadCloser
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return
	}

	//Start the command exec
	if err = cmd.Start(); err != nil {
		return
	}

	//Emit output (if any)
	slurpout, _ := ioutil.ReadAll(stdout)
	slurpoutstr := strings.TrimSpace(string(slurpout))
	if slurpoutstr != "" {
		fmt.Printf("%s\n", slurpoutstr)
	}

	//Emit error (if any)
	slurperr, _ := ioutil.ReadAll(stderr)
	slurperrstr := strings.TrimSpace(string(slurperr))
	if slurperrstr != "" {
		fmt.Printf("%s\n", slurperrstr)
	}

	// Wait till the command is completed/exited
	if err = cmd.Wait(); err != nil {
		return
	}
	return
}
