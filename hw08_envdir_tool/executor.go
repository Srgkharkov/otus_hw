package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmdstr []string, env Environment) (returnCode int) {
	for key, mapvalue := range env {
		if mapvalue.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				return -1
			}
			continue
		}
		err := os.Setenv(key, mapvalue.Value)
		if err != nil {
			return -1
		}

	}
	cmd := exec.Command(cmdstr[0], cmdstr[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			returnCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Printf("Could not get exit code for failed program: %v", cmdstr)
			returnCode = 1
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		returnCode = ws.ExitStatus()
	}
	return
}
