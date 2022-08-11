package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessCopy(t *testing.T) {
	tests := []struct {
		fileTo string
		offset int64
		limit  int64
	}{
		{
			fileTo: "out_offset0_limit0.txt",
			offset: 0,
			limit:  0,
		},
		{
			fileTo: "out_offset0_limit10.txt",
			offset: 0,
			limit:  10,
		},
		{
			fileTo: "out_offset0_limit1000.txt",
			offset: 0,
			limit:  1000,
		},
		{
			fileTo: "out_offset0_limit10000.txt",
			offset: 0,
			limit:  10000,
		},
		{
			fileTo: "out_offset100_limit1000.txt",
			offset: 100,
			limit:  1000,
		},
		{
			fileTo: "out_offset6000_limit1000.txt",
			offset: 6000,
			limit:  1000,
		},
	}

	fileFrom := "input.txt"
	copyFile := "copy.txt"

	err := os.Chdir("./testdata")
	if err != nil {
		require.Fail(t, "не удалось сменить рабочую папку для тестов")
	}

	for _, tc := range tests {
		t.Run(tc.fileTo, func(t *testing.T) {
			defer os.Remove(copyFile)

			err := Copy(fileFrom, copyFile, tc.offset, tc.limit)
			require.Nil(t, err)
		})
	}
}

func TestErrors(t *testing.T) {
	t.Run("пробуем скопировать папку", func(t *testing.T) {
		err := Copy("./testdata", "./tmp", 0, 1000)

		require.Error(t, err)
	})

	t.Run("слишком большой offset", func(t *testing.T) {
		err := Copy("input.txt", "copy.txt", 15000, 200)

		require.Error(t, err)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})
}
