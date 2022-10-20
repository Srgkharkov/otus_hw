package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCopy(t *testing.T) {
	t.Run("Successful copy", func(t *testing.T) {
		fsrc, err := os.OpenFile("./testdata/input.txt", os.O_RDONLY, 0755)
		assert.Nil(t, err, "can not CreateTemp")
		defer fsrc.Close()

		fdst, err := os.CreateTemp("", "sample")
		assert.Nil(t, err, "can not CreateTemp")
		defer os.Remove(fdst.Name())

		err = Copy(fsrc.Name(), fdst.Name(), 0, 0)
		assert.Nil(t, err, "func Copy return error")

		srcbytes, err := os.ReadFile(fsrc.Name())
		assert.Nil(t, err, "can not read source file")

		destbytes, err := os.ReadFile(fdst.Name())
		assert.Nil(t, err, "can not read destination file")

		require.True(t, Equal(srcbytes, destbytes), "not equal source file and destination file")
	})

	t.Run("Offset more than size file", func(t *testing.T) {
		fsrc, err := os.OpenFile("./testdata/input.txt", os.O_RDONLY, 0755)
		assert.Nil(t, err, "can not CreateTemp")
		defer fsrc.Close()

		fdst, err := os.CreateTemp("", "sample")
		assert.Nil(t, err, "can not CreateTemp")
		defer os.Remove(fdst.Name())

		err = Copy(fsrc.Name(), fdst.Name(), 10000, 0)

		assert.Equal(t, ErrOffsetExceedsFileSize, err, "expected ErrOffsetExceedsFileSize")
	})
}
