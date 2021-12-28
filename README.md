# stripe-integration-api
Backend api integration with stripe payment gateway

# Dependencies
* stripe account in test mode
* Go 1.17

# To Build the project
Go to project root dir and run
* Go build && ./stripe-integration-api

The api server will be started running on 8000 port 

# Endpoints
* POST /api/v1/create_charge
* POST /api/v1/capture_charge/{chargeId}
* POST /api/v1/create_refund/{chargeId}
* GET /api/v1/get_charges

# Postman Collection
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/1637f3ae5c1608e48e49)