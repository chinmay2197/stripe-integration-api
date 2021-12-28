package stripecharge

import (
	"github.com/chinmay2197/stripe-integration-api/logging"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

type StripeCharge struct {
	CLIENT              *client.API
	ChargeRequestParams ChargeRequest
	CaptureChargeID     string
	ChargeListParams    ChargeListParams
}

// create stripe charge
func (sc StripeCharge) CreateStripeCharge() (*stripe.Charge, error) {
	logging.Logger.Info("function start: CreateStripeCharge")
	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(sc.ChargeRequestParams.Amount),
		Currency:    stripe.String(sc.ChargeRequestParams.Currency),
		Description: stripe.String(sc.ChargeRequestParams.Description),
		Customer:    stripe.String(sc.ChargeRequestParams.Customer),
		Capture:     stripe.Bool(sc.ChargeRequestParams.Capture),
	}
	c, err := sc.CLIENT.Charges.New(chargeParams)

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
	logging.Logger.Info("function end: CreateStripeCharge")
	return c, nil
}

// capture stripe charge
func (sc StripeCharge) CaptureStripeCharge() (*stripe.Charge, error) {
	logging.Logger.Info("function start: CaptureStripeCharge")
	c, err := sc.CLIENT.Charges.Capture(
		sc.CaptureChargeID,
		nil,
	)

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
	logging.Logger.Info("function end: CaptureStripeCharge")
	return c, nil
}

// get charges list with pagination or auto pagination
func (sc StripeCharge) GetCharges() []*stripe.Charge {
	logging.Logger.Info("function start: GetCharges")
	chargeList := []*stripe.Charge{}
	var listParams stripe.ListParams
	listParams.Limit = stripe.Int64(sc.ChargeListParams.Limit)
	listParams.Single = *stripe.Bool(sc.ChargeListParams.Single)
	if sc.ChargeListParams.StartingAfter != "" {
		listParams.StartingAfter = &sc.ChargeListParams.StartingAfter
	}
	if sc.ChargeListParams.EndingBefore != "" {
		listParams.EndingBefore = &sc.ChargeListParams.EndingBefore
	}
	params := &stripe.ChargeListParams{
		ListParams: listParams,
	}
	charges := sc.CLIENT.Charges.List(params)

	for charges.Next() {
		c := charges.Charge()
		chargeList = append(chargeList, c)
	}
	logging.Logger.Info("function end: GetCharges")
	return chargeList
}
