package main

import (
	"net/http"

	"github.com/chinmay2197/stripe-integration-api/logging"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/client"
)

var Client *client.API

func main() {
	router := mux.NewRouter()

	logging.Logger.Info("=================Mux Router Start===================")
	router.HandleFunc("/api/v1/create_charge", createCharge).Methods("POST")
	router.HandleFunc("/api/v1/capture_charge/{chargeId}", captureCharge).Methods("POST")
	router.HandleFunc("/api/v1/create_refund/{chargeId}", createRefund).Methods("POST")
	router.HandleFunc("/api/v1/get_charges", getCharges).Methods("GET")
	router.Use(loggingMiddleware)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		logging.Logger.Fatal("Mux router error", err)
	}
	logging.Logger.Info("=================Mux Router End===================")

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Logger.Info(r.RequestURI)
		apiKey := r.Header.Get("x-api-key")
		if apiKey != "" {
			Client = &client.API{}
			Client.Init(apiKey, nil)
			next.ServeHTTP(w, r)
		} else {
			logging.Logger.Error("Header x-api-key is not valid or empty string")
			http.Error(w, GetStripeError("Header x-api-key is not valid or empty string", http.StatusForbidden).Error(), http.StatusForbidden)
		}
	})
}
