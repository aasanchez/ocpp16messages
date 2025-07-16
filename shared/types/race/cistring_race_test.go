package sharedtypes_race

import (
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// TestCiString20Type_RaceCondition is designed to detect race conditions
// for the CiString20Type. It concurrently calls the Value and Validate methods
// to ensure that no data races occur when the object is accessed from multiple goroutines.
func TestCiString20Type_RaceCondition(t *testing.T) {
	t.Parallel()

	// Create a valid CiString20Type instance that will be shared across goroutines.
	ciString, err := st.SetCiString20Type("shared-value")
	if err != nil {
		t.Fatalf("Failed to create CiString20Type: %v", err)
	}

	// Use a WaitGroup to wait for all goroutines to complete.
	var wg sync.WaitGroup
	// Set the number of concurrent goroutines to run.
	numGoroutines := 10
	wg.Add(numGoroutines)

	// Launch multiple goroutines to concurrently access the ciString instance.
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			// Concurrently call the Value and Validate methods.
			// The race detector will report an error if there is a data race.
			_ = ciString.Value()
			_ = ciString.Validate()
		}()
	}

	// Wait for all goroutines to finish.
	wg.Wait()
}

// TestCiString25Type_RaceCondition is designed to detect race conditions
// for the CiString25Type. It concurrently calls the Value and Validate methods
// to ensure that no data races occur when the object is accessed from multiple goroutines.
func TestCiString25Type_RaceCondition(t *testing.T) {
	t.Parallel()

	ciString, err := st.SetCiString25Type("shared-value-for-25")
	if err != nil {
		t.Fatalf("Failed to create CiString25Type: %v", err)
	}

	var wg sync.WaitGroup

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			_ = ciString.Value()
			_ = ciString.Validate()
		}()
	}

	wg.Wait()
}

// TestCiString50Type_RaceCondition is designed to detect race conditions
// for the CiString50Type. It concurrently calls the Value and Validate methods
// to ensure that no data races occur when the object is accessed from multiple goroutines.
func TestCiString50Type_RaceCondition(t *testing.T) {
	t.Parallel()

	ciString, err := st.SetCiString50Type("shared-value-for-50-type-race-condition-test")
	if err != nil {
		t.Fatalf("Failed to create CiString50Type: %v", err)
	}

	var wg sync.WaitGroup

	numGoroutines := 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			_ = ciString.Value()
			_ = ciString.Validate()
		}()
	}

	wg.Wait()
}

// TestCiString255Type_RaceCondition is designed to detect race conditions
// for the CiString255Type. It concurrently calls the Value and Validate methods
// to ensure that no data races occur when the object is accessed from multiple goroutines.
func TestCiString255Type_RaceCondition(t *testing.T) {
	t.Parallel()

	ciString, err := st.SetCiString255Type("shared-value-for-255-type-race-condition-test")
	if err != nil {
		t.Fatalf("Failed to create CiString255Type: %v", err)
	}

	var wg sync.WaitGroup

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			_ = ciString.Value()
			_ = ciString.Validate()
		}()
	}

	wg.Wait()
}

// TestCiString500Type_RaceCondition is designed to detect race conditions
// for the CiString500Type. It concurrently calls the Value and Validate methods
// to ensure that no data races occur when the object is accessed from multiple goroutines.
func TestCiString500Type_RaceCondition(t *testing.T) {
	t.Parallel()

	ciString, err := st.SetCiString500Type("shared-value-for-500-type-race-condition-test")
	if err != nil {
		t.Fatalf("Failed to create CiString500Type: %v", err)
	}

	var wg sync.WaitGroup

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			_ = ciString.Value()
			_ = ciString.Validate()
		}()
	}

	wg.Wait()
}
