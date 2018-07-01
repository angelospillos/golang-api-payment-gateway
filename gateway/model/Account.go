package model

import (
	"time"
	"math/rand"
	"fmt"
)

type Account struct {
	Id               string    `json:"id"`
	Available        int64     `json:"available"`
	Blocked          int64     `json:"blocked"`
	Deposited        int64     `json:"deposited"`
	Withdrawn        int64     `json:"withdrawn"`
	Currency         string    `json:"currency"`
	CardName         string    `json:"card_name"`
	CardType         string    `json:"card_type"`
	CardNumber       int       `json:"card_number"`
	CardExpiryMonth  int       `json:"card_expiry_month"`
	CardExpiryYear   int       `json:"card_expiry_year"`
	CardSecurityCode int       `json:"card_security_code"`
	Statement        []string  `json:"statement"`
	CreationTime     time.Time `json:"creation_time"`
}

func GenerateAccount() Account {

	var identifier = rand.Intn(4666778181156223-4666000000000000) + 4666000000000000
	var account Account

	account.Id = fmt.Sprintf("%v", identifier)
	account.Available = 0
	account.Blocked = 0
	account.Deposited = 0
	account.Withdrawn = 0
	account.Currency = "GBP"
	account.CardName = "Mr Payment"
	account.CardType = "VISA"
	account.CardNumber = identifier
	account.CardExpiryMonth = rand.Intn(12-1) + 1
	account.CardExpiryYear = rand.Intn(24-18) + 18
	account.CardSecurityCode = rand.Intn(999-100) + 100
	account.CreationTime = time.Now()

	return account
}

type AccountStatementDto struct {
	Statement []Payment `json:"statement"`
}
