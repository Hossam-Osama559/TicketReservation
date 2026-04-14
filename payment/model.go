package payment

type Payment struct {

	// the session id of the stripe session for the payment
	SessionId string

	TicketId int

	Completed bool
}

func NewPayment() *Payment {

	return &Payment{}
}
