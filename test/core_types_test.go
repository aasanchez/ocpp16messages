package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestCiStringValidation(t *testing.T) {
	tests := []struct {
		name    string
		value   core.CiString20
		wantErr bool
	}{
		{"Valid short string", "abc", false},
		{"Valid max length string", core.CiString20("12345678901234567890"), false},
		{"Empty string", "", true},
		{"Too long", core.CiString20("123456789012345678901"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.value.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("CiString20.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnumAuthorizationStatus(t *testing.T) {
	tests := []struct {
		name   string
		status core.AuthorizationStatus
		valid  bool
	}{
		{"Valid Accepted", core.AuthorizationAccepted, true},
		{"Valid Blocked", core.AuthorizationBlocked, true},
		{"Invalid Custom", core.AuthorizationStatus("Custom"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.status.IsValid(); got != tt.valid {
				t.Errorf("AuthorizationStatus.IsValid() = %v, want %v", got, tt.valid)
			}
		})
	}
}
