package cmds

import (
	"errors"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"time"
)

var INFO_SET *cli.Command

func InfoSetExecute(cmd *cli.Command) error {
	u, s, e, err := parseSetInfoData(&S1, &S2, &S3)

	// if there are boolean flags
	// check in the switch statement
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case err == nil:
		return SetInfo(&u, &s, &e)
	default:
		cmd.PrintHelp()
	}
	return err
}

// parseSetInfoData parses and validates the data passed as args when invoking
// the `info set` command.
func parseSetInfoData(UUID *string, startDate *string, endDate *string) (uuid.UUID, time.Time, time.Time, error) {
	var u uuid.UUID
	var s time.Time
	var e time.Time
	var ok bool

	if *UUID != "" {
		ok = true
		uParsed, err := uuid.Parse(*UUID)
		u = uParsed
		if err != nil {
			return u, s, e, err
		}
	}

	if *startDate != "" {
		ok = true
		t, err := time.Parse("2006-01-02", *startDate)
		s = t
		if err != nil {
			return u, s, e, err
		}
	}

	if *endDate != "" {
		ok = true
		t, err := time.Parse("2006-01-02", *endDate)
		e = t
		if err != nil {
			return u, s, e, err
		}
	}

	if !ok {
		return u, s, e, errors.New("invalid args: at least one argument is required")
	}

	return u, s, e, nil
}

// SetInfo checks if there are any changes
func SetInfo(UUID *uuid.UUID, startDate *time.Time, endDate *time.Time) error {
	i := LoadInfo()
	b := false
	if i.PropertyUUID != *UUID {
		i.PropertyUUID = *UUID
		b = true
	}
	if i.StartDate != *startDate {
		i.StartDate = *startDate
		b = true
	}
	if i.EndDate != *endDate {
		i.EndDate = *endDate
		b = true
	}
	if b {
		i.Save()
	}
	return nil
}
