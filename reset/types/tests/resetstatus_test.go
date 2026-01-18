package types_test

import (
	"testing"

	rt "github.com/aasanchez/ocpp16messages/reset/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusRejectedStr = "Rejected"
)

func TestResetStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !rt.ResetStatusAccepted.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"ResetStatusAccepted",
		)
	}
}

func TestResetStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !rt.ResetStatusRejected.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"ResetStatusRejected",
		)
	}
}

func TestResetStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := rt.ResetStatus("")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetStatus(\"\")",
		)
	}
}

func TestResetStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := rt.ResetStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetStatus(\"Unknown\")",
		)
	}
}

func TestResetStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := rt.ResetStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetStatus(\"accepted\")",
		)
	}
}

func TestResetStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := rt.ResetStatus("Pending")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"ResetStatus(\"Pending\")",
		)
	}
}

func TestResetStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := rt.ResetStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ResetStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestResetStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := rt.ResetStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"ResetStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}
