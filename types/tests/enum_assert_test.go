package types_test

import (
	"fmt"
	"testing"
)

// enumValidator pairs a Stringer with an IsValid check.
type enumValidator struct {
	value     fmt.Stringer
	isValidFn func() bool
}

func assertEnumValid(
	t *testing.T,
	validator enumValidator,
	wantString string,
) {
	t.Helper()

	if validator.value.String() != wantString {
		t.Errorf("String() = %q, want %q", validator.value.String(), wantString)
	}

	if !validator.isValidFn() {
		t.Errorf("IsValid() = %v, want %v", validator.isValidFn(), true)
	}
}

func assertEnumInvalid(t *testing.T, validator enumValidator) {
	t.Helper()

	if validator.isValidFn() {
		t.Errorf("IsValid() = %v, want false", validator.isValidFn())
	}
}
