package main

import (
	"TicketRservation/auth"
	"TicketRservation/db"
	"TicketRservation/env"
	"TicketRservation/payment"
	"TicketRservation/qrcode"
	"TicketRservation/renderer"
	"TicketRservation/session"
	"TicketRservation/ticket"
	"TicketRservation/user"
	"fmt"
	"net/http"
	"text/template"

	"github.com/aarondl/authboss/v3"
)

var (
	tickethandler          *ticket.TicketHandler
	mux                    = http.NewServeMux()
	mysql_db               = &db.DbInstance
	stripepaymentprocessor payment.PaymentProcessor
	ticketservice          *ticket.TicketService
	paymentservice         *payment.PaymentService
	qrcodeservice          *qrcode.QrcodeService
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	clinetstate := r.Context().Value(authboss.CTXKeySessionState)

	if clinetstate == nil {

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	http.ServeFile(w, r, "./static/reserve/reserve.html")

}

func servestatic(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	path = "." + path
	fmt.Println(path)
	http.ServeFile(w, r, path)
}

func successpayment(w http.ResponseWriter, r *http.Request) {

	homepageurl := struct{ Link string }{Link: env.HomePageUrl.GetValue()}

	fmt.Println("the home page is ", homepageurl.Link)

	tmpl, err := template.ParseFiles("./static/success/success.html")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, homepageurl)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func scaner(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./static/scan/scan.html")
}

func main() {

	// loading the env vars
	env.LoadEnv()

	// init the db
	db.InitDb()

	// migrate the db
	db.MigrateUp()

	// init the ticker service
	InitTicketservice()

	// init the payment service
	InitPaymentservice()

	// init the qrcode service
	InitQrcodeservice()

	// init stripe and register the webhook
	InitStripe()

	// init the ticket handler and register the handlers
	newtickethandler()

	SetupAuthboss()

	mux.Handle("/register", auth.AuthbossInstance.Core.Router)

	mux.Handle("/login", auth.AuthbossInstance.Core.Router)

	mux.HandleFunc("/", MainPage)

	mux.HandleFunc("/scan", scaner)

	mux.HandleFunc("/static/", servestatic)

	mux.HandleFunc("/success/", successpayment)

	final_handler := auth.SetAuthMiddleWare(mux)

	fmt.Println("the server just started")
	port := env.ServerPort.GetValue()
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), final_handler)

	if err != nil {

		fmt.Println(err)
	}
}

func InitTicketservice() {
	mysqlticketrepo := ticket.NewMysqlTicketRepo(*mysql_db)

	ticketservice = ticket.NewTicketService(mysqlticketrepo)

}

func InitPaymentservice() {
	mysqlpaymentrepo := payment.NewMysqlPaymentRepo(*mysql_db)
	paymentservice = payment.NewPaymentService(mysqlpaymentrepo)
}

func InitQrcodeservice() {

	qrcodeservice = qrcode.NewQrcodeService()
}

func InitStripe() {

	stripepaymentprocessor = payment.NewStripeProcessor(paymentservice, qrcodeservice)

	mux.HandleFunc("/webhook", stripepaymentprocessor.Webhook)
}

func newtickethandler() {

	tickethandler = ticket.NewTicketHandler(ticketservice, paymentservice, stripepaymentprocessor)

	tickethandler.RegisterHandlers(mux)
}

func SetupAuthboss() {

	mysqluserRepo := user.NewMysqlUserRepo(*mysql_db)

	mysqlsessionrepo := session.NewMysqlSessionRepo(*mysql_db)

	renderer := renderer.NewFileRenderer()

	auth.Setup(mysqluserRepo, mysqlsessionrepo, renderer)
}
