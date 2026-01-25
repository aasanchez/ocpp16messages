package types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testIntegerValue    = 73
	testIntegerValueStr = "73"
)

func TestNewInteger(t *testing.T) {
	t.Parallel()

	integer, err := st.NewInteger(testIntegerValueStr)
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

func TestInteger_String(t *testing.T) {
	t.Parallel()

	integer, _ := st.NewInteger(testIntegerValueStr)
	got := integer.String()

	if got != testIntegerValueStr {
		t.Errorf(st.ErrorMismatch, got, testIntegerValueStr)
	}
}

func TestInteger_String_Zero(t *testing.T) {
	t.Parallel()

	integer, _ := st.NewInteger("0")

	got := integer.String()
	want := "0"

	if got != want {
		t.Errorf(st.ErrorMismatch, got, want)
	}
}

func TestInteger_String_MaxValue(t *testing.T) {
	t.Parallel()

	integer, _ := st.NewInteger("65535")

	got := integer.String()
	want := "65535"

	if got != want {
		t.Errorf(st.ErrorMismatch, got, want)
	}
}
