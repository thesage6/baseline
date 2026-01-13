package main

import (
	"errors"
	"sync"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID     string
	Amount float64
}

// TransactionProcessor handles the processing of transactions
type TransactionProcessor struct {
	processedIDs sync.Map
}

// NewTransactionProcessor creates a new TransactionProcessor
func NewTransactionProcessor() *TransactionProcessor {
	return &TransactionProcessor{}
}

// ProcessTransaction processes a transaction ensuring idempotency
func (tp *TransactionProcessor) ProcessTransaction(tx Transaction) error {
	// Check if transaction has already been processed
	if _, loaded := tp.processedIDs.LoadOrStore(tx.ID, true); loaded {
		return errors.New("transaction already processed")
	}

	// In a real application, you would do the actual processing here
	// e.g., charging a card, updating a database, etc.

	return nil
}
