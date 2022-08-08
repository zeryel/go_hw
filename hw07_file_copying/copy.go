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
	if nil != err {
		return err
	}

	if !fromFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fromFileInfo.Size() {
		return ErrOffsetExceedsFileSize
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

	if 0 == limit || limit > fromFileInfo.Size() {
		limit = fromFileInfo.Size()
	}

	_, err = fromFile.Seek(offset, 0)
	if err != nil {
		return err
	}

	progressBar := pb.New(int(limit))
	_, err = io.CopyN(toFile, progressBar.NewProxyReader(fromFile), limit)
	if err != nil {
		return err
	}

	progressBar.Finish()

	return nil
}
