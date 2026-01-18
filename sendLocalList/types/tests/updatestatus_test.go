package types_test

import (
	"testing"

	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr        = "Accepted"
	statusFailedStr          = "Failed"
	statusNotSupportedStr    = "NotSupported"
	statusVersionMismatchStr = "VersionMismatch"
	methodUpdateStatusString = "UpdateStatus.String()"
)

func TestUpdateStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !slt.UpdateStatusAccepted.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UpdateStatusAccepted",
		)
	}
}

func TestUpdateStatus_IsValid_Failed(t *testing.T) {
	t.Parallel()

	if !slt.UpdateStatusFailed.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UpdateStatusFailed",
		)
	}
}

func TestUpdateStatus_IsValid_NotSupported(t *testing.T) {
	t.Parallel()

	if !slt.UpdateStatusNotSupported.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UpdateStatusNotSupported",
		)
	}
}

func TestUpdateStatus_IsValid_VersionMismatch(t *testing.T) {
	t.Parallel()

	if !slt.UpdateStatusVersionMismatch.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"UpdateStatusVersionMismatch",
		)
	}
}

func TestUpdateStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	updateStatus := slt.UpdateStatus("")
	if updateStatus.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UpdateStatus(\"\")",
		)
	}
}

func TestUpdateStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	updateStatus := slt.UpdateStatus("Unknown")
	if updateStatus.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UpdateStatus(\"Unknown\")",
		)
	}
}

func TestUpdateStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	updateStatus := slt.UpdateStatus("accepted")
	if updateStatus.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"UpdateStatus(\"accepted\")",
		)
	}
}

func TestUpdateStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := slt.UpdateStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodUpdateStatusString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestUpdateStatus_String_Failed(t *testing.T) {
	t.Parallel()

	got := slt.UpdateStatusFailed.String()
	if got != statusFailedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodUpdateStatusString,
			got,
			statusFailedStr,
		)
	}
}

func TestUpdateStatus_String_NotSupported(t *testing.T) {
	t.Parallel()

	got := slt.UpdateStatusNotSupported.String()
	if got != statusNotSupportedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodUpdateStatusString,
			got,
			statusNotSupportedStr,
		)
	}
}

func TestUpdateStatus_String_VersionMismatch(t *testing.T) {
	t.Parallel()

	got := slt.UpdateStatusVersionMismatch.String()
	if got != statusVersionMismatchStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodUpdateStatusString,
			got,
			statusVersionMismatchStr,
		)
	}
}
