package types_test

import (
	"errors"
	"strings"
	"testing"

	types "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestCiString20Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := "ABCDEFGHIJKLMNOPQRST"
	ciStr, err := types.SetCiString20(input)

	if err != nil {
		t.Fatalf("CiString20 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString20 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	_, err := types.SetCiString20(strings.Repeat("X", 21))
	if !errors.Is(err, types.ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength for >20 chars, got: %v", err)
	}
}

func TestCiString25Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)
	ciStr, err := types.SetCiString25(input)

	if err != nil {
		t.Fatalf("CiString25 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString25 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString50Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)
	ciStr, err := types.SetCiString50(input)

	if err != nil {
		t.Fatalf("CiString50 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString50 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString255Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)
	ciStr, err := types.SetCiString255(input)

	if err != nil {
		t.Fatalf("CiString255 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString255 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString500Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)
	ciStr, err := types.SetCiString500(input)

	if err != nil {
		t.Fatalf("CiString500 creation failed: %v", err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString500 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}
