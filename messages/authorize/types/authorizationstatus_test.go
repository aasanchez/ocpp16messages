package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestAuthorizationStatus_validAccepted(t *testing.T) {
	t.Parallel()

	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)
	if err != nil {
		t.Fatalf("validAccepted: unexpected error on: %v", err)
	}

	if status.Value() != authorizetypes.Accepted {
		t.Errorf("validAccepted: expected value %q, got %q", authorizetypes.Accepted, status.Value())
	}
}

func TestAuthorizationStatus_validBlocked(t *testing.T) {
	t.Parallel()

	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Blocked)
	if err != nil {
		t.Fatalf("validBlocked: unexpected error: %v", err)
	}

	if status.Value() != authorizetypes.Blocked {
		t.Errorf("validBlocked: expected value %q, got %q", authorizetypes.Blocked, status.Value())
	}
}

func TestAuthorizationStatus_validExpired(t *testing.T) {
	t.Parallel()

	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Expired)
	if err != nil {
		t.Fatalf("validExpired: unexpected error: %v", err)
	}

	if status.Value() != authorizetypes.Expired {
		t.Errorf("validExpired: expected value %q, got %q", authorizetypes.Expired, status.Value())
	}
}

func TestAuthorizationStatus_validInvalid(t *testing.T) {
	t.Parallel()

	status, err := authorizetypes.AuthorizationStatus(authorizetypes.Invalid)
	if err != nil {
		t.Fatalf("validInvalid: unexpected error: %v", err)
	}

	if status.Value() != authorizetypes.Invalid {
		t.Errorf("validInvalid: expected value %q, got %q", authorizetypes.Invalid, status.Value())
	}
}

func TestAuthorizationStatus_validConcurrentTx(t *testing.T) {
	t.Parallel()

	status, err := authorizetypes.AuthorizationStatus(authorizetypes.ConcurrentTx)
	if err != nil {
		t.Fatalf("validConcurrentTx: unexpected error: %v", err)
	}

	if status.Value() != authorizetypes.ConcurrentTx {
		t.Errorf("validConcurrentTx: expected value %q, got %q", authorizetypes.ConcurrentTx, status.Value())
	}
}

func TestAuthorizationStatus_invalidStatus(t *testing.T) {
	t.Parallel()

	_, err := authorizetypes.AuthorizationStatus("SomethingElse")
	if err == nil {
		t.Fatal("invalidStatus: expected error for invalid status, got nil")
	}
}
