package models

type Transaction struct {
	ID       string
	Merchant string
	Amount   float64
	Currency string
	Status   string // "approved", "declined", "pending"
}
