package ticket

type TicketService struct {
	repo TicketRepository
}

func NewTicketService(repo TicketRepository) *TicketService {

	return &TicketService{repo: repo}
}

func (service *TicketService) Create(t *Ticket) error {

	// validate the ticket
	err := t.Validate()

	if err != nil {

		return err
	}

	service.repo.Create(t)

	return nil

}

func (service *TicketService) GetByUser(userid int) ([]Ticket, error) {

	return service.repo.GetByUser(userid)
}
