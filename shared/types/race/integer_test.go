package sharedtypes_race

import (
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func RaceCondition_IntegerValue(t *testing.T) {
	var wg sync.WaitGroup

	var shared st.Integer

	shared, err := st.SetInteger("42")
	if err != nil {
		t.Fatalf("SetInteger failed: %v", err)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			_ = shared.Value()
		}()

		go func() {
			defer wg.Done()
			shared, _ = st.SetInteger("42")
		}()
	}

	wg.Wait()
}
