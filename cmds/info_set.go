package cmds

import (
	"fmt"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"time"
)

func infoSetExecute(cmd *cli.Command) error {
	switch {
	case premisesUUID != "":
		SetInfo(&premisesUUID, &startDate, &endDate)
	case startDate != "":
		SetInfo(&premisesUUID, &startDate, &endDate)
	case endDate != "":
		SetInfo(&premisesUUID, &startDate, &endDate)
	default:
		cmd.PrintHelp()
	}
	return nil
}

func SetInfo(premisesUUID *string, startDate *string, endDate *string) {
	i := LoadInfo()
	if *premisesUUID != "" {
		UUID, err := uuid.Parse(*premisesUUID)
		if err != nil {
			fmt.Printf("Unable to set property UUID: %v\n", err)
		} else {
			i.PropertyUUID = UUID
		}
	}
	if *startDate != "" {
		t, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			fmt.Printf("Unable to set start date: %v\n", err)
		} else {
			i.StartDate = t
		}
	}
	if *endDate != "" {
		t, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			fmt.Printf("Unable to set end date: %v\n", err)
		} else {
			i.EndDate = t
		}
	}
	i.Save()
}
