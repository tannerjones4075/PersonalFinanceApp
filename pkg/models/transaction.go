package models

type Transaction struct {
	Id            string  `json:"Id"`
	TransDate     string  `json:"Transaction Date"` // TODO: Need to know how to handle dates
	Merchant      string  `json:"Merchant"`
	Amount        float64 `json:"Amount"`
	AccountNumber int     `json:"Account Number"`
	Category      string  `json:"Category"`
}

// TODO: Add categories
