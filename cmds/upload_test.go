package cmds

import (
	"errors"
	"github.com/dottics/cli"
	"testing"
)

func TestValidateCSV(t *testing.T) {
	tt := []struct {
		name  string
		path  string
		err   error
		xbLen int
	}{
		{
			name:  "invalid path",
			path:  "./invalid-file.csv",
			err:   errors.New("stat ./invalid-file.csv: no such file or directory"),
			xbLen: 0,
		},
		{
			name:  "path to directory",
			path:  "./../testdata",
			err:   errors.New("invalid path: points to a directory not a file"),
			xbLen: 0,
		},
		{
			name:  "invalid extension",
			path:  "./../testdata/test-bank-statement.ssv",
			err:   errors.New("invalid file extension: expected 'csv' got 'ssv'"),
			xbLen: 0,
		},
		{
			name:  "valid path to return slice of bytes",
			path:  "./../testdata/test-bank-statement.csv",
			err:   nil,
			xbLen: 3824,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			xb, err := validateCSV(&tc.path)
			if cli.ErrorNotEqual(tc.err, err) {
				t.Errorf("expected error '%v' got '%v'", tc.err, err)
			}
			if len(xb) != tc.xbLen {
				t.Errorf("expected bytes to have lenth %d got %d", tc.xbLen, len(xb))
			}
		})
	}
}
