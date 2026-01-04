package types

import "testing"

const (
	statusAcceptedStr    = "Accepted"
	statusBlockedStr     = "Blocked"
	statusExpiredStr     = "Expired"
	statusInvalidStr     = "Invalid"
	statusConcurrentTxStr = "ConcurrentTx"
)

func TestAuthorizationStatus_IsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status AuthorizationStatus
		want   bool
	}{
		{statusAcceptedStr, AuthorizationStatusAccepted, true},
		{statusBlockedStr, AuthorizationStatusBlocked, true},
		{statusExpiredStr, AuthorizationStatusExpired, true},
		{statusInvalidStr, AuthorizationStatusInvalid, true},
		{statusConcurrentTxStr, AuthorizationStatusConcurrentTx, true},
		{"Empty", AuthorizationStatus(""), false},
		{"Unknown", AuthorizationStatus("Unknown"), false},
		{"Lowercase", AuthorizationStatus("accepted"), false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			if got := testCase.status.IsValid(); got != testCase.want {
				t.Errorf(
					"AuthorizationStatus.IsValid() = %v, want %v",
					got,
					testCase.want,
				)
			}
		})
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
