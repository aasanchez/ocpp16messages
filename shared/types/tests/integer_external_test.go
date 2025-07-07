package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/shared/types"
)

// TestSetInteger_ValidInteger verifies that a normal numeric string parses without error.
func TestSetInteger_ValidInteger(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger("123"); err != nil {
		t.Fatalf("expected no error for valid integer, got %v", err)
	}
}

// TestSetInteger_Zero verifies that "0" is accepted as a valid integer.
func TestSetInteger_Zero(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger("0"); err != nil {
		t.Fatalf("expected no error for zero, got %v", err)
	}
}

// TestSetInteger_MaxUint32 verifies that the maximum uint32 value is accepted.
func TestSetInteger_MaxUint32(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger("4294967295"); err != nil {
		t.Fatalf("expected no error for max uint32, got %v", err)
	}
}

// TestSetInteger_NegativeInteger verifies that negative strings are rejected.
func TestSetInteger_NegativeInteger(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger("-1"); err == nil {
		t.Fatal("expected error for negative integer, got none")
	}
}

// TestSetInteger_NotANumber verifies that non-numeric strings are rejected.
func TestSetInteger_NotANumber(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger("abc"); err == nil {
		t.Fatal("expected error for non-numeric input, got none")
	}
}

// TestSetInteger_EmptyString verifies that empty input is rejected.
func TestSetInteger_EmptyString(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger(""); err == nil {
		t.Fatal("expected error for empty string, got none")
	}
}

// TestSetInteger_TooLarge verifies that values exceeding uint32 range are rejected.
func TestSetInteger_TooLarge(t *testing.T) {
	t.Parallel()
	if _, err := types.SetInteger("4294967296"); err == nil {
		t.Fatal("expected error for value exceeding uint32 range, got none")
	}
}

// TestInteger_Value verifies the Value() accessor returns the stored uint32.
func TestInteger_Value(t *testing.T) {
	t.Parallel()
	i, err := types.SetInteger("7")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := i.Value(); got != 7 {
		t.Errorf("Value() = %d; want %d", got, 7)
	}
}
