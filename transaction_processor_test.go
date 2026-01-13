package main

import (
	"sync"
	"testing"
)

func TestProcessTransaction_Idempotency(t *testing.T) {
	processor := NewTransactionProcessor()
	tx := Transaction{ID: "tx1", Amount: 100.0}

	// First processing should succeed
	err := processor.ProcessTransaction(tx)
	if err != nil {
		t.Errorf("Expected first processing to succeed, but got error: %v", err)
	}

	// Second processing should fail (idempotency check)
	err = processor.ProcessTransaction(tx)
	if err == nil {
		t.Error("Expected second processing to fail due to idempotency, but it succeeded")
	} else if err.Error() != "transaction already processed" {
		t.Errorf("Expected error 'transaction already processed', got: %v", err)
	}
}

func TestProcessTransaction_Concurrent(t *testing.T) {
	processor := NewTransactionProcessor()
	tx := Transaction{ID: "tx_concurrent", Amount: 50.0}

	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			err := processor.ProcessTransaction(tx)
			if err == nil {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	if successCount != 1 {
		t.Errorf("Expected exactly 1 successful processing, but got %d", successCount)
	}
}
