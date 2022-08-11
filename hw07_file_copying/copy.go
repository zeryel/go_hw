package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFileInfo, err := os.Stat(fromPath)
	switch {
	case err != nil:
		return err
	case !fromFileInfo.Mode().IsRegular():
		return ErrUnsupportedFile
	case offset > fromFileInfo.Size():
		return ErrOffsetExceedsFileSize
	case limit == 0 || limit > fromFileInfo.Size():
		limit = fromFileInfo.Size()
	}

	fromFile, err := os.Open(fromPath)
	if nil != err {
		return err
	}

	defer fromFile.Close()

	toFile, err := os.Create(toPath)
	if nil != err {
		return err
	}

	defer toFile.Close()

	_, err = fromFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	if limit+offset > fromFileInfo.Size() {
		limit = fromFileInfo.Size() - offset
	}
	progressBar := pb.New(int(limit))
	_, err = io.CopyN(toFile, progressBar.NewProxyReader(fromFile), limit)
	if err != nil {
		return err
	}

	progressBar.Finish()

	return nil
}
