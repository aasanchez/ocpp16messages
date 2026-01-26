//go:build race

package race

import (
	"fmt"
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	workers          = 32
	iterations       = 500
	ciStringTemplate = "TAG-%d-%d"
)

func TestRace_NewCiString20Type(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(worker int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				s := fmt.Sprintf(ciStringTemplate, worker, j)
				value, err := st.NewCiString20Type(s)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				_ = value.String()
			}
		}(i)
	}

	wg.Wait()
}

func TestRace_NewInteger(t *testing.T) {
	t.Parallel()

	const maxValue = 65535

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(worker int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				n := (worker + j) % maxValue
				value, err := st.NewInteger(n)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				_ = value.String()
			}
		}(i)
	}

	wg.Wait()
}
