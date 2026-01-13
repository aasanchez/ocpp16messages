package types_test

import (
	"testing"

	mat "github.com/aasanchez/ocpp16messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusAcceptedStr     = "Accepted"
	statusBlockedStr      = "Blocked"
	statusExpiredStr      = "Expired"
	statusInvalidStr      = "Invalid"
	statusConcurrentTxStr = "ConcurrentTx"
	methodString          = "AuthorizationStatus.String()"
)

func TestAuthorizationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !mat.AuthorizationStatusAccepted.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusAccepted")
	}
}

func TestAuthorizationStatus_IsValid_Blocked(t *testing.T) {
	t.Parallel()

	if !mat.AuthorizationStatusBlocked.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusBlocked")
	}
}

func TestAuthorizationStatus_IsValid_Expired(t *testing.T) {
	t.Parallel()

	if !mat.AuthorizationStatusExpired.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusExpired")
	}
}

func TestAuthorizationStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	if !mat.AuthorizationStatusInvalid.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusInvalid")
	}
}

func TestAuthorizationStatus_IsValid_ConcurrentTx(t *testing.T) {
	t.Parallel()

	if !mat.AuthorizationStatusConcurrentTx.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "AuthorizationStatusConcurrentTx")
	}
}

func TestAuthorizationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := mat.AuthorizationStatus("")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "AuthorizationStatus(\"\")")
	}
}

func TestAuthorizationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := mat.AuthorizationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "AuthorizationStatus(\"Unknown\")")
	}
}

func TestAuthorizationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := mat.AuthorizationStatus("accepted")
	if status.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "AuthorizationStatus(\"accepted\")")
	}
}

func TestAuthorizationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := mat.AuthorizationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestAuthorizationStatus_String_Blocked(t *testing.T) {
	t.Parallel()

	got := mat.AuthorizationStatusBlocked.String()
	if got != statusBlockedStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusBlockedStr,
		)
	}
}

func TestAuthorizationStatus_String_Expired(t *testing.T) {
	t.Parallel()

	got := mat.AuthorizationStatusExpired.String()
	if got != statusExpiredStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusExpiredStr,
		)
	}
}

func TestAuthorizationStatus_String_Invalid(t *testing.T) {
	t.Parallel()

	got := mat.AuthorizationStatusInvalid.String()
	if got != statusInvalidStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusInvalidStr,
		)
	}
}

func TestAuthorizationStatus_String_ConcurrentTx(t *testing.T) {
	t.Parallel()

	got := mat.AuthorizationStatusConcurrentTx.String()
	if got != statusConcurrentTxStr {
		t.Errorf(
			st.ErrorMethodMismatch,
			methodString,
			got,
			statusConcurrentTxStr,
		)
	}
}
