package types_test

import (
	"testing"

	tt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	acceptedStr        = "Accepted"
	rejectedStr        = "Rejected"
	notImplementedStr  = "NotImplemented"
	methodStringStatus = "TriggerMessageStatus.String()"
)

func TestTriggerMessageStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !tt.TriggerMessageStatusAccepted.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"TriggerMessageStatusAccepted",
		)
	}
}

func TestTriggerMessageStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !tt.TriggerMessageStatusRejected.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"TriggerMessageStatusRejected",
		)
	}
}

func TestTriggerMessageStatus_IsValid_NotImplemented(t *testing.T) {
	t.Parallel()

	if !tt.TriggerMessageStatusNotImplemented.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"TriggerMessageStatusNotImplemented",
		)
	}
}

func TestTriggerMessageStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := tt.TriggerMessageStatus("")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"TriggerMessageStatus(\"\")",
		)
	}
}

func TestTriggerMessageStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := tt.TriggerMessageStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"TriggerMessageStatus(\"Unknown\")",
		)
	}
}

func TestTriggerMessageStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := tt.TriggerMessageStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"TriggerMessageStatus(\"accepted\")",
		)
	}
}

func TestTriggerMessageStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := tt.TriggerMessageStatusAccepted.String()
	if got != acceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringStatus,
			got,
			acceptedStr,
		)
	}
}

func TestTriggerMessageStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := tt.TriggerMessageStatusRejected.String()
	if got != rejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringStatus,
			got,
			rejectedStr,
		)
	}
}

func TestTriggerMessageStatus_String_NotImplemented(t *testing.T) {
	t.Parallel()

	got := tt.TriggerMessageStatusNotImplemented.String()
	if got != notImplementedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodStringStatus,
			got,
			notImplementedStr,
		)
	}
}
