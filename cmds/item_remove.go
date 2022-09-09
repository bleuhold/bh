package cmds

import (
	"fmt"
	"github.com/dottics/cli"
	"github.com/google/uuid"
)

var ITEM_REMOVE *cli.Command

func ItemRemoveExecute(cmd *cli.Command) error {
	UUID, err := uuid.Parse(S1)
	switch {
	case Help:
		cmd.PrintHelp()
	case err == nil:
		return removeItem(UUID)
	}
	return nil
}

func removeItem(UUID uuid.UUID) error {
	xi := LoadItems()
	items := *xi
	found := false
	for i, item := range items {
		if item.UUID == UUID {
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("item not found for uuid: %v", UUID)
	}
	xi = &items
	xi.Save()
	return nil
}
