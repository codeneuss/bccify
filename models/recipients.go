package models

import "fmt"

type Recipents []string

func (r *Recipents) Add(c Recipient) {
	*r = append(*r, c.Recipient())
}

func (c *Recipient) Email() string {
	return c.email
}

func (c *Recipient) Recipient() string {
	return fmt.Sprintf("\"%s\" <%s>", c.name, c.email)
}

func NewRecipient(email, name string) Recipient {
	return Recipient{
		email: email,
		name:  name,
	}
}
