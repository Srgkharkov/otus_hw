package main

import (
	"errors"
	"log"
	"os"
)

var (
	ErrNeedMinTwoArgs = errors.New("Need minimum two arguments: path and command")
	ErrReadDir        = errors.New("Can not read directory")
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 2 {
		log.Fatal(ErrNeedMinTwoArgs)
	}

	env, err := ReadDir(argsWithoutProg[0])
	if err != nil {
		log.Fatal(ErrReadDir, err)
	}
	os.Exit(RunCmd(argsWithoutProg[1:], env))
}
