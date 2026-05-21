package main

import (
	"errors"
	"fmt"
)

type Transaction struct {
	ID       string
	Merchant string
	Amount   float64
	Currency string
	Status   string
}

type TransactionStore interface {
	Save(t Transaction) error
	Get(id string) (Transaction, error)
	All() []Transaction
}

type InMemoryStore struct {
	store map[string]Transaction
}

func (ims *InMemoryStore) Save(t Transaction) error {
	ims.store[t.ID] = t

	return nil
}

func (ims *InMemoryStore) Get(id string) (Transaction, error) {

	for k, v := range ims.store {
		if k == id {
			return v, nil
		}
	}

	return Transaction{}, errors.New("id does not exist")
}

func (ims *InMemoryStore) All() []Transaction {
	var txns []Transaction

	for _, v := range ims.store {

		txns = append(txns, v)
	}

	return txns
}

type LoggingStore struct {
	ts TransactionStore
}

func main() {

	fmt.Println("Hello World")

	txns := []Transaction{
		{ID: "tx1", Merchant: "Boeing", Amount: 500, Currency: "USD", Status: "approved"},
		{ID: "tx2", Merchant: "Boeing", Amount: 700, Currency: "USD", Status: "pending"},
		{ID: "tx3", Merchant: "American Express", Amount: 50, Currency: "USD", Status: "pending"},
		{ID: "tx4", Merchant: "HooHah", Amount: 0, Currency: "USD", Status: "declined"},
		{ID: "tx5", Merchant: "Boeing", Amount: 3000, Currency: "USD", Status: "approved"},
		{ID: "tx6", Merchant: "Delta", Amount: 250, Currency: "USD", Status: "approved"},
		{ID: "tx7", Merchant: "Tesla", Amount: 5500, Currency: "USD", Status: "pproved"},
		{ID: "tx8", Merchant: "Tesla", Amount: 1100, Currency: "USD", Status: "pending"},
		{ID: "", Merchant: "Tomato", Amount: 750, Currency: "ASD", Status: "pending"},
	}

	fmt.Println(txns)

}
