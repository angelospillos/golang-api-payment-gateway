package dto

type SuccessiveRequestDto struct {
	OrderId     string `json:"orderId"`
	ReferenceId string `json:"referenceId"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
}
