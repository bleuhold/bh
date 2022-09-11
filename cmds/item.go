package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
	"time"
)

var ITEM *cli.Command
var itemsFilename = "items.json"

type Item struct {
	UUID            uuid.UUID `json:"uuid"`
	TransactionUUID uuid.UUID `json:"transactionUUID"`
	Date            time.Time `json:"date"`
	Description     string    `json:"description"`
	Debit           float64   `json:"debit"`
	Credit          float64   `json:"credit"`
	Tags            []string  `json:"tags"`
}

// NewItem creates a new Item and returns the value pointed to.
func NewItem(TransactionUUID uuid.UUID) *Item {
	i := &Item{
		UUID:            uuid.New(),
		TransactionUUID: TransactionUUID,
	}
	return i
}

func (i *Item) AddTags(tags []string) {
	addTags := make([]string, 0)
	for _, tag := range tags {
		add := true
		for _, iTag := range i.Tags {
			if tag == iTag {
				// this tag already exists
				add = false
			}
		}
		if add {
			addTags = append(addTags, tag)
		}
	}
	// add all the tags at once
	i.Tags = append(i.Tags, addTags...)
}

func (i *Item) RemoveTags(tags []string) {
	keepTags := make([]string, 0)
	for _, iTag := range i.Tags {
		keep := true
		for _, tag := range tags {
			if tag == iTag {
				keep = false
			}
		}
		if keep {
			keepTags = append(keepTags, iTag)
		}
	}
	i.Tags = keepTags
}

// TODO ITEMS: update to a map

type Items []Item

// LoadItems loads all the items from the file system.
func LoadItems() *Items {
	xi := &Items{}
	filesys.LoadInterface(itemsFilename, xi)
	return xi
}

// Save writes the items to the file system.
func (xi *Items) Save() {
	xb, err := json.Marshal(*xi)
	if err != nil {
		log.Fatalf("Unable to marshal items data: %v", err)
	}
	filesys.WriteFile(itemsFilename, xb)
}

// String returns a string representation of the items.
func (xi *Items) String() string {
	s := fmt.Sprintf("%-36s %-36s %10s %-40s %-11s %-11s %s", "UUID", "TRANSACTION UUID", "DATE", "DESCRIPTION", "DEBIT", "CREDIT", "TAGS")
	for _, i := range *xi {
		desc := i.Description
		if len(i.Description) > 40 {
			desc = i.Description[:40]
		}
		s += fmt.Sprintf("%36s %36s %10s %-40s %11.2f %11.2f %s\n", i.UUID, i.TransactionUUID, i.Date.Format("2006-01-02"), desc, i.Debit, i.Credit, i.Tags)
	}
	return s
}

// Add appends items to the Items value pointed to.
func (xi *Items) Add(items ...Item) {
	*xi = append(*xi, items...)
}

/*
	COMMAND
*/

func ItemExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case B1:
		ListItems()
	}
	return nil
}

// ListItems is the executable function to list items.
func ListItems() {
	xi := LoadItems()
	fmt.Println(xi.String())
}
