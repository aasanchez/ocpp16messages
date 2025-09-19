//go:build race

package sharedtypes_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// randomSleep increases goroutine interleaving to maximize race detection.
func randomSleepDT() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
}

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

func TestDateTimeRace_SharedInstanceStringValue(t *testing.T) {
	t.Parallel()
	dt, err := st.SetDateTime("2025-09-19T12:34:56.789Z")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
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

// Document: If DateTime ever becomes mutable, add tests for concurrent mutation and reading.
