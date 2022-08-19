package cmds

/*
ALL FLAGS ARE DEFINED HERE

Flags are shared with all subcommands.
*/

/* ALL GENERIC FLAGS */

// to show help information
var help bool

// to list all available data
var list bool

// to add a new data structure
var add bool

// to remove a data structure
var remove bool

// to update a data structure
var update bool

// to set a piece of information/config/settings
var set bool

/* MORE SPECIFIC FLAGS */

// the name of a data structure
var name string

var address string

// the plotNumber of a data structure
var plotNumber string

// the premisesUUID for filtering premises
var premisesUUID string

// the startDate for filtering
var startDate string

// the endDate for filtering
var endDate string
