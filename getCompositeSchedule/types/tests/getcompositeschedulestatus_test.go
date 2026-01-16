package types_test

import (
	"testing"

	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusRejectedStr = "Rejected"
	statusTypeString  = "GetCompositeScheduleStatus.String()"
)

func TestGetCompositeScheduleStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !gt.GetCompositeScheduleStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "GetCompositeScheduleStatusAccepted")
	}
}

func TestGetCompositeScheduleStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !gt.GetCompositeScheduleStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "GetCompositeScheduleStatusRejected")
	}
}

func TestGetCompositeScheduleStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := gt.GetCompositeScheduleStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "GetCompositeScheduleStatus(\"\")")
	}
}

func TestGetCompositeScheduleStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	status := gt.GetCompositeScheduleStatus("Invalid")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "GetCompositeScheduleStatus(\"Invalid\")")
	}
}

func TestGetCompositeScheduleStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := gt.GetCompositeScheduleStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"GetCompositeScheduleStatus(\"accepted\")",
		)
	}
}

func TestGetCompositeScheduleStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := gt.GetCompositeScheduleStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(st.ErrorMethodMismatch, statusTypeString, got, statusAcceptedStr)
	}
}

func TestGetCompositeScheduleStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := gt.GetCompositeScheduleStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(st.ErrorMethodMismatch, statusTypeString, got, statusRejectedStr)
	}
}
