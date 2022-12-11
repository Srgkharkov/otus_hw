package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	ExpectedResult := make(Environment, 0)
	ExpectedResult["BAR"] = EnvValue{Value: "bar", NeedRemove: false}
	ExpectedResult["EMPTY"] = EnvValue{Value: "", NeedRemove: true}
	ExpectedResult["FOO"] = EnvValue{Value: "   foo\nwith new line", NeedRemove: false}
	ExpectedResult["HELLO"] = EnvValue{Value: "\"hello\"", NeedRemove: false}
	ExpectedResult["UNSET"] = EnvValue{Value: "", NeedRemove: true}
	ReadDir, err := ReadDir("./testdata/env")
	if err != nil {
		return
	}

	t.Run("scan env dir", func(t *testing.T) {

		require.Equal(t, ExpectedResult, ReadDir)
	})
}
