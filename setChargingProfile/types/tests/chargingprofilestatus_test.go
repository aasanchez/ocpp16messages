//nolint:dupl // enum test pattern is intentionally similar across types
package types_test

import (
	"testing"

	spt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	"github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr     = "Accepted"
	statusRejectedStr     = "Rejected"
	statusNotSupportedStr = "NotSupported"
	methodStatusString    = "ChargingProfileStatus.String()"
)

func TestChargingProfileStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !spt.ChargingProfileStatusAccepted.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"ChargingProfileStatusAccepted",
		)
	}
}

func TestChargingProfileStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !spt.ChargingProfileStatusRejected.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"ChargingProfileStatusRejected",
		)
	}
}

func TestChargingProfileStatus_IsValid_NotSupported(t *testing.T) {
	t.Parallel()

	if !spt.ChargingProfileStatusNotSupported.IsValid() {
		t.Errorf(
			types.ErrorIsValidFalse,
			"ChargingProfileStatusNotSupported",
		)
	}
}

func TestChargingProfileStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := spt.ChargingProfileStatus("")
	if status.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"ChargingProfileStatus(\"\")",
		)
	}
}

func TestChargingProfileStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := spt.ChargingProfileStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"ChargingProfileStatus(\"Unknown\")",
		)
	}
}

func TestChargingProfileStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := spt.ChargingProfileStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			types.ErrorIsValidTrue,
			"ChargingProfileStatus(\"accepted\")",
		)
	}
}

func TestChargingProfileStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := spt.ChargingProfileStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodStatusString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestChargingProfileStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := spt.ChargingProfileStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodStatusString,
			got,
			statusRejectedStr,
		)
	}
}

func TestChargingProfileStatus_String_NotSupported(t *testing.T) {
	t.Parallel()

	got := spt.ChargingProfileStatusNotSupported.String()
	if got != statusNotSupportedStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodStatusString,
			got,
			statusNotSupportedStr,
		)
	}
}
