package types

import "testing"

const (
	statusAcceptedStr     = "Accepted"
	statusBlockedStr      = "Blocked"
	statusExpiredStr      = "Expired"
	statusInvalidStr      = "Invalid"
	statusConcurrentTxStr = "ConcurrentTx"
)

func TestAuthorizationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusAccepted.IsValid() {
		t.Error("AuthorizationStatusAccepted.IsValid() = false, want true")
	}
}

func TestAuthorizationStatus_IsValid_Blocked(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusBlocked.IsValid() {
		t.Error("AuthorizationStatusBlocked.IsValid() = false, want true")
	}
}

func TestAuthorizationStatus_IsValid_Expired(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusExpired.IsValid() {
		t.Error("AuthorizationStatusExpired.IsValid() = false, want true")
	}
}

func TestAuthorizationStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusInvalid.IsValid() {
		t.Error("AuthorizationStatusInvalid.IsValid() = false, want true")
	}
}

func TestAuthorizationStatus_IsValid_ConcurrentTx(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusConcurrentTx.IsValid() {
		t.Error("AuthorizationStatusConcurrentTx.IsValid() = false, want true")
	}
}

func TestAuthorizationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatus("")
	if status.IsValid() {
		t.Error("AuthorizationStatus(\"\").IsValid() = true, want false")
	}
}

func TestAuthorizationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatus("Unknown")
	if status.IsValid() {
		t.Error("AuthorizationStatus(\"Unknown\").IsValid() = true, want false")
	}
}

func TestAuthorizationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatus("accepted")
	if status.IsValid() {
		t.Error(
			"AuthorizationStatus(\"accepted\").IsValid() = true, want false",
		)
	}
}

func TestAuthorizationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	if got := AuthorizationStatusAccepted.String(); got != statusAcceptedStr {
		t.Errorf(ErrorMethodMismatch, got, statusAcceptedStr)
	}
}

func TestAuthorizationStatus_String_Blocked(t *testing.T) {
	t.Parallel()

	if got := AuthorizationStatusBlocked.String(); got != statusBlockedStr {
		t.Errorf(ErrorMethodMismatch, got, statusBlockedStr)
	}
}

func TestAuthorizationStatus_String_Expired(t *testing.T) {
	t.Parallel()

	if got := AuthorizationStatusExpired.String(); got != statusExpiredStr {
		t.Errorf(ErrorMethodMismatch, got, statusExpiredStr)
	}
}

func TestAuthorizationStatus_String_Invalid(t *testing.T) {
	t.Parallel()

	if got := AuthorizationStatusInvalid.String(); got != statusInvalidStr {
		t.Errorf(ErrorMethodMismatch, got, statusInvalidStr)
	}
}

func TestAuthorizationStatus_String_ConcurrentTx(t *testing.T) {
	t.Parallel()

	got := AuthorizationStatusConcurrentTx.String()
	if got != statusConcurrentTxStr {
		t.Errorf(ErrorMethodMismatch, got, statusConcurrentTxStr)
	}
}
