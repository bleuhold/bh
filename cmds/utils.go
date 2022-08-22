package cmds

import (
	"log"
	"time"
)

// MustParseTime is used for testing, to parse a date without error.
func MustParseTime(layout string, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		log.Fatalln(err)
	}
	return t
}
