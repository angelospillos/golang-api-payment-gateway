package model

import (
	"time"
	"gateway/dto"
	"gateway/constant"
	"github.com/google/uuid"
)

type Payment struct {
	Id              string    `json:"id"`
	BusinessId      string    `json:"businessId"`
	OrderId         string    `json:"order_id"`
	Operation       string    `json:"operation"`
	OriginalAmount  int64     `json:"originalAmount"`
	CurrentAmount   int64     `json:"currentAmount"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	Currency        string    `json:"currency"`
	CardName        string    `json:"card_name"`
	CardType        string    `json:"card_type"`
	CardNumber      int       `json:"card_number"`
	CardExpiryMonth int       `json:"card_expiry_month"`
	CardExpiryYear  int       `json:"card_expiry_year"`
	CreationTime    time.Time `json:"creation_time"`
}

func CreateAuthorizationPayment(authorizationRequestDto dto.AuthorizationRequestDto,
	personalAccount Account,
	businessAccount Account,
	status string,
	description string) Payment {

	var payment Payment

	payment.Id = uuid.New().String()
	payment.BusinessId = businessAccount.Id
	payment.OrderId = authorizationRequestDto.OrderId
	payment.Operation = constant.AUTHORIZATION
	payment.OriginalAmount = authorizationRequestDto.Amount
	payment.CurrentAmount = authorizationRequestDto.Amount
	payment.Status = status
	payment.Description = description
	payment.Currency = authorizationRequestDto.Currency
	payment.CardName = authorizationRequestDto.CardName
	payment.CardType = personalAccount.CardType
	payment.CardNumber = personalAccount.CardNumber
	payment.CardExpiryMonth = personalAccount.CardExpiryMonth
	payment.CardExpiryYear = personalAccount.CardExpiryYear
	payment.CreationTime = time.Now()

	return payment
}

func CreateSuccessivePayment(successiveRequestDto dto.SuccessiveRequestDto,
	referencedPayment Payment,
	status string,
	description string) Payment {

	var successivePayment Payment

	successivePayment.Id = uuid.New().String()
	successivePayment.BusinessId = referencedPayment.BusinessId
	successivePayment.OrderId = successiveRequestDto.OrderId
	successivePayment.Operation = successiveRequestDto.Type
	successivePayment.OriginalAmount = successiveRequestDto.Amount
	successivePayment.CurrentAmount = successiveRequestDto.Amount
	successivePayment.Status = status
	successivePayment.Description = description
	successivePayment.Currency = referencedPayment.Currency
	successivePayment.CardName = referencedPayment.CardName
	successivePayment.CardType = referencedPayment.CardType
	successivePayment.CardNumber = referencedPayment.CardNumber
	successivePayment.CardExpiryMonth = referencedPayment.CardExpiryMonth
	successivePayment.CardExpiryYear = referencedPayment.CardExpiryYear
	successivePayment.CreationTime = time.Now()

	return successivePayment
}
