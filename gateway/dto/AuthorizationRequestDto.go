package dto

type AuthorizationRequestDto struct {
	OrderId          string `json:"orderId"`
	Amount           int64  `json:"amount"`
	Currency         string `json:"currency"`
	CardName         string `json:"card_name"`
	CardNumber       int    `json:"card_number"`
	CardExpiryMonth  int    `json:"card_expiry_month"`
	CardExpiryYear   int    `json:"card_expiry_year"`
	CardSecurityCode int    `json:"card_security_code"`
}
