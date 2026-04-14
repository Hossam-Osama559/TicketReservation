package payment

import "net/http"

type Session interface {

	// geting the url of the checkout page of this session
	GetUrl() string

	// getting the session id
	GetId() string
}

type PaymentProcessor interface {

	// make a new payment and return the
	NewSession(string, string, int64) (Session, error)

	// handling the webhook of the processor
	Webhook(http.ResponseWriter, *http.Request)
}
