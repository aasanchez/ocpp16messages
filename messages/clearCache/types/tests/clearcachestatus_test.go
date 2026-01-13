package types_test

import (
	"testing"

	mct "github.com/aasanchez/ocpp16messages/messages/clearCache/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusRejectedStr = "Rejected"
)

func TestClearCacheStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !mct.ClearCacheStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ClearCacheStatusAccepted")
	}
}

func TestClearCacheStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !mct.ClearCacheStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "ClearCacheStatusRejected")
	}
}

func TestClearCacheStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := mct.ClearCacheStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ClearCacheStatus(\"\")")
	}
}

func TestClearCacheStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := mct.ClearCacheStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ClearCacheStatus(\"Unknown\")")
	}
}

func TestClearCacheStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := mct.ClearCacheStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ClearCacheStatus(\"accepted\")")
	}
}

func TestClearCacheStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := mct.ClearCacheStatus("Pending")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "ClearCacheStatus(\"Pending\")")
	}
}

func TestClearCacheStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := mct.ClearCacheStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ClearCacheStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestClearCacheStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := mct.ClearCacheStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ClearCacheStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}
