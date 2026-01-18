package types_test

import (
	"testing"

	rt "github.com/aasanchez/ocpp16messages/remoteStopTransaction/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusRejectedStr = "Rejected"
)

func TestRemoteStopTransactionStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !rt.RemoteStopTransactionStatusAccepted.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"RemoteStopTransactionStatusAccepted",
		)
	}
}

func TestRemoteStopTransactionStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !rt.RemoteStopTransactionStatusRejected.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"RemoteStopTransactionStatusRejected",
		)
	}
}

func TestRemoteStopTransactionStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStopTransactionStatus("")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStopTransactionStatus(\"\")",
		)
	}
}

func TestRemoteStopTransactionStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStopTransactionStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStopTransactionStatus(\"Unknown\")",
		)
	}
}

func TestRemoteStopTransactionStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStopTransactionStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStopTransactionStatus(\"accepted\")",
		)
	}
}

func TestRemoteStopTransactionStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStopTransactionStatus("Pending")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStopTransactionStatus(\"Pending\")",
		)
	}
}

func TestRemoteStopTransactionStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := rt.RemoteStopTransactionStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"RemoteStopTransactionStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestRemoteStopTransactionStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := rt.RemoteStopTransactionStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"RemoteStopTransactionStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}
