package model

type CardIssuer struct {
	Bin           string `json:"bin"`
	PaymentSystem string `json:"payment_system"`
	Bank          string `json:"bank"`
	CardType      string `json:"card_type"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	Website       string `json:"website"`
}

type CardIssuers struct {
	Issuers []CardIssuer `json:"issuers"`
}

type CardIssuerStore interface {
	GetCardIssuer(Bin string) (CardIssuer, bool)
}

type Payload struct {
	CardNumber string `json:"card_number,omitempty"`
}

type UnknownCardIssuer struct {
	PaymentSystem string `json:"payment_system"`
}
