package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment, 0)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		file, err := os.Open(dir + "/" + entry.Name())
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(file)
		line, _, err := reader.ReadLine()

		file.Close()

		if err != nil && err != io.EOF {
			return nil, err
		}

		line = bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))

		str := string(line)
		str = strings.TrimRight(str, " :")

		if len(str) > 0 {
			env[entry.Name()] = EnvValue{Value: str, NeedRemove: false}
		} else {
			env[entry.Name()] = EnvValue{Value: "", NeedRemove: true}
		}
	}
	return env, nil
}
