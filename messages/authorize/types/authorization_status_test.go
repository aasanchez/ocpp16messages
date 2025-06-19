package authorizetypes

import (
	"errors"
	"testing"
)

const (
	errExpectedValidateNil  = "expected Validate() to return nil, got %v"
	errExpectedValidateFail = "expected Validate() to fail, but got nil"
)

func TestAuthorizationStatus_validAccepted(t *testing.T) {
	t.Parallel()

	_, err := AuthorizationStatus("Accepted")
	if err != nil {
		t.Errorf("expected no error for valid status 'Accepted', got: %v", err)
	}
}

func TestAuthorizationStatus_invalidValueFails(t *testing.T) {
	t.Parallel()

	_, err := AuthorizationStatus("UnknownStatus")
	if err == nil {
		t.Error("expected error for invalid status, got nil")
	} else if !errors.Is(err, ErrInvalidAuthorizationStatus) {
		t.Errorf("expected ErrInvalidAuthorizationStatus, got: %v", err)
	}
}

func TestAuthorizationStatus_emptyFails(t *testing.T) {
	t.Parallel()

	_, err := AuthorizationStatus("")
	if err == nil {
		t.Error("expected error for empty status, got nil")
	} else if !errors.Is(err, ErrInvalidAuthorizationStatus) {
		t.Errorf("expected ErrInvalidAuthorizationStatus, got: %v", err)
	}
}

func TestAuthorizationStatus_nullStringFails(t *testing.T) {
	t.Parallel()

	_, err := AuthorizationStatus("null")
	if err == nil {
		t.Error("expected error for 'null' string, got nil")
	} else if !errors.Is(err, ErrInvalidAuthorizationStatus) {
		t.Errorf("expected ErrInvalidAuthorizationStatus, got: %v", err)
	}
}

func TestAuthorizationStatus_validateValid(t *testing.T) {
	t.Parallel()

	s, _ := AuthorizationStatus("Blocked")
	if err := s.Validate(); err != nil {
		t.Errorf(errExpectedValidateNil, err)
	}
}

func TestAuthorizationStatus_validateFailsForGarbage(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatusType{value: "Invalid"}
	err := status.Validate()

	if err == nil {
		t.Error(errExpectedValidateFail)
	} else if !errors.Is(err, ErrInvalidAuthorizationStatus) {
		t.Errorf("expected ErrInvalidAuthorizationStatus, got: %v", err)
	}
}

func TestAuthorizationStatus_constantsMatch(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"Accepted":     Accepted,
		"Blocked":      Blocked,
		"Expired":      Expired,
		"Invalid":      Invalid,
		"ConcurrentTx": ConcurrentTx,
	}

	for expected, actual := range tests {
		if expected != actual {
			t.Errorf("expected constant %q to match value %q", expected, actual)
		}
	}
}
