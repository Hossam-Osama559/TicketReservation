package ticket

import (
	"TicketRservation/payment"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aarondl/authboss/v3"
)

type TicketHandler struct {
	TicketService *TicketService

	// payment service ----->make the payment record
	paymentservice *payment.PaymentService

	// payment processor --->make the stripe session
	paymentprocessor payment.PaymentProcessor
}

func NewTicketHandler(ticketservice *TicketService, paymentservice *payment.PaymentService, paymentprocessor payment.PaymentProcessor) *TicketHandler {

	return &TicketHandler{TicketService: ticketservice, paymentservice: paymentservice, paymentprocessor: paymentprocessor}
}

// post tickets {Origin:"",Arrival:""}
// response {Url:""}
func (handler *TicketHandler) CreateTicketHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("post ticket")

	// make the ticket record
	// make the payment record
	// make the session and return the url of the checkout

	ticket := NewTicket()

	// dummy price for now
	ticket.Price = 3000

	// loading the clinet state

	clinetsession := r.Context().Value(authboss.CTXKeySessionState)

	if clinetsession == nil {

		// the user not authenticated
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	usersession, _ := clinetsession.(authboss.ClientState)

	userid, _ := usersession.Get("userid")

	ticket.UserId, _ = strconv.Atoi(userid)

	err := json.NewDecoder(r.Body).Decode(ticket)

	fmt.Println("got new ticket ", ticket)

	if err != nil {

		fmt.Println("error in des in the ticket ", err)

		// to do : handle the error
	}

	err = handler.TicketService.Create(ticket)

	if err != nil {

		fmt.Println("error in creating the ticket ", err)

		// to do : handle the error
	}

	// make the session to get the url and the id
	session, _ := handler.paymentprocessor.NewSession(ticket.Origin, ticket.Arrival, int64(ticket.Price))

	// make the payment with the sessionid
	payment := payment.NewPayment()
	payment.Completed = false
	payment.SessionId = session.GetId()
	payment.TicketId = ticket.Id
	handler.paymentservice.Create(*payment)

	json.NewEncoder(w).Encode(struct{ Url string }{Url: session.GetUrl()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// PUT /tickets {"Ticket":ticketid}
// res {"State":valid or not valid}
func (handler *TicketHandler) UpdateTicketHandler(w http.ResponseWriter, r *http.Request) {

	// checks if the session belong to employee or not

	responsebody := struct{ State string }{}

	ticketid := struct{ Ticket string }{Ticket: "0"}

	json.NewDecoder(r.Body).Decode(&ticketid)

	intticketid, _ := strconv.Atoi(ticketid.Ticket)
	fmt.Println("scanning now the ticket ", ticketid, "  ", intticketid)

	ticket, err := handler.TicketService.repo.GetById(intticketid)

	if err != nil {

		responsebody.State = err.Error()

		w.WriteHeader(http.StatusNotFound)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(responsebody)

		return

	}

	if ticket.Scanned == true {

		responsebody.State = "invalid"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(responsebody)

		return

	} else {

		responsebody.State = "valid"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(responsebody)

		// update the ticket
		ticket.Scanned = true
		handler.TicketService.repo.Update(ticket)
		return
	}

}

// GET /tickets
// {"tickets":["",""]}
func (handler *TicketHandler) GetTicketHandler(w http.ResponseWriter, r *http.Request) {

	userstate := r.Context().Value(authboss.CTXKeySessionState)

	if userstate == nil {

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	usersession := userstate.(authboss.ClientState)

	userid, _ := usersession.Get("userid")

	intuserid, _ := strconv.Atoi(userid)

	tickets, err := handler.TicketService.GetByUser(intuserid)

	if err != nil {

		// handler the error
	}

	json.NewEncoder(w).Encode(struct{ Tickets []Ticket }{Tickets: tickets})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (handler *TicketHandler) DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {

}

func (handler *TicketHandler) RegisterHandlers(router *http.ServeMux) {

	router.HandleFunc("GET /tickets/", handler.GetTicketHandler)

	router.HandleFunc("POST /tickets/", handler.CreateTicketHandler)

	router.HandleFunc("PUT /tickets/", handler.UpdateTicketHandler)

	router.HandleFunc("DELETE /tickets/", handler.DeleteTicketHandler)

}
