package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
)

const (
	errStatusMismatch = "Expected %v, got %v"
	statusAccepted    = "Accepted"
	statusRejected    = "Rejected"
	statusScheduled   = "Scheduled"
)

func TestAvailabilityStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus(statusAccepted)
	if !as.IsValid() {
		t.Errorf(errStatusMismatch, true, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus(statusRejected)
	if !as.IsValid() {
		t.Errorf(errStatusMismatch, true, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Scheduled(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus(statusScheduled)
	if !as.IsValid() {
		t.Errorf(errStatusMismatch, true, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("")
	if as.IsValid() {
		t.Errorf(errStatusMismatch, false, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("Unknown")
	if as.IsValid() {
		t.Errorf(errStatusMismatch, false, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("accepted")
	if as.IsValid() {
		t.Errorf(errStatusMismatch, false, as.IsValid())
	}
}

func TestAvailabilityStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatus("Pending")
	if as.IsValid() {
		t.Errorf(errStatusMismatch, false, as.IsValid())
	}
}

func TestAvailabilityStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatusAccepted
	if as.String() != statusAccepted {
		t.Errorf(errStatusMismatch, statusAccepted, as.String())
	}
}

func TestAvailabilityStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatusRejected
	if as.String() != statusRejected {
		t.Errorf(errStatusMismatch, statusRejected, as.String())
	}
}

func TestAvailabilityStatus_String_Scheduled(t *testing.T) {
	t.Parallel()

	as := types.AvailabilityStatusScheduled
	if as.String() != statusScheduled {
		t.Errorf(errStatusMismatch, statusScheduled, as.String())
	}
}

func TestAvailabilityStatus_Constants(t *testing.T) {
	t.Parallel()

	if types.AvailabilityStatusAccepted != statusAccepted {
		t.Errorf(
			errStatusMismatch,
			statusAccepted,
			types.AvailabilityStatusAccepted,
		)
	}

	if types.AvailabilityStatusRejected != statusRejected {
		t.Errorf(
			errStatusMismatch,
			statusRejected,
			types.AvailabilityStatusRejected,
		)
	}

	if types.AvailabilityStatusScheduled != statusScheduled {
		t.Errorf(
			errStatusMismatch,
			statusScheduled,
			types.AvailabilityStatusScheduled,
		)
	}
}
