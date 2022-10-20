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
		file, err := os.CreateTemp("", "sample")
		assert.Nil(t, err, "Can not CreateTemp")
		defer os.Remove(file.Name())
		file.WriteString("Sergei Kharkov 27.09.2022, 19:41 HW5 is completed!\nSergei Kharkov 27.09.2022, 19:36 HW5 is completed!")

		destpath := "./temp/export.txt"

		err = Copy(file.Name(), destpath, 0, 0)
		assert.Nil(t, err, "Func Copy return error")
		defer os.Remove(destpath)

		srcbytes, err := os.ReadFile(file.Name())
		assert.Nil(t, err, "Can not read source file")

		destbytes, err := os.ReadFile(destpath)
		assert.Nil(t, err, "Can not read destination file")

		require.True(t, Equal(srcbytes, destbytes), "Not equal source file and destination file")
	})

	t.Run("Offset more than size file", func(t *testing.T) {
		file, err := os.CreateTemp("", "sample")
		defer os.Remove(file.Name())
		file.WriteString("Sergei Kharkov 27.09.2022, 19:41 HW5 is completed!\nSergei Kharkov 27.09.2022, 19:36 HW5 is completed!")
		destpath := "./temp/export.txt"
		err = Copy(file.Name(), destpath, 1000, 0)
		if err == nil {
			os.Remove(destpath)
		}
		assert.Equal(t, ErrOffsetExceedsFileSize, err, "Expected ErrOffsetExceedsFileSize")
	})
}
