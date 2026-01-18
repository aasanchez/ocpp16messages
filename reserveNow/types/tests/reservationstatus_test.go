package types_test

import (
	"testing"

	rt "github.com/aasanchez/ocpp16messages/reserveNow/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr    = "Accepted"
	statusFaultedStr     = "Faulted"
	statusOccupiedStr    = "Occupied"
	statusRejectedStr    = "Rejected"
	statusUnavailableStr = "Unavailable"
)

func TestReservationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !rt.ReservationStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReservationStatusAccepted")
	}
}

func TestReservationStatus_IsValid_Faulted(t *testing.T) {
	t.Parallel()

	if !rt.ReservationStatusFaulted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReservationStatusFaulted")
	}
}

func TestReservationStatus_IsValid_Occupied(t *testing.T) {
	t.Parallel()

	if !rt.ReservationStatusOccupied.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReservationStatusOccupied")
	}
}

func TestReservationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !rt.ReservationStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReservationStatusRejected")
	}
}

func TestReservationStatus_IsValid_Unavailable(t *testing.T) {
	t.Parallel()

	if !rt.ReservationStatusUnavailable.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ReservationStatusUnavailable")
	}
}

func TestReservationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := rt.ReservationStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReservationStatus(\"\")")
	}
}

func TestReservationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := rt.ReservationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReservationStatus(\"Unknown\")")
	}
}

func TestReservationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := rt.ReservationStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReservationStatus(\"accepted\")")
	}
}

func TestReservationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := rt.ReservationStatus("Pending")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ReservationStatus(\"Pending\")")
	}
}

func TestReservationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := rt.ReservationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ReservationStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestReservationStatus_String_Faulted(t *testing.T) {
	t.Parallel()

	got := rt.ReservationStatusFaulted.String()
	if got != statusFaultedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ReservationStatus.String()",
			got,
			statusFaultedStr,
		)
	}
}

func TestReservationStatus_String_Occupied(t *testing.T) {
	t.Parallel()

	got := rt.ReservationStatusOccupied.String()
	if got != statusOccupiedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ReservationStatus.String()",
			got,
			statusOccupiedStr,
		)
	}
}

func TestReservationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := rt.ReservationStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ReservationStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}

func TestReservationStatus_String_Unavailable(t *testing.T) {
	t.Parallel()

	got := rt.ReservationStatusUnavailable.String()
	if got != statusUnavailableStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ReservationStatus.String()",
			got,
			statusUnavailableStr,
		)
	}
}
