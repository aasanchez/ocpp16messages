package sharedtypes_test

import (
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// TestDateTime_RaceCondition is designed to detect race conditions
// for the DateTime type. It concurrently calls the Value and String methods
// to ensure that no data races occur when the object is accessed from multiple goroutines.
func TestDateTime_RaceCondition(t *testing.T) {
	t.Parallel()

	dt, err := st.SetDateTime("2013-02-01T20:00:00.000Z")
	if err != nil {
		t.Fatalf("failed to parse DateTime: %v", err)
	}

	var wg sync.WaitGroup

	numGoroutines := 10
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()

			_ = dt.Value()
			_ = dt.String()
		}()
	}

	wg.Wait()
}
