package payment

import (
	"TicketRservation/env"
	"TicketRservation/qrcode"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
	"github.com/stripe/stripe-go/v74/webhook"
)

type StripeSession stripe.CheckoutSession

func (session *StripeSession) GetUrl() string {

	return session.URL
}

func (session *StripeSession) GetId() string {

	return session.ID
}

type StripeProcessor struct {
	key string

	// must depend on a payment service to handle the change of the payment when recieving a webhook
	paymentservice *PaymentService

	// depend on a qrcode service to handle the creation of the qrcodes
	qrcodeservice *qrcode.QrcodeService
}

func NewStripeProcessor(paymentservice *PaymentService, qrcodeservice *qrcode.QrcodeService) *StripeProcessor {

	// getting the key from the env vars

	key := env.StripeKey.GetValue()
	return &StripeProcessor{key: key, paymentservice: paymentservice, qrcodeservice: qrcodeservice}
}

func (processor *StripeProcessor) NewSession(startstation, endstation string, price int64) (Session, error) {

	stripe.Key = processor.key

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(startstation + " → " + endstation + " Ticket"),
					},
					UnitAmount: stripe.Int64(price), // in cents
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(env.StripeSuccessUrl.GetValue()),
		CancelURL:  stripe.String(env.StripeCancelUrl.GetValue()),
	}

	s, err := session.New(params)

	stripesession := StripeSession(*s)

	return &stripesession, err

}

func (processor *StripeProcessor) Webhook(w http.ResponseWriter, r *http.Request) {

	payload, err := io.ReadAll(r.Body)

	if err != nil {

	} else {

		stripe_sign := r.Header.Get("Stripe-Signature")

		webhook_secret := env.WebhookSecret.GetValue()

		eve, _ := webhook.ConstructEvent(payload, stripe_sign, webhook_secret)

		switch eve.Type {
		case "checkout.session.completed":

			// in this case we must make the qrcode and modify the state of the payment of this session

			var session stripe.CheckoutSession

			err = json.Unmarshal(eve.Data.Raw, &session)

			// mark the payment as completed

			processor.paymentservice.MarkCompleter(session.ID)

			paymetrecord, _ := processor.paymentservice.Get(session.ID)

			ticketid := paymetrecord.TicketId

			// make the qrcode
			processor.qrcodeservice.NewQrcode(strconv.Itoa(ticketid), strconv.Itoa(ticketid))

		// to do : handle all the cases of the event

		default:

			// not handled yet

		}

	}

}
