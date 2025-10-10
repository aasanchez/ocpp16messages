//go:build race

// OCPP 1.6 DateTime race tests.
// Verifies RFC3339 parse, String, Value
// under heavy concurrency and jitter.
package sharedtypes_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// randomSleepDT adds jitter to goroutines
// to expose races under load.
func randomSleepDT() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
}

// Parses many inputs in parallel, then
// calls String and Value on each.
// Mirrors OCPP 1.6 time handling.
func TestDateTimeRace_ConcurrentSetAndStringValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	inputs := []string{
		"2025-09-19T12:34:56Z", "2025-09-19T12:34:56.789Z", "0001-01-01T00:00:00Z",
		"9999-12-31T23:59:59Z", "2016-12-31T23:59:60Z", "not-a-date", "", "2025-13-01T00:00:00Z",
	}
	for i := 0; i < 100; i++ {
		for _, input := range inputs {
			wg.Add(1)
			go func(val string) {
				defer wg.Done()
				randomSleepDT()
				dt, err := st.SetDateTime(val)
				if err == nil {
					_ = dt.String()
					_ = dt.Value()
				}
			}(input)
		}
	}
	wg.Wait()
}

// Validates String and Value calls on a
// single parsed instance across goroutines.
// Ensures read-only access is race safe.
func TestDateTimeRace_SharedInstanceStringValue(t *testing.T) {
	t.Parallel()
	dt, err := st.SetDateTime("2025-09-19T12:34:56.789Z")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randomSleepDT()
			str := dt.String()
			if str == "" {
				t.Errorf("String() returned empty string")
			}
			v := dt.Value()
			if v.IsZero() {
				t.Errorf("Value() returned zero time")
			}
		}()
	}
	wg.Wait()
}

// Interleaves parse with later String and
// Value calls using mixed input quality.
// Reflects OCPP 1.6 field diversity.
func TestDateTimeRace_InterleavedSetStringValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	inputs := []string{
		"2025-09-19T12:34:56Z", "2025-09-19T12:34:56.789Z", "0001-01-01T00:00:00Z",
		"9999-12-31T23:59:59Z", "2016-12-31T23:59:60Z", "not-a-date", "", "2025-13-01T00:00:00Z",
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := inputs[i%len(inputs)]
			dt, err := st.SetDateTime(val)
			randomSleepDT()
			if err == nil {
				_ = dt.String()
				_ = dt.Value()
			}
		}(i)
	}
	wg.Wait()
}

// Hammers parse, String, and Value with
// thousands of goroutines and inputs.
// Models bursty OCPP timestamp traffic.
func TestDateTimeRace_Stress(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	inputs := []string{
		"2025-09-19T12:34:56Z", "2025-09-19T12:34:56.789Z", "0001-01-01T00:00:00Z",
		"9999-12-31T23:59:59Z", "2016-12-31T23:59:60Z", "not-a-date", "", "2025-13-01T00:00:00Z",
	}
	nGoroutines := 2000
	for i := 0; i < nGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := inputs[rand.Intn(len(inputs))]
			dt, err := st.SetDateTime(val)
			randomSleepDT()
			if err == nil {
				_ = dt.String()
				_ = dt.Value()
			}
		}(i)
	}
	wg.Wait()
}
