package types_test

import (
	"testing"

	mct "github.com/aasanchez/ocpp16messages/cancelReservation/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusRejectedStr = "Rejected"
)

func TestCancelReservationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !mct.CancelReservationStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "CancelReservationStatusAccepted")
	}
}

func TestCancelReservationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !mct.CancelReservationStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "CancelReservationStatusRejected")
	}
}

func TestCancelReservationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "CancelReservationStatus(\"\")")
	}
}

func TestCancelReservationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "CancelReservationStatus(\"Unknown\")")
	}
}

func TestCancelReservationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "CancelReservationStatus(\"accepted\")")
	}
}

func TestCancelReservationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := mct.CancelReservationStatus("Pending")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "CancelReservationStatus(\"Pending\")")
	}
}

func TestCancelReservationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := mct.CancelReservationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"CancelReservationStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestCancelReservationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := mct.CancelReservationStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"CancelReservationStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}
