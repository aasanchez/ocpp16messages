package types_test

import (
	"testing"

	dt "github.com/aasanchez/ocpp16messages/dataTransfer/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr         = "Accepted"
	statusRejectedStr         = "Rejected"
	statusUnknownMessageIdStr = "UnknownMessageId"
	statusUnknownVendorStr    = "UnknownVendor"
)

func TestDataTransferStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !dt.DataTransferStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DataTransferStatusAccepted")
	}
}

func TestDataTransferStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !dt.DataTransferStatusRejected.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DataTransferStatusRejected")
	}
}

func TestDataTransferStatus_IsValid_UnknownMessageId(t *testing.T) {
	t.Parallel()

	if !dt.DataTransferStatusUnknownMessageId.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DataTransferStatusUnknownMessageId")
	}
}

func TestDataTransferStatus_IsValid_UnknownVendor(t *testing.T) {
	t.Parallel()

	if !dt.DataTransferStatusUnknownVendor.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "DataTransferStatusUnknownVendor")
	}
}

func TestDataTransferStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := dt.DataTransferStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "DataTransferStatus(\"\")")
	}
}

func TestDataTransferStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	status := dt.DataTransferStatus("Invalid")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "DataTransferStatus(\"Invalid\")")
	}
}

func TestDataTransferStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := dt.DataTransferStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "DataTransferStatus(\"accepted\")")
	}
}

func TestDataTransferStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := dt.DataTransferStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"DataTransferStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestDataTransferStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := dt.DataTransferStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"DataTransferStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}

func TestDataTransferStatus_String_UnknownMessageId(t *testing.T) {
	t.Parallel()

	got := dt.DataTransferStatusUnknownMessageId.String()
	if got != statusUnknownMessageIdStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"DataTransferStatus.String()",
			got,
			statusUnknownMessageIdStr,
		)
	}
}

func TestDataTransferStatus_String_UnknownVendor(t *testing.T) {
	t.Parallel()

	got := dt.DataTransferStatusUnknownVendor.String()
	if got != statusUnknownVendorStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"DataTransferStatus.String()",
			got,
			statusUnknownVendorStr,
		)
	}
}
