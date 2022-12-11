package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	env, _ := ReadDir("./testdata/env")
	command := []string{"ls", "testdata/env"}
	RunCmd(command, env)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	StrResult := buf.String()
	StrExpected := "BAR\nEMPTY\nFOO\nHELLO\nUNSET\n"
	require.Equal(t, StrExpected, StrResult)

}
