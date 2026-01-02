package types

import "testing"

func TestAuthorizationStatus_IsValid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status AuthorizationStatus
		want   bool
	}{
		{"Accepted", AuthorizationStatusAccepted, true},
		{"Blocked", AuthorizationStatusBlocked, true},
		{"Expired", AuthorizationStatusExpired, true},
		{"Invalid", AuthorizationStatusInvalid, true},
		{"ConcurrentTx", AuthorizationStatusConcurrentTx, true},
		{"Empty", AuthorizationStatus(""), false},
		{"Unknown", AuthorizationStatus("Unknown"), false},
		{"Lowercase", AuthorizationStatus("accepted"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf(
					"AuthorizationStatus.IsValid() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestAuthorizationStatus_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status AuthorizationStatus
		want   string
	}{
		{"Accepted", AuthorizationStatusAccepted, "Accepted"},
		{"Blocked", AuthorizationStatusBlocked, "Blocked"},
		{"Expired", AuthorizationStatusExpired, "Expired"},
		{"Invalid", AuthorizationStatusInvalid, "Invalid"},
		{"ConcurrentTx", AuthorizationStatusConcurrentTx, "ConcurrentTx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.status.String(); got != tt.want {
				t.Errorf(
					"AuthorizationStatus.String() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
