package cmds

import (
	"errors"
	"fmt"
	"github.com/dottics/cli"
	"os"
	"strings"
)

// uploadExecute is the function executed when the upload command is called.
func uploadExecute(cmd *cli.Command) error {
	// since both -f and -file point to variable s1
	_, err := validateCSV(&s1)
	switch {
	case help:
		cmd.PrintHelp()
		return nil
	case err != nil:

	}
	return err
}

// validateCSV validates that the path points to a CSV file.
func validateCSV(path *string) ([]byte, error) {
	fileInfo, err := os.Stat(*path)
	if err != nil {
		return []byte{}, err
	}
	if fileInfo.IsDir() {
		return []byte{}, errors.New("invalid path: points to a directory not a file")
	}
	s := strings.Split(fileInfo.Name(), ".")
	// get the file extension
	ext := s[len(s)-1]
	ext = strings.ToLower(ext)
	if ext != "csv" {
		return []byte{}, fmt.Errorf("invalid file extension: expected '%s' got '%s'", "csv", ext)
	}
	xb, err := os.ReadFile(*path)
	if err != nil {
		return []byte{}, err
	}
	return xb, nil
}
