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
		fsrc, err := os.CreateTemp("", "sample")
		assert.Nil(t, err, "Can not CreateTemp")
		defer os.Remove(fsrc.Name())
		fsrc.WriteString("Sergei Kharkov 27.09.2022, 19:41 HW5 is completed!\nSergei Kharkov 27.09.2022, 19:36 HW5 is completed!")

		fdst, err := os.CreateTemp("", "sample")
		assert.Nil(t, err, "Can not CreateTemp")
		defer os.Remove(fdst.Name())

		err = Copy(fsrc.Name(), fdst.Name(), 0, 0)
		assert.Nil(t, err, "Func Copy return error")

		srcbytes, err := os.ReadFile(fsrc.Name())
		assert.Nil(t, err, "Can not read source file")

		destbytes, err := os.ReadFile(fdst.Name())
		assert.Nil(t, err, "Can not read destination file")

		require.True(t, Equal(srcbytes, destbytes), "Not equal source file and destination file")
	})

	t.Run("Offset more than size file", func(t *testing.T) {
		fsrc, err := os.CreateTemp("", "sample")
		defer os.Remove(fsrc.Name())
		fsrc.WriteString("Sergei Kharkov 27.09.2022, 19:41 HW5 is completed!\nSergei Kharkov 27.09.2022, 19:36 HW5 is completed!")

		fdst, err := os.CreateTemp("", "sample")
		assert.Nil(t, err, "Can not CreateTemp")
		defer os.Remove(fdst.Name())

		err = Copy(fsrc.Name(), fdst.Name(), 1000, 0)

		assert.Equal(t, ErrOffsetExceedsFileSize, err, "Expected ErrOffsetExceedsFileSize")
	})
}
