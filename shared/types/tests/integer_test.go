package types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const testIntegerValue = 73

func Test_sharedtypes_NewInteger(t *testing.T) {
	t.Parallel()

	i, err := st.NewInteger("73")
	if err != nil {
		t.Fatalf("unexpected error from NewInteger: %v", err)
	}

	got := i.Value()
	if got != testIntegerValue {
		t.Errorf("Value() = %d; want %d", got, testIntegerValue)
	}
}

func Test_sharedtypes_NewInteger_Overflow(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger("4294967296")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

func Test_sharedtypes_NewInteger_Negative(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger("-10")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}

func Test_sharedtypes_NewInteger_Alphanumeric(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger("abc")
	if err == nil {
		t.Errorf(st.ErrorExpectedError, err)
	}
}
