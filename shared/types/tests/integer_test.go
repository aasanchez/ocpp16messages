package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestExternal_SharedTypes_integer_Default(t *testing.T) {
	t.Parallel()

	i, err := st.SetInteger("73")
	if err != nil {
		t.Fatalf("unexpected error from SetInteger: %v", err)
	}

	got := i.Value()
	if got != 73 {
		t.Errorf("Value() = %d; want 73", got)
	}
}

func TestExternal_SharedTypes_integer_Overflow(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("4294967296")
	if err == nil {
		t.Fatalf(st.ErrorExpectedError, err)
	}
}

func TestExternal_SharedTypes_integer_Negative(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("-10")
	if err == nil {
		t.Fatalf(st.ErrorExpectedError, err)
	}
}

func TestExternal_SharedTypes_integer_Alphanumeric(t *testing.T) {
	t.Parallel()

	_, err := st.SetInteger("abc")
	if err == nil {
		t.Fatalf(st.ErrorExpectedError, err)
	}
}
