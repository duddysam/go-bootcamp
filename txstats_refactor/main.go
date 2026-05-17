package main

import (
	"fmt"
	"txstats/models"
)

type TransactionSet []models.Transaction

func (ts TransactionSet) Total() float64 {

	total := 0.0

	for _, v := range ts {
		total += v.Amount
	}

	return total

}

func (ts TransactionSet) CountByStatus() map[string]int {

	statusMap := map[string]int{
		"approved": 0,
		"declined": 0,
		"pending":  0,
	}

	for _, v := range ts {
		statusMap[v.Status] += 1
	}
	return statusMap
}

func (ts *TransactionSet) Add(txn models.Transaction) {

	*ts = append(*ts, txn)
}

func main() {

	var txns TransactionSet = []models.Transaction{
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

	newtxn := models.Transaction{
		ID: "tx10", Merchant: "Paste", Amount: 3211, Currency: "USD", Status: "pending",
	}

	fmt.Println(txns.Total())

	fmt.Println(txns.CountByStatus())
	fmt.Println(txns)

	txns.Add(newtxn)

	fmt.Println(txns)

}
