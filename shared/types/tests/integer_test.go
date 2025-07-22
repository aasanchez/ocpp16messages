package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestHappyValue(t *testing.T) {
	t.Parallel()

	i, err := st.SetInteger("42")
	if err != nil {
		t.Fatalf("unexpected error from SetInteger: %v", err)
	}

	got := i.Value()
	if got != 42 {
		t.Errorf("Value() = %d; want 42", got)
	}
}

func TestOverflow(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("4294967296")
	if err == nil {
		t.Fatalf(st.ErrorExpectedError, err)
	}
}

func TestNegative(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("-10")
	if err == nil {
		t.Fatalf(st.ErrorExpectedError, err)
	}
}

func TestAlphanumeric(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("abc")
	if err == nil {
		t.Fatalf(st.ErrorExpectedError, err)
	}
}
