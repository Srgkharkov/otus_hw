package hw02unpackstring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcatandrepeatrune(t *testing.T) {
	tests := []struct {
		inputstr string
		rune     rune
		count    int
		expected string
	}{
		{inputstr: "a4bc2d5e", rune: 'a', count: 3, expected: "a4bc2d5eaaa"},
		{inputstr: "a4bc2d5e", rune: 0, count: 3, expected: "a4bc2d5e"},
		{inputstr: "a4bc2d5e", rune: 0, count: 100, expected: "a4bc2d5e"},
		{inputstr: "", rune: '0', count: 10, expected: "0000000000"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.inputstr, func(t *testing.T) {
			str := tc.inputstr
			Concatandrepeatrune(&str, tc.rune, tc.count)
			require.Equal(t, tc.expected, str)
		})
	}
}
