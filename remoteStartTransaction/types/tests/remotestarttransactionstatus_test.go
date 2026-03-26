package types_test

import (
	"testing"

	rt "github.com/aasanchez/ocpp16messages/remotestarttransaction/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr = "Accepted"
	statusRejectedStr = "Rejected"
)

func TestRemoteStartTransactionStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !rt.RemoteStartTransactionStatusAccepted.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"RemoteStartTransactionStatusAccepted",
		)
	}
}

func TestRemoteStartTransactionStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !rt.RemoteStartTransactionStatusRejected.IsValid() {
		t.Errorf(
			st.ErrorIsValidFalse,
			"RemoteStartTransactionStatusRejected",
		)
	}
}

func TestRemoteStartTransactionStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStartTransactionStatus("")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStartTransactionStatus(\"\")",
		)
	}
}

func TestRemoteStartTransactionStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStartTransactionStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStartTransactionStatus(\"Unknown\")",
		)
	}
}

func TestRemoteStartTransactionStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStartTransactionStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStartTransactionStatus(\"accepted\")",
		)
	}
}

func TestRemoteStartTransactionStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	status := rt.RemoteStartTransactionStatus("Pending")
	if status.IsValid() {
		t.Errorf(
			st.ErrorIsValidTrue,
			"RemoteStartTransactionStatus(\"Pending\")",
		)
	}
}

func TestRemoteStartTransactionStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := rt.RemoteStartTransactionStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"RemoteStartTransactionStatus.String()",
			got,
			statusAcceptedStr,
		)
	}
}

func TestRemoteStartTransactionStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := rt.RemoteStartTransactionStatusRejected.String()
	if got != statusRejectedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			"RemoteStartTransactionStatus.String()",
			got,
			statusRejectedStr,
		)
	}
}
