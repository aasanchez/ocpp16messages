package authorizetypes

import (
	"testing"
)

func TestAuthorizationStatus_validAccepted(t *testing.T) {
	t.Parallel()

	status, err := AuthorizationStatus(Accepted)
	if err != nil {
		t.Fatalf("validAccepted: unexpected error on: %v", err)
	}

	if status.Value() != Accepted {
		t.Errorf("validAccepted: expected value %q, got %q", Accepted, status.Value())
	}
}

func TestAuthorizationStatus_validBlocked(t *testing.T) {
	t.Parallel()

	status, err := AuthorizationStatus(Blocked)
	if err != nil {
		t.Fatalf("validBlocked: unexpected error: %v", err)
	}

	if status.Value() != Blocked {
		t.Errorf("validBlocked: expected value %q, got %q", Blocked, status.Value())
	}
}

func TestAuthorizationStatus_validExpired(t *testing.T) {
	t.Parallel()

	status, err := AuthorizationStatus(Expired)
	if err != nil {
		t.Fatalf("validExpired: unexpected error: %v", err)
	}

	if status.Value() != Expired {
		t.Errorf("validExpired: expected value %q, got %q", Expired, status.Value())
	}
}

func TestAuthorizationStatus_validInvalid(t *testing.T) {
	t.Parallel()

	status, err := AuthorizationStatus(Invalid)
	if err != nil {
		t.Fatalf("validInvalid: unexpected error: %v", err)
	}

	if status.Value() != Invalid {
		t.Errorf("validInvalid: expected value %q, got %q", Invalid, status.Value())
	}
}

func TestAuthorizationStatus_validConcurrentTx(t *testing.T) {
	t.Parallel()

	status, err := AuthorizationStatus(ConcurrentTx)
	if err != nil {
		t.Fatalf("validConcurrentTx: unexpected error: %v", err)
	}

	if status.Value() != ConcurrentTx {
		t.Errorf("validConcurrentTx: expected value %q, got %q", ConcurrentTx, status.Value())
	}
}

func TestAuthorizationStatus_invalidStatus(t *testing.T) {
	t.Parallel()

	_, err := AuthorizationStatus("SomethingElse")
	if err == nil {
		t.Fatal("invalidStatus: expected error for invalid status, got nil")
	}
}
