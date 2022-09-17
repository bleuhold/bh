package cmds

import (
	"fmt"
	"github.com/bleuhold/bh/filesys"
	"github.com/dottics/cli"
	"github.com/google/uuid"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

var STATEMENT *cli.Command
var tpl *template.Template

func formatDateYMD(t time.Time) string {
	return t.Format("2006-01-02")
}

var funcMap = template.FuncMap{
	"formatDate": formatDateYMD,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
}

type Statement struct {
	filename  string
	Reference string
	Premises  Premises
	Date      struct {
		Start time.Time
		End   time.Time
	}
	Account  Account
	Landlord string
	Agent    struct {
		FirstName     string
		LastName      string
		ContactNumber string
		Email         string
	}
	Tenants Tenants
	Items   Items
}

// NewStatement creates a new statement based on a specific contract's details.
func NewStatement(c *Contract) *Statement {
	endDate := time.Now()
	return &Statement{
		filename:  fmt.Sprintf("%s-%s-statement", endDate.Format("20060102"), c.Reference),
		Reference: c.Reference,
		Premises:  c.Premises,
		Date: struct {
			Start time.Time
			End   time.Time
		}{
			Start: c.Dates.Commencement,
			End:   endDate,
		},
		Account:  c.Account,
		Landlord: "J and M Scribante",
		Agent: struct {
			FirstName     string
			LastName      string
			ContactNumber string
			Email         string
		}{
			FirstName:     c.Agent.FirstName,
			LastName:      c.Agent.LastName,
			ContactNumber: c.Agent.ContactNumber,
			Email:         c.Agent.Email,
		},
		Tenants: c.Tenants,
		Items:   nil,
	}
}

func (s *Statement) LoadTransactions() {}

//func (s *Statement) filename(extension string) string {
//	endDate := s.Date.End.Format("20060102")
//	return fmt.Sprintf("%s-%s-statement.%s", endDate, s.Reference, extension)
//}

// Write executes the HTML template to populate all data.
func (s *Statement) Write() {
	file, err := filesys.CreateFile(s.filename + ".html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(file, "statement.html", s)
	if err != nil {
		log.Fatalln(err)
	}
}

// PDF converts and generates the PDF from the HTML statement template.
func (s *Statement) PDF() error {
	home := os.Getenv("HOME")
	pdfPath := path.Join(home, "Downloads", s.filename+".pdf")
	htmlPath := filesys.FilePath(s.filename + ".html")
	cmd := exec.Command("wkhtmltopdf", htmlPath, pdfPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	unicodeCircle := fmt.Sprintf("\u25cf")
	fmt.Printf("\u001b[32m%s\u001B[0m PDF generated successfully\n\n", unicodeCircle)
	return nil
}

/*
	STATEMENT
*/

func StatementExecute(cmd *cli.Command) error {
	switch {
	case Help:
		cmd.PrintHelp()
		return nil
	case S1 != "" && B1 == false:
		UUID, err := uuid.Parse(S1)
		if err != nil {
			return err
		}
		c := GetContract(UUID)
		if c == nil {
			return fmt.Errorf("contract not found with UUID: %v", UUID)
		}
		s := NewStatement(c)
		xi := LoadItems()
		items := xi.DateRange(c.Dates.Occupation, c.Dates.Termination)
		items = items.FilterTags(c.References())
		s.Items = *items

		fmt.Printf("\n\u001B[1mSTATEMENT\u001B[0m\n\n")
		c.Print(false)
		fmt.Println(items.String())

	case S1 != "" && B1:
		UUID, err := uuid.Parse(S1)
		if err != nil {
			return err
		}
		c := GetContract(UUID)
		if c == nil {
			return fmt.Errorf("contract not found with UUID: %v", UUID)
		}
		s := NewStatement(c)
		xi := LoadItems()
		items := xi.DateRange(c.Dates.Occupation, c.Dates.Termination)
		items = items.FilterTags(c.References())
		s.Items = *items
		s.Write()
		err = s.PDF()
		if err != nil {
			return err
		}
		fmt.Printf("")
		return nil
	}
	return nil
}
