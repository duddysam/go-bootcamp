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

	txn, ok := ims.store[id]

	if !ok {
		return Transaction{}, errors.New("id does not exist")
	}
	return txn, nil
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

func (ls *LoggingStore) Save(t Transaction) error {

	fmt.Printf("Logging txn: %s\n", t.ID)
	return ls.ts.Save(t)
}

func (ls *LoggingStore) Get(id string) (Transaction, error) {

	fmt.Printf("Getting txn: %s\n", id)
	return ls.ts.Get(id)
}

func (ls *LoggingStore) All() []Transaction {
	fmt.Println("Getting all Transactions")
	return ls.ts.All()
}

func ProcessBatch(store TransactionStore, txs []Transaction) (saved int, errs []error) {

	for _, v := range txs {
		err := store.Save(v)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		saved += 1
	}

	return saved, errs

}

func main() {

	ms := &InMemoryStore{
		store: map[string]Transaction{},
	}

	ls := &LoggingStore{
		ts: ms,
	}

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

	saved, errs := ProcessBatch(ls, txns)

	fmt.Println(saved, errs)

	fmt.Println(ms.store)

}
