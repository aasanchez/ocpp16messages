package sharedtypes_test

import (
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestRaceCondition_IntegerValue(t *testing.T) {
	t.Parallel()

	var (
		mu        sync.RWMutex
		shared    st.Integer
		waitGroup sync.WaitGroup
	)

	v42, err := st.SetInteger("42")
	if err != nil {
		t.Fatalf("SetInteger(42): %v", err)
	}

	v43, err := st.SetInteger("43")
	if err != nil {
		t.Fatalf("SetInteger(43): %v", err)
	}

	shared = v42

	for i := range 1000 {
		waitGroup.Add(2)

		go func() { // reader
			defer waitGroup.Done()

			mu.RLock()

			_ = shared.Value()

			mu.RUnlock()
		}()

		go func(i int) { // writer
			defer waitGroup.Done()

			mu.Lock()

			if i%2 == 0 {
				shared = v43
			} else {
				shared = v42
			}

			mu.Unlock()
		}(i)
	}

	waitGroup.Wait()
}
