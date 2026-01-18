package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
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

	if !types.AuthorizationStatusAccepted.IsValid() {
		t.Errorf(types.ErrorIsValidFalse, "AuthorizationStatusAccepted")
	}
}

func TestAuthorizationStatus_IsValid_Blocked(t *testing.T) {
	t.Parallel()

	if !types.AuthorizationStatusBlocked.IsValid() {
		t.Errorf(types.ErrorIsValidFalse, "AuthorizationStatusBlocked")
	}
}

func TestAuthorizationStatus_IsValid_Expired(t *testing.T) {
	t.Parallel()

	if !types.AuthorizationStatusExpired.IsValid() {
		t.Errorf(types.ErrorIsValidFalse, "AuthorizationStatusExpired")
	}
}

func TestAuthorizationStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	if !types.AuthorizationStatusInvalid.IsValid() {
		t.Errorf(types.ErrorIsValidFalse, "AuthorizationStatusInvalid")
	}
}

func TestAuthorizationStatus_IsValid_ConcurrentTx(t *testing.T) {
	t.Parallel()

	if !types.AuthorizationStatusConcurrentTx.IsValid() {
		t.Errorf(types.ErrorIsValidFalse, "AuthorizationStatusConcurrentTx")
	}
}

func TestAuthorizationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := types.AuthorizationStatus("")
	if status.IsValid() {
		t.Errorf(types.ErrorIsValidTrue, "AuthorizationStatus(\"\")")
	}
}

func TestAuthorizationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := types.AuthorizationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(types.ErrorIsValidTrue, "AuthorizationStatus(\"Unknown\")")
	}
}

func TestAuthorizationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := types.AuthorizationStatus("accepted")
	if status.IsValid() {
		t.Errorf(types.ErrorIsValidTrue, "AuthorizationStatus(\"accepted\")")
	}
}

func TestAuthorizationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := types.AuthorizationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestAuthorizationStatus_String_Blocked(t *testing.T) {
	t.Parallel()

	got := types.AuthorizationStatusBlocked.String()
	if got != statusBlockedStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodString,
			got,
			statusBlockedStr,
		)
	}
}

func TestAuthorizationStatus_String_Expired(t *testing.T) {
	t.Parallel()

	got := types.AuthorizationStatusExpired.String()
	if got != statusExpiredStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodString,
			got,
			statusExpiredStr,
		)
	}
}

func TestAuthorizationStatus_String_Invalid(t *testing.T) {
	t.Parallel()

	got := types.AuthorizationStatusInvalid.String()
	if got != statusInvalidStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodString,
			got,
			statusInvalidStr,
		)
	}
}

func TestAuthorizationStatus_String_ConcurrentTx(t *testing.T) {
	t.Parallel()

	got := types.AuthorizationStatusConcurrentTx.String()
	if got != statusConcurrentTxStr {
		t.Errorf(
			types.ErrorMethodMismatch,
			methodString,
			got,
			statusConcurrentTxStr,
		)
	}
}
