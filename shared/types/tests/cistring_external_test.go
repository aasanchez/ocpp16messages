package types_test

import (
	"strings"
	"testing"

	types "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	errCiString20CreationFailed  = "CiString20 creation failed: %v"
	errCiString25CreationFailed  = "CiString25 creation failed: %v"
	errCiString50CreationFailed  = "CiString50 creation failed: %v"
	errCiString255CreationFailed = "CiString255 creation failed: %v"
	errCiString500CreationFailed = "CiString500 creation failed: %v"
)

func TestCiString20Type_CreateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)
	if _, err := types.SetCiString20(input); err != nil {
		t.Fatalf(errCiString20CreationFailed, err)
	}
}

func TestCiString20Type_ValidateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, err := types.SetCiString20(input)
	if err != nil {
		t.Fatalf(errCiString20CreationFailed, err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString20 validation failed: %v", err)
	}
}

func TestCiString20Type_ValueReturnsOriginal(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("A", 20)

	ciStr, err := types.SetCiString20(input)
	if err != nil {
		t.Fatalf(errCiString20CreationFailed, err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString25Type_CreateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)
	if _, err := types.SetCiString25(input); err != nil {
		t.Fatalf(errCiString25CreationFailed, err)
	}
}

func TestCiString25Type_ValidateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, err := types.SetCiString25(input)
	if err != nil {
		t.Fatalf(errCiString25CreationFailed, err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString25 validation failed: %v", err)
	}
}

func TestCiString25Type_ValueReturnsOriginal(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("B", 25)

	ciStr, err := types.SetCiString25(input)
	if err != nil {
		t.Fatalf(errCiString25CreationFailed, err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString50Type_CreateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)
	if _, err := types.SetCiString50(input); err != nil {
		t.Fatalf(errCiString50CreationFailed, err)
	}
}

func TestCiString50Type_ValidateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, err := types.SetCiString50(input)
	if err != nil {
		t.Fatalf(errCiString50CreationFailed, err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString50 validation failed: %v", err)
	}
}

func TestCiString50Type_ValueReturnsOriginal(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("C", 50)

	ciStr, err := types.SetCiString50(input)
	if err != nil {
		t.Fatalf(errCiString50CreationFailed, err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString255Type_CreateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	if _, err := types.SetCiString255(input); err != nil {
		t.Fatalf(errCiString255CreationFailed, err)
	}
}

func TestCiString255Type_ValidateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, err := types.SetCiString255(input)
	if err != nil {
		t.Fatalf(errCiString255CreationFailed, err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString255 validation failed: %v", err)
	}
}

func TestCiString255Type_ValueReturnsOriginal(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("D", 255)

	ciStr, err := types.SetCiString255(input)
	if err != nil {
		t.Fatalf(errCiString255CreationFailed, err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

func TestCiString500Type_CreateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	if _, err := types.SetCiString500(input); err != nil {
		t.Fatalf(errCiString500CreationFailed, err)
	}
}

func TestCiString500Type_ValidateSucceeds(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, err := types.SetCiString500(input)
	if err != nil {
		t.Fatalf(errCiString500CreationFailed, err)
	}

	if err := ciStr.Validate(); err != nil {
		t.Errorf("CiString500 validation failed: %v", err)
	}
}

func TestCiString500Type_ValueReturnsOriginal(t *testing.T) {
	t.Parallel()

	input := strings.Repeat("E", 500)

	ciStr, err := types.SetCiString500(input)
	if err != nil {
		t.Fatalf(errCiString500CreationFailed, err)
	}

	if ciStr.Value() != input {
		t.Errorf(types.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}
