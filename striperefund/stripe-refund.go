package striperefund

import (
	"github.com/chinmay2197/stripe-integration-api/logging"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

type StripeRefund struct {
	CLIENT         *client.API
	RefundChargeID string
}

func (sr StripeRefund) CreateStripeRefund() (*stripe.Refund, error) {
	params := &stripe.RefundParams{
		Charge: stripe.String(sr.RefundChargeID),
	}
	re, err := sr.CLIENT.Refunds.New(params)

	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			switch stripeErr.Code {

			case stripe.ErrorCodeCardDeclined:
				logging.Logger.Error("Card declined Error:", stripeErr.Error())
			case stripe.ErrorCodeExpiredCard:
				logging.Logger.Error("Card expired Error:", stripeErr.Error())
			case stripe.ErrorCodeIncorrectCVC:
				logging.Logger.Error("CVC incorrect Error:", stripeErr.Error())
			}
			logging.Logger.Error("Other Stripe error occurred:", stripeErr.Error())
		} else {
			logging.Logger.Error("Error occurred %v\n", err.Error())
		}
		return nil, err
	}
	return re, nil
}
