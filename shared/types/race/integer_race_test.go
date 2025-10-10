//go:build race

package sharedtypes_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func randomSleep() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
}

func TestIntegerRace_ConcurrentSetAndValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	inputs := []string{
		"0", "1", "42", "65535", "32768", "12345", "9999", "00001", "65536", "-1", "notanint", "", "18446744073709551615",
	}
	for i := 0; i < 100; i++ {
		for _, input := range inputs {
			wg.Add(1)
			go func(val string) {
				defer wg.Done()
				randomSleep()
				integer, err := st.SetInteger(val)
				if err == nil {
					_ = integer.Value()
				}
			}(input)
		}
	}
	wg.Wait()
}

func TestIntegerRace_SharedInstanceValue(t *testing.T) {
	t.Parallel()
	integer, err := st.SetInteger("12345")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randomSleep()
			v := integer.Value()
			if v != 12345 {
				t.Errorf("expected 12345, got %d", v)
			}
		}()
	}
	wg.Wait()
}

func TestIntegerRace_InterleavedSetAndValue(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := strconv.Itoa(i % 65536)
			integer, err := st.SetInteger(val)
			if err != nil {
				return
			}
			randomSleep()
			v := integer.Value()
			if v != uint16(i%65536) {
				t.Errorf("expected %d, got %d", i%65536, v)
			}
		}(i)
	}
	wg.Wait()
}

func TestIntegerRace_Stress(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup
	nGoroutines := 2000
	for i := 0; i < nGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val := strconv.Itoa(rand.Intn(65536))
			integer, err := st.SetInteger(val)
			randomSleep()
			if err == nil {
				_ = integer.Value()
			}
		}(i)
	}
	wg.Wait()
}
