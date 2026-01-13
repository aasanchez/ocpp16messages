package types_test

import (
	"testing"

	ct "github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusUnknownStr  = "Unknown"
)

func TestClearChargingProfileStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !ct.ClearChargingProfileStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ClearChargingProfileStatusAccepted")
	}
}

func TestClearChargingProfileStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	if !ct.ClearChargingProfileStatusUnknown.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ClearChargingProfileStatusUnknown")
	}
}

func TestClearChargingProfileStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := ct.ClearChargingProfileStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ClearChargingProfileStatus(\"\")")
	}
}

func TestClearChargingProfileStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	status := ct.ClearChargingProfileStatus("Invalid")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ClearChargingProfileStatus(\"Invalid\")")
	}
}

func TestClearChargingProfileStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := ct.ClearChargingProfileStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ClearChargingProfileStatus(\"accepted\")",
		)
	}
}

func TestClearChargingProfileStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	status := ct.ClearChargingProfileStatus("Rejected")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ClearChargingProfileStatus(\"Rejected\")",
		)
	}
}

func TestClearChargingProfileStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := ct.ClearChargingProfileStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ClearChargingProfileStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestClearChargingProfileStatus_String_Unknown(t *testing.T) {
	t.Parallel()

	got := ct.ClearChargingProfileStatusUnknown.String()
	if got != statusUnknownStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ClearChargingProfileStatus.String()",
			got,
			statusUnknownStr,
		)
	}
}
