package ecsv

import (
	"runtime"
	"strings"
)

type CSV struct {
	StartOffset int
	Records     [][]string
}

// ReadData uses the offset to determine from which line/row to start parse the
// data into a CSV, excluding the header data and stores the result on the CSV
// Records field.
func (c *CSV) ReadData(xb []byte) {
	xxs := make([][]string, 0)
	// get the separator based on the runtime operating system
	sep := "\n"
	switch runtime.GOOS {
	case "darwin":
		sep = "\n"
	case "windows":
		sep = "\r\n"
	}

	// first split records
	xs := strings.Split(string(xb), sep)
	// then offset the header
	ds := xs[c.StartOffset:]
	for _, s := range ds {
		// now split the columns
		s = strings.Replace(s, "\n", "", -1)
		s = strings.Replace(s, "\r", "", -1)
		if len(s) > 0 {
			xxs = append(xxs, strings.Split(s, ","))
		}
	}
	//return xxs
	c.Records = xxs
}
