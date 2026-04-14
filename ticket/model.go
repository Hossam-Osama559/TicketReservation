package ticket

import (
	"errors"
)

type Ticket struct {
	Id int

	Origin string

	Arrival string

	Price int

	UserId int

	Scanned bool
}

func NewTicket() *Ticket {

	return &Ticket{Scanned: false}
}

func (t *Ticket) Validate() error {

	if t.Origin == "" {

		return errors.New("origin is required field")
	} else if t.Arrival == "" {

		return errors.New("arrival is required field")
	}

	return nil
}
