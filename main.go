package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	name  string
	email string
}

type Recipents []string

func (r *Recipents) Add(c Contact) {
	*r = append(*r, c.Recipient())
}

func (c *Contact) Email() string {
	return c.email
}

func (c *Contact) Recipient() string {
	return fmt.Sprintf("%s <%s>", c.name, c.email)
}

func NewContact(email, name string) Contact {
	return Contact{
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
	fd, err := os.Open(args[1])
	if err != nil {
		fmt.Printf("Error opening file %s\n", args[1])
		os.Exit(2)
	}

	csv := csv.NewReader(fd)

	var recipients Recipents
	for {
		record, err := csv.Read()
		if err != nil {
			break
		}
		if strings.Contains(record[1], ";") {
			emailaddesses := strings.Split(record[1], ";")
			for i := range emailaddesses {
				recipients.Add(NewContact(emailaddesses[i], record[0]))
			}
		} else {
			recipients.Add(NewContact(record[1], record[0]))
		}
	}

	for i := range recipients {
		fmt.Println(recipients[i])
	}

}
