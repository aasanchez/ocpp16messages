package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	statusAccepted  = "Accepted"
	statusRejected  = "Rejected"
	statusScheduled = "Scheduled"
)

func TestAvailabilityStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus(statusAccepted)
	if !as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus(statusRejected)
	if !as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Scheduled(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus(statusScheduled)
	if !as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, true, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("")
	if as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("Unknown")
	if as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("accepted")
	if as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("Pending")
	if as.IsValid() {
		t.Errorf(st.ErrorMismatchValue, false, as.IsValid())
	}
}

func TestAvailabilityStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatusAccepted
	if as.String() != statusAccepted {
		t.Errorf(st.ErrorMismatchValue, statusAccepted, as.String())
	}
}

func TestAvailabilityStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatusRejected
	if as.String() != statusRejected {
		t.Errorf(st.ErrorMismatchValue, statusRejected, as.String())
	}
}

func TestAvailabilityStatus_String_Scheduled(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatusScheduled
	if as.String() != statusScheduled {
		t.Errorf(st.ErrorMismatchValue, statusScheduled, as.String())
	}
}

func TestAvailabilityStatus_Constants(t *testing.T) {
	t.Parallel()

	if types.AvailabilityStatusAccepted != statusAccepted {
		t.Errorf(
			st.ErrorMismatchValue,
			statusAccepted,
			types.AvailabilityStatusAccepted,
		)
	}

	if types.AvailabilityStatusRejected != statusRejected {
		t.Errorf(
			st.ErrorMismatchValue,
			statusRejected,
			types.AvailabilityStatusRejected,
		)
	}

	if types.AvailabilityStatusScheduled != statusScheduled {
		t.Errorf(
			st.ErrorMismatchValue,
			statusScheduled,
			types.AvailabilityStatusScheduled,
		)
	}
}
