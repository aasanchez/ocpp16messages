package types_test

import (
	"testing"

	mbt "github.com/aasanchez/ocpp16messages/messages/bootNotification/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusPendingStr  = "Pending"
	statusRejectedStr = "Rejected"
)

func TestRegistrationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !mbt.RegistrationStatusAccepted.IsValid() {
		t.Error("RegistrationStatusAccepted.IsValid() = false, want true")
	}
}

func TestRegistrationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	if !mbt.RegistrationStatusPending.IsValid() {
		t.Error("RegistrationStatusPending.IsValid() = false, want true")
	}
}

func TestRegistrationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !mbt.RegistrationStatusRejected.IsValid() {
		t.Error("RegistrationStatusRejected.IsValid() = false, want true")
	}
}

func TestRegistrationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := mbt.RegistrationStatus("")
	if status.IsValid() {
		t.Error("RegistrationStatus(\"\").IsValid() = true, want false")
	}
}

func TestRegistrationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := mbt.RegistrationStatus("Unknown")
	if status.IsValid() {
		t.Error("RegistrationStatus(\"Unknown\").IsValid() = true, want false")
	}
}

func TestRegistrationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := mbt.RegistrationStatus("accepted")
	if status.IsValid() {
		t.Error("RegistrationStatus(\"accepted\").IsValid() = true, want false")
	}
}

func TestRegistrationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := mbt.RegistrationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(ErrorMethodMismatch, got, statusAcceptedStr)
	}
}

func TestRegistrationStatus_String_Pending(t *testing.T) {
	t.Parallel()

	got := mbt.RegistrationStatusPending.String()
	if got != statusPendingStr {
		t.Errorf(ErrorMethodMismatch, got, statusPendingStr)
	}
}

func TestRegistrationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := mbt.RegistrationStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(ErrorMethodMismatch, got, statusRejectedStr)
	}
}
