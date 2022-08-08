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
		{fileTo: "out_offset0_limit0.txt", offset: 0, limit: 0},
	}

	for _, tc := range tests {
		os.Chdir("./testdata")

		t.Run(tc.fileTo, func(t *testing.T) {
			err := Copy("input.txt", tc.fileTo, tc.offset, tc.limit)
			require.Nil(t, err)
		})
	}
}
