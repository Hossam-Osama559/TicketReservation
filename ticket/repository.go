package ticket

type TicketRepository interface {
	Create(t *Ticket) error

	GetById(id int) (*Ticket, error)

	Update(t *Ticket) error

	Delete(id int) error
	GetByUser(userId int) ([]Ticket, error)
}
