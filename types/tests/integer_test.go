package types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testIntegerValue    = 73
	testIntegerOverflow = 65536
	testIntegerMax      = 65535
	testIntegerZero     = 0
	testIntegerNegative = -10
)

func TestNewInteger(t *testing.T) {
	t.Parallel()

	integer, err := st.NewInteger(testIntegerValue)
	if err != nil {
		t.Errorf(st.ErrorExpectedError, err)
	}

	got := integer.Value()
	if got != testIntegerValue {
		t.Errorf(st.ErrorMismatch, got, testIntegerValue)
	}
}

func TestNewInteger_Overflow(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger(testIntegerOverflow)
	if err == nil {
		t.Error(err)
	}
}

func TestNewInteger_Negative(t *testing.T) {
	t.Parallel()

	_, err := st.NewInteger(testIntegerNegative)
	if err == nil {
		t.Error(err)
	}
}

func TestInteger_String(t *testing.T) {
	t.Parallel()

	integer, _ := st.NewInteger(testIntegerValue)
	got := integer.String()

	if got != "73" {
		t.Errorf(st.ErrorMismatch, got, "73")
	}
}

func TestInteger_String_Zero(t *testing.T) {
	t.Parallel()

	integer, _ := st.NewInteger(testIntegerZero)

	got := integer.String()
	want := "0"

	if got != want {
		t.Errorf(st.ErrorMismatch, got, want)
	}
}

func TestInteger_String_MaxValue(t *testing.T) {
	t.Parallel()

	integer, _ := st.NewInteger(testIntegerMax)

	got := integer.String()
	want := "65535"

	if got != want {
		t.Errorf(st.ErrorMismatch, got, want)
	}
}
