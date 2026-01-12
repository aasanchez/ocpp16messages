package types_test

import (
	"testing"

	mct "github.com/aasanchez/ocpp16messages/messages/cancelReservation/types"
)

const (
	statusAcceptedStr   = "Accepted"
	statusRejectedStr   = "Rejected"
	errIsValidUnknown   = "Status(\"Unknown\").IsValid() = true, want false"
	errIsValidLowercase = "Status(\"accepted\").IsValid() = true, want false"
	errIsValidPending   = "Status(\"Pending\").IsValid() = true, want false"
	errMethodMismatch   = "CancelReservationStatus.String() = %v, want %v"
)

func TestCancelReservationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !mct.CancelReservationStatusAccepted.IsValid() {
		t.Error("CancelReservationStatusAccepted.IsValid() = false, want true")
	}
}

func TestCancelReservationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !mct.CancelReservationStatusRejected.IsValid() {
		t.Error("CancelReservationStatusRejected.IsValid() = false, want true")
	}
}

func TestCancelReservationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("")
	if status.IsValid() {
		t.Error("CancelReservationStatus(\"\").IsValid() = true, want false")
	}
}

func TestCancelReservationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("Unknown")
	if status.IsValid() {
		t.Error(errIsValidUnknown)
	}
}

func TestCancelReservationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("accepted")
	if status.IsValid() {
		t.Error(errIsValidLowercase)
	}
}

func TestCancelReservationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("Pending")
	if status.IsValid() {
		t.Error(errIsValidPending)
	}
}

func TestCancelReservationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := mct.CancelReservationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(errMethodMismatch, got, statusAcceptedStr)
	}
}

func TestCancelReservationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := mct.CancelReservationStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(errMethodMismatch, got, statusRejectedStr)
	}
}
