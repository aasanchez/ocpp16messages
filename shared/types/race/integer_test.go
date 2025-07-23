package sharedtypes_test

import (
	"sync"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func RaceCondition_IntegerValue(t *testing.T) {
	t.Parallel()

	var wait sync.WaitGroup

	var shared st.Integer

	shared, err := st.SetInteger("42")
	if err != nil {
		t.Fatalf("SetInteger failed: %v", err)
	}

	for range 1000 {
		wait.Add(2)

		go func() {
			defer wait.Done()

			_ = shared.Value()
		}()

		go func() {
			defer wait.Done()

			shared, _ = st.SetInteger("42")
		}()
	}

	wait.Wait()
}
