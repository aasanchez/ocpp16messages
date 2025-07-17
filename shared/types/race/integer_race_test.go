package sharedtypes_test

import (
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// TestInteger_RaceCondition detects data races when accessing an Integer concurrently.
// It calls the Value method from multiple goroutines to ensure thread-safety.
func TestInteger_RaceCondition(t *testing.T) {
	t.Parallel()

	// Create a shared Integer instance.
	integer, err := st.SetInteger("12345")
	if err != nil {
		t.Fatalf("failed to create Integer: %v", err)
	}

	var wait sync.WaitGroup

	const goroutines = 10

	wait.Add(goroutines)

	for range make([]struct{}, goroutines) {
		go func() {
			defer wait.Done()

			_ = integer.Value()
		}()
	}

	wait.Wait()
}
