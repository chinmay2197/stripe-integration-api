package main

import (
	"encoding/json"
	"net/http"

	"github.com/chinmay2197/stripe-integration-api/logging"
	"github.com/chinmay2197/stripe-integration-api/stripecharge"
	"github.com/chinmay2197/stripe-integration-api/striperefund"
	"github.com/chinmay2197/stripe-integration-api/utils"
	"github.com/gorilla/mux"
)

/*
	Create charge for customer
	`cutomer` is required in api request
*/
func createCharge(response http.ResponseWriter, request *http.Request) {
	logging.Logger.Info("endpoint `create_charge` triggered")
	var chargeRequest *stripecharge.ChargeRequest

	err := json.NewDecoder(request.Body).Decode(&chargeRequest)
	if err != nil {
		logging.Logger.Error("Error occurred while decoding request params Error:", err)
		http.Error(response, GetStripeError("Bad Request params", http.StatusBadRequest).Error(), http.StatusBadRequest)
		return
	}

	var stripeCharge stripecharge.StripeCharge
	stripeCharge.CLIENT = Client
	stripeCharge.ChargeRequestParams = *chargeRequest

	charge, err := stripeCharge.CreateStripeCharge()
	if err != nil {
		logging.Logger.Error("Error occurred while processing charge :", err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(charge)
}

/*
	Capture a charge
	chargeId is required as part of the url
*/
func captureCharge(response http.ResponseWriter, request *http.Request) {
	logging.Logger.Info("endpoint `capture_charge` triggered")
	params := mux.Vars(request)
	if params["chargeId"] == "" {
		logging.Logger.Error("charge id is empty string")
		http.Error(response, GetStripeError("Bad Request params: chargeId is not set", http.StatusBadRequest).Error(), http.StatusBadRequest)
		return
	}
	var stripeCharge stripecharge.StripeCharge

	stripeCharge.CLIENT = Client
	stripeCharge.CaptureChargeID = params["chargeId"]
	charge, err := stripeCharge.CaptureStripeCharge()
	if err != nil {
		logging.Logger.Error("Error occurred while processing charge :", err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(charge)
}

/*
	List charges
	auto-pagination is enabled if `single` is set to false in api request params
	Otherwise, use starting_after and ending_before to navigate to prev and next pages
	`limit` is set to 10 by default
*/
func getCharges(response http.ResponseWriter, request *http.Request) {
	logging.Logger.Info("endpoint `get_charges` triggered")
	var stripeCharge stripecharge.StripeCharge
	var err error
	params := request.URL.Query()
	startingAfter := params.Get("starting_after")
	if params.Get("starting_after") != "" {
		stripeCharge.ChargeListParams.StartingAfter = startingAfter
	}
	endingBefore := params.Get("ending_before")
	if params.Get("ending_before") != "" {
		stripeCharge.ChargeListParams.EndingBefore = endingBefore
	}
	stripeCharge.ChargeListParams.Limit, err = utils.ConvertStringToInt64(params.Get("limit"))
	if err != nil {
		logging.Logger.Error("string to int64 conversion failed for limit param", err)
		stripeCharge.ChargeListParams.Limit = 10 // setting default limit 10 if limit is provided as non int
	}
	if params.Get("single") == "true" {
		stripeCharge.ChargeListParams.Single = true
	} else {
		stripeCharge.ChargeListParams.Single = false
	}

	stripeCharge.CLIENT = Client
	charges := stripeCharge.GetCharges()
	json.NewEncoder(response).Encode(charges)
}

/*
	Create refund for charge
	chargeId is required as part of the requested url
*/
func createRefund(response http.ResponseWriter, request *http.Request) {
	logging.Logger.Info("endpoint `create_refund` triggered")
	params := mux.Vars(request)
	if params["chargeId"] == "" {
		logging.Logger.Error("charge id is empty string")
		http.Error(response, GetStripeError("Bad Request params: chargeId is not set", http.StatusBadRequest).Error(), http.StatusBadRequest)
		return
	}

	var stripeRefund striperefund.StripeRefund
	stripeRefund.RefundChargeID = params["chargeId"]
	stripeRefund.CLIENT = Client
	refund, err := stripeRefund.CreateStripeRefund()
	if err != nil {
		logging.Logger.Error("Error occurred while processing refund :", err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(response).Encode(refund)
}
