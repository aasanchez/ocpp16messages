package types_test

import (
	"testing"

	mbt "github.com/aasanchez/ocpp16messages/bootNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusPendingStr  = "Pending"
	statusRejectedStr = "Rejected"
	methodString      = "RegistrationStatus.String()"
)

func TestRegistrationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !mbt.RegistrationStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RegistrationStatusAccepted")
	}
}

func TestRegistrationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	if !mbt.RegistrationStatusPending.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RegistrationStatusPending")
	}
}

func TestRegistrationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !mbt.RegistrationStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "RegistrationStatusRejected")
	}
}

func TestRegistrationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := mbt.RegistrationStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "RegistrationStatus(\"\")")
	}
}

func TestRegistrationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := mbt.RegistrationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "RegistrationStatus(\"Unknown\")")
	}
}

func TestRegistrationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := mbt.RegistrationStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "RegistrationStatus(\"accepted\")")
	}
}

func TestRegistrationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := mbt.RegistrationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestRegistrationStatus_String_Pending(t *testing.T) {
	t.Parallel()

	got := mbt.RegistrationStatusPending.String()
	if got != statusPendingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusPendingStr,
		)
	}
}

func TestRegistrationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := mbt.RegistrationStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusRejectedStr,
		)
	}
}
