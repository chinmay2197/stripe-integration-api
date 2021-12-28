package stripecharge

type ChargeRequest struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Customer    string `json:"customer"`
	Description string `json:"description"`
	Capture     bool   `json:"capture"`
}

type ChargeListParams struct {
	StartingAfter string `json:"starting_after"`
	EndingBefore  string `json:"ending_before"`
	Limit         int64  `json:"limit"`
	Single        bool   `json:"single"`
}
