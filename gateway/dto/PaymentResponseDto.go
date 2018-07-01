package dto

type PaymentResponseDto struct {
	ReferenceId string `json:"referenceId"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func CreatePaymentResponseDto(referenceId string, status string, description string) PaymentResponseDto {
	var paymentResponseDto PaymentResponseDto
	paymentResponseDto.ReferenceId = referenceId
	paymentResponseDto.Status = status
	paymentResponseDto.Description = description
	return paymentResponseDto
}
