package main

import (
	"fmt"
	"txstats/models"
)

func Total(txns []models.Transaction) float64 {

	var total float64

	for _, v := range txns {
		total += v.Amount
	}

	return total

}

func main() {

	txns := []models.Transaction{
		{ID: "tx1", Merchant: "Boeing", Amount: 500, Currency: "USD", Status: "approved"},
		{ID: "tx2", Merchant: "Boeing", Amount: 700, Currency: "USD", Status: "pending"},
		{ID: "tx3", Merchant: "American Express", Amount: 50, Currency: "USD", Status: "pending"},
		{ID: "tx4", Merchant: "HooHah", Amount: 1000, Currency: "USD", Status: "declined"},
		{ID: "tx5", Merchant: "Boeing", Amount: 3000, Currency: "USD", Status: "approved"},
		{ID: "tx6", Merchant: "Delta", Amount: 250, Currency: "USD", Status: "approved"},
		{ID: "tx7", Merchant: "Tesla", Amount: 5500, Currency: "USD", Status: "approved"},
		{ID: "tx8", Merchant: "Tesla", Amount: 1100, Currency: "USD", Status: "pending"},
		{ID: "tx9", Merchant: "Tomato", Amount: 750, Currency: "USD", Status: "pending"},
	}

	total := Total(txns)
	fmt.Printf("Total of transactions is: %.2f\n", total)

}
