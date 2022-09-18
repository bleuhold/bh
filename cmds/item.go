package cmds

import (
	"encoding/json"
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"log"
	"sort"
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
	Balance         float64   `json:"-"`
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
	sort.Sort(xi)
	xb, err := json.Marshal(*xi)
	if err != nil {
		log.Fatalf("Unable to marshal items data: %v", err)
	}
	filesys.WriteFile(itemsFilename, xb)
}

// Len returns the length of the slice of items.
func (xi Items) Len() int {
	return len(xi)
}

// Less return whether item i is before item j.
func (xi Items) Less(i, j int) bool {
	return xi[i].Date.Before(xi[j].Date)
}

// Swap interchanges the positions of i and j.
func (xi Items) Swap(i, j int) {
	xi[i], xi[j] = xi[j], xi[i]
}

func (xi *Items) DateRange(start, end time.Time) *Items {
	items := Items{}
	startDate := start.AddDate(0, 0, -1)
	endDate := end.AddDate(0, 0, 1)
	for _, item := range *xi {
		if item.Date.After(startDate) && item.Date.Before(endDate) {
			items = append(items, item)
		}
	}
	return &items
}

func (xi *Items) FilterTags(tags map[string]bool) *Items {
	items := Items{}
	for _, item := range *xi {
		for _, tag := range item.Tags {
			if _, ok := tags[tag]; ok {
				items = append(items, item)
				break
			}
		}
	}
	return &items
}

//func FilterTransactionUUID(UUID uuid.UUID) *Items {
//
//}

func (xi *Items) RemoveTransactionUUID(UUID uuid.UUID) *Items {
	items := Items{}
	for _, item := range *xi {
		if item.UUID != UUID {
			items = append(items, item)
		}
	}
	return &items
}

// String returns a string representation of the items.
func (xi *Items) String() string {
	s := fmt.Sprintf("%-36s %-36s %-10s %-40s %11s %11s %s\n", "UUID", "TRANSACTION UUID", "DATE", "DESCRIPTION", "DEBIT", "CREDIT", "TAGS")
	for _, i := range *xi {
		desc := i.Description
		if len(i.Description) > 40 {
			desc = i.Description[:40]
		}
		s += fmt.Sprintf("%36s %36s %10s %-40s %11.2f %11.2f %s\n", i.UUID, i.TransactionUUID, i.Date.Format("2006-01-02"), desc, i.Debit, i.Credit, i.Tags)
	}
	return s
}

// StatementString returns a string representation of the items.
func (xi *Items) StatementString() string {
	s := fmt.Sprintf("%-36s %-36s %-10s %-40s %11s %11s %11s %s\n", "UUID", "TRANSACTION UUID", "DATE", "DESCRIPTION", "DEBIT", "CREDIT", "BALANCE", "TAGS")
	for _, i := range *xi {
		desc := i.Description
		if len(i.Description) > 40 {
			desc = i.Description[:40]
		}
		s += fmt.Sprintf("%36s %36s %10s %-40s %11.2f %11.2f %11.2f %s\n", i.UUID, i.TransactionUUID, i.Date.Format("2006-01-02"), desc, i.Debit, i.Credit, i.Balance, i.Tags)
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
