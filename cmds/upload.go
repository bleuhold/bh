package cmds

import (
	"errors"
	"fmt"
	"github.com/bleuhold/bh/ecsv"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"os"
	"strings"
)

var UPLOAD *cli.Command

// UploadExecute is the function executed when the upload command is called.
func UploadExecute(cmd *cli.Command) error {
	// since both -f and -file point to variable s1
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	default:
		xb, err := validateCSV(&S1)
		if err != nil {
			return err
		}
		xt, err := marshalCSV(xb)
		xt, err = appendTransactions(xt)
		xt.Save()
		return err
	}
}

// validateCSV validates that the path points to a CSV file.
//
// Validates:
// 1. The path is a path to a file.
// 2. The file extension is csv.
// 3. Finally, read the file and return the []bytes or the error.
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
	return xb, err
}

func marshalCSV(xb []byte) (*Transactions, error) {
	c := ecsv.CSV{
		StartOffset: 2, // for Investec CSV files
	}
	c.ReadData(xb)
	xt := make(Transactions, 0)
	err := xt.MarshalCSV("investec", uuid.New(), &c)
	if err != nil {
		return &xt, err
	}
	//fmt.Println(xt.String())
	return &xt, nil
}

// appendTransactions takes a Transactions, loads the saved transactions.
// It compares the new transactions to the current transactions and only appends
// non-duplicate transactions to the current transactions.
// the returns the transactions.
func appendTransactions(xt *Transactions) (*Transactions, error) {
	// cxt denotes current slice of transactions
	cxt := LoadTransactions()
	fmt.Println(cxt.String())
	for _, ti := range *xt {
		// by default, we will want to add a new transaction
		add := true
		for _, t := range *cxt {
			// TODO: remove condition and implement bottom, once account are introduced
			if t.Date == ti.Date && t.Description == ti.Description && t.Debit == ti.Debit && t.Credit == ti.Credit {
				// if the transaction is already in the current slice of
				// transaction, then do not add the transaction.
				add = false
			}
			//if t == ti {
			//	// if the transaction is already in the current slice of
			//	// transaction, then do not add the transaction.
			//	add = false
			//}
		}
		if add {
			*cxt = append(*cxt, ti)
		}
	}
	return cxt, nil
}
