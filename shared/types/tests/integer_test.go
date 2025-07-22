package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestHappyValue(t *testing.T) {
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
	_, err := st.SetInteger("4294967296")
	if err == nil {
		t.Fatalf("Expected Error: %v", err)
	}
}

func TestNegative(t *testing.T) {
	_, err := st.SetInteger("-10")
	if err == nil {
		t.Fatalf("Expected Error: %v", err)
	}
}

func TestAlphanumeric(t *testing.T) {
	_, err := st.SetInteger("abc")
	if err == nil {
		t.Fatalf("Expected Error: %v", err)
	}
}
