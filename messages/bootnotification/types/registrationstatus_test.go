package bootnotificationtypes

import (
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestRegistrationStatus_Accepted(t *testing.T) {
	t.Parallel()

	registrationStatus, err := RegistrationStatus("Accepted")
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}

	if registrationStatus.Value() != RegistrationAccepted {
		t.Errorf(sharedtypes.ErrExpectedValueMismatch, RegistrationAccepted, registrationStatus.Value())
	}
}

func TestRegistrationStatus_Pending(t *testing.T) {
	t.Parallel()

	registrationStatus, err := RegistrationStatus("Pending")
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}

	if registrationStatus.Value() != RegistrationPending {
		t.Errorf(sharedtypes.ErrExpectedValueMismatch, RegistrationPending, registrationStatus.Value())
	}
}

func TestRegistrationStatus_Rejected(t *testing.T) {
	t.Parallel()

	registrationStatus, err := RegistrationStatus("Rejected")
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}

	if registrationStatus.Value() != RegistrationRejected {
		t.Errorf(sharedtypes.ErrExpectedValueMismatch, RegistrationRejected, registrationStatus.Value())
	}
}

func TestRegistrationStatus_Empty_IsInvalid(t *testing.T) {
	t.Parallel()

	_, err := RegistrationStatus("")
	if err == nil {
		t.Error("expected error for empty string, got nil")
	}
}

func TestRegistrationStatus_LowercaseAccepted_IsInvalid(t *testing.T) {
	t.Parallel()

	_, err := RegistrationStatus("accepted")
	if err == nil {
		t.Error("expected error for 'accepted', got nil")
	}
}

func TestRegistrationStatus_Approved_IsInvalid(t *testing.T) {
	t.Parallel()

	_, err := RegistrationStatus("approved")
	if err == nil {
		t.Error("expected error for 'approved', got nil")
	}
}

func TestRegistrationStatus_Declined_IsInvalid(t *testing.T) {
	t.Parallel()

	_, err := RegistrationStatus("Declined")
	if err == nil {
		t.Error("expected error for 'Declined', got nil")
	}
}

func TestRegistrationStatus_Unknown_IsInvalid(t *testing.T) {
	t.Parallel()

	_, err := RegistrationStatus("unknown")
	if err == nil {
		t.Error("expected error for 'unknown', got nil")
	}
}
