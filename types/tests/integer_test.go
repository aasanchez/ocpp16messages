package types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

const testIntegerValue = 73

func TestNewInteger(t *testing.T) {
	t.Parallel()

	i, err := st.NewInteger("73")
	if err != nil {
		t.Errorf(st.ErrorExpectedError, err)
	}

	got := i.Value()
	if got != testIntegerValue {
		t.Errorf(st.ErrorMismatch, got, testIntegerValue)
	}
}

func TestNewInteger_Overflow(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger("4294967296")
	if err == nil {
		t.Error(err)
	}
}

func TestNewInteger_Negative(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger("-10")
	if err == nil {
		t.Error(err)
	}
}

func TestNewInteger_Alphanumeric(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger("abc")
	if err == nil {
		t.Error(err)
	}
}
