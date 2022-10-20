package main

import (
	"errors"
	"io"
	"math"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrCanNotOpenFile        = errors.New("Can not open file")
	ErrNoDataRead            = errors.New("No data for read")
	ErrLimitNotCorrect       = errors.New("Limit is not correct")
	ErrCannotCreateFile      = errors.New("Can not create file")
	ErrCopyFile              = errors.New("Err copy")
	ErrSeek                  = errors.New("Err Seek")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fsrc, err := os.OpenFile(fromPath, os.O_RDONLY, 0755)
	if err != nil {
		return ErrCanNotOpenFile
	}
	defer fsrc.Close()

	fsrci, err := fsrc.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	size := fsrci.Size()

	if offset < 0 {
		offset += size
	}

	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	if limit < 0 {
		return ErrLimitNotCorrect
	}

	if limit == 0 {
		limit = size - offset
	}

	if limit+offset > size {
		limit = size - offset
	}

	if limit == 0 {
		return ErrNoDataRead
	}

	_, err = fsrc.Seek(offset, 1)
	if err != nil {
		return ErrSeek
	}
	step := int64(math.Ceil(float64(limit) / 100))

	fdst, err := os.Create(toPath)
	if err != nil {
		return ErrCannotCreateFile
	}
	defer fdst.Close()

	// create bar
	bar := pb.New64(limit).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.Start()

	// create proxy writer
	barWriter := bar.NewProxyWriter(fdst)

	var i int64

	for i = 0; i < limit; i += step {
		if i+step > limit {
			step = limit - i
		}
		_, err := io.CopyN(barWriter, fsrc, step)
		if err != nil {
			return ErrCopyFile
		}
	}

	bar.Finish()

	return nil
}
