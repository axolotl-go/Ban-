package models

type Movement struct {
	IdUser     uint    `json:"id_User"`
	Amount     float64 `json:"amount"`
	Email      string  `json:"Email"`
	CardNumber string  `json:"cardNumber"`
}
