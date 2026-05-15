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

func CountByStatus(txns []models.Transaction) map[string]int {
	statusMap := map[string]int{
		"approved": 0,
		"declined": 0,
		"pending":  0,
	}

	for _, v := range txns {
		statusMap[v.Status] += 1
	}

	return statusMap
}

func LargestByMerchant(txns []models.Transaction) map[string]models.Transaction {

	largestTxnMap := map[string]models.Transaction{}

	for _, v := range txns {
		_, ok := largestTxnMap[v.Merchant]
		if !ok {
			largestTxnMap[v.Merchant] = v
		} else if v.Amount > largestTxnMap[v.Merchant].Amount {
			largestTxnMap[v.Merchant] = v
		}
	}

	return largestTxnMap
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

	statusCount := CountByStatus(txns)
	fmt.Println("Total count of each status type:", statusCount)

	largestTxns := LargestByMerchant(txns)
	fmt.Println("Largest transaction for each merchant:", largestTxns)

}
