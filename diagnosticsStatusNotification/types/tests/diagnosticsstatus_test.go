package types_test

import (
	"testing"

	dn "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusIdleStr         = "Idle"
	statusUploadedStr     = "Uploaded"
	statusUploadFailedStr = "UploadFailed"
	statusUploadingStr    = "Uploading"
	methodStringName      = "methodStringName"
)

func TestDiagnosticsStatus_IsValid_Idle(t *testing.T) {
	t.Parallel()

	if !dn.DiagnosticsStatusIdle.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DiagnosticsStatusIdle")
	}
}

func TestDiagnosticsStatus_IsValid_Uploaded(t *testing.T) {
	t.Parallel()

	if !dn.DiagnosticsStatusUploaded.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DiagnosticsStatusUploaded")
	}
}

func TestDiagnosticsStatus_IsValid_UploadFailed(t *testing.T) {
	t.Parallel()

	if !dn.DiagnosticsStatusUploadFailed.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DiagnosticsStatusUploadFailed")
	}
}

func TestDiagnosticsStatus_IsValid_Uploading(t *testing.T) {
	t.Parallel()

	if !dn.DiagnosticsStatusUploading.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DiagnosticsStatusUploading")
	}
}

func TestDiagnosticsStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := dn.DiagnosticsStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "DiagnosticsStatus(\"\")")
	}
}

func TestDiagnosticsStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	status := dn.DiagnosticsStatus("Invalid")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "DiagnosticsStatus(\"Invalid\")")
	}
}

func TestDiagnosticsStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := dn.DiagnosticsStatus("idle")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "DiagnosticsStatus(\"idle\")")
	}
}

func TestDiagnosticsStatus_String_Idle(t *testing.T) {
	t.Parallel()

	got := dn.DiagnosticsStatusIdle.String()
	if got != statusIdleStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusIdleStr,
		)
	}
}

func TestDiagnosticsStatus_String_Uploaded(t *testing.T) {
	t.Parallel()

	got := dn.DiagnosticsStatusUploaded.String()
	if got != statusUploadedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusUploadedStr,
		)
	}
}

func TestDiagnosticsStatus_String_UploadFailed(t *testing.T) {
	t.Parallel()

	got := dn.DiagnosticsStatusUploadFailed.String()
	if got != statusUploadFailedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusUploadFailedStr,
		)
	}
}

func TestDiagnosticsStatus_String_Uploading(t *testing.T) {
	t.Parallel()

	got := dn.DiagnosticsStatusUploading.String()
	if got != statusUploadingStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringName,
			got,
			statusUploadingStr,
		)
	}
}
