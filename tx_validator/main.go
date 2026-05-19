package main

import (
	"errors"
	"fmt"
	"slices"
)

type Transaction struct {
	ID       string
	Merchant string
	Amount   float64
	Currency string
	Status   string
}

type ValidationError struct {
	Field  string
	Value  string
	Reason string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("field: %v", v.Field)
}

func Validate(t Transaction) error {

	var errs []error

	if t.ID == "" {
		errs = append(errs, errors.New("ID cannot be empty"))
	}
	if t.Merchant == "" {
		errs = append(errs, errors.New("Merchant cannot be empty"))
	}
	if t.Amount <= 0 {
		errs = append(errs, ValidationError{Field: "Amount", Value: fmt.Sprintf("%.2f", t.Amount), Reason: "Must be >= 0"})
	}
	if !slices.Contains([]string{"USD", "EUR", "GBP"}, t.Currency) {
		errs = append(errs, errors.New("Currency is invalid"))
	}
	if !slices.Contains([]string{"pending", "approved", "declined"}, t.Status) {
		errs = append(errs, errors.New("Status is invalid"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil

}

func ValidateAll(t []Transaction) []error {
	var errs []error
	for _, v := range t {
		if err := Validate(v); err != nil {
			errs = append(errs, fmt.Errorf("txn %s: %w", v.ID, err))
		}
	}
	if len(errs) > 0 {
		return errs
	} else {
		return nil
	}

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

	var errs []error
	errs = ValidateAll(txns)
	fmt.Println(errs)

	err := Validate(txns[3])
	fmt.Println(err)
	var ValidationError ValidationError
	if errors.As(err, &ValidationError) {
		fmt.Println("errors.As returned true")
		fmt.Println(ValidationError.Field)
	}

}
