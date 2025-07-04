package authorizetypes_test

import (
	"testing"

	at "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestAuthorizationStatus_validAccepted(t *testing.T) {
	t.Parallel()

	status, err := at.SetAuthorizationStatus("Accepted")
	if err != nil {
		t.Fatalf("validAccepted: unexpected error: %v", err)
	}
	if status.Value() != at.Accepted {
		t.Errorf("validAccepted: expected value %q, got %q", at.Accepted, status.Value())
	}
}

func TestAuthorizationStatus_validBlocked(t *testing.T) {
	t.Parallel()

	status, err := at.SetAuthorizationStatus("Blocked")
	if err != nil {
		t.Fatalf("validBlocked: unexpected error: %v", err)
	}
	if status.Value() != at.Blocked {
		t.Errorf("validBlocked: expected value %q, got %q", at.Blocked, status.Value())
	}
}

func TestAuthorizationStatus_validExpired(t *testing.T) {
	t.Parallel()

	status, err := at.SetAuthorizationStatus("Expired")
	if err != nil {
		t.Fatalf("validExpired: unexpected error: %v", err)
	}
	if status.Value() != at.Expired {
		t.Errorf("validExpired: expected value %q, got %q", at.Expired, status.Value())
	}
}

func TestAuthorizationStatus_validInvalid(t *testing.T) {
	t.Parallel()

	status, err := at.SetAuthorizationStatus("Invalid")
	if err != nil {
		t.Fatalf("validInvalid: unexpected error: %v", err)
	}
	if status.Value() != at.Invalid {
		t.Errorf("validInvalid: expected value %q, got %q", at.Invalid, status.Value())
	}
}

func TestAuthorizationStatus_validConcurrentTx(t *testing.T) {
	t.Parallel()

	status, err := at.SetAuthorizationStatus("ConcurrentTx")
	if err != nil {
		t.Fatalf("validConcurrentTx: unexpected error: %v", err)
	}
	if status.Value() != at.ConcurrentTx {
		t.Errorf("validConcurrentTx: expected value %q, got %q", at.ConcurrentTx, status.Value())
	}
}

func TestAuthorizationStatus_invalidStatus(t *testing.T) {
	t.Parallel()

	_, err := at.SetAuthorizationStatus("SomethingElse")
	if err == nil {
		t.Fatal("invalidStatus: expected error for invalid status, got nil")
	}
}
