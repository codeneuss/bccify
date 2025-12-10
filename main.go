package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codeneuss/bccify/importer"
)

type Recipient struct {
	name  string
	email string
}

type Recipents []string

func (r *Recipents) Add(c Recipient) {
	*r = append(*r, c.Recipient())
}

func (c *Recipient) Email() string {
	return c.email
}

func (c *Recipient) Recipient() string {
	return fmt.Sprintf("%s <%s>", c.name, c.email)
}

func NewContact(email, name string) Recipient {
	return Recipient{
		email: email,
		name:  name,
	}
}

func main() {
	fmt.Println("Welcome to BCCify!")
	args := os.Args

	if len(args) < 1 {
		fmt.Println("You need to provide a file as an argument.")
		os.Exit(1)
	}

	fmt.Println(args[1])

	var imp importer.Importer
	switch {
	case strings.HasSuffix(args[1], ".csv"):
		imp = &importer.CSVImporter{
			Filename:   args[1],
			HasHeaders: true,
		}
	}

	if imp == nil {
		fmt.Println("No Importer found")
		os.Exit(1)
	}

	if err := imp.Import(); err != nil {
		fmt.Println("Error on import: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(imp.FilterColumns())

	for _, record := range imp.Filter(func(r importer.Record) bool { return strings.HasPrefix(r["name"], "J") }) {
		fmt.Println(record["emailaddress"])
	}

}
