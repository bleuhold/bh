package cmds

import (
	"fmt"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"strings"
)

var ITEM_TAG *cli.Command

func ItemTagExecute(cmd *cli.Command) error {
	UUID, tags, err := validateItemTag(&S1, &S2)
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case err != nil:
		return err
	case B1: // add
		err = addItemTags(UUID, tags)
		break
	case B2: // remove
		err = removeItemTags(UUID, tags)
		break
	}
	return err
}

// validateItemTag validates the data passed to add/remove tags from an item.
func validateItemTag(s1, s2 *string) (uuid.UUID, []string, error) {
	UUID, err := uuid.Parse(*s1)
	if err != nil {
		return uuid.UUID{}, []string{}, err
	}
	tags := strings.Split(*s2, ",")
	return UUID, tags, nil
}

func addItemTags(UUID uuid.UUID, tags []string) error {
	xi := LoadItems()
	found := -1
	//x := *xi
	for idx, i := range *xi {
		if i.UUID == UUID {
			found = idx
		}
	}
	if found < 0 {
		return fmt.Errorf("item not found for UUID: %v", UUID)
	}
	(*xi)[found].AddTags(tags)
	xi.Save()
	return nil
}

func removeItemTags(UUID uuid.UUID, tags []string) error {
	xi := LoadItems()
	found := -1
	for idx, i := range *xi {
		if i.UUID == UUID {
			found = idx
		}
	}
	if found < 0 {
		return fmt.Errorf("item not found for UUID: %v", UUID)
	}
	(*xi)[found].RemoveTags(tags)
	xi.Save()
	return nil
}
