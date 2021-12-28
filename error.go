package main

import (
	"errors"

	"github.com/stripe/stripe-go/v72"
)

func GetStripeError(message string, status int) *stripe.Error {
	return &stripe.Error{
		Err:            errors.New(message),
		Msg:            message,
		HTTPStatusCode: status,
		Type:           stripe.ErrorTypeAPI,
	}
}
