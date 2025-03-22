// Package test contains integration-style tests for the ocpp16_messages package.
// These tests validate the parsing and schema enforcement of specific OCPP 1.6J messages.
package test

import (
	"testing"
	"time"

	"github.com/aasanchez/ocpp16_messages"
	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

// TestAuthorizeReqValidation checks multiple cases for Authorize.req validation.
func TestAuthorizeReqValidation(t *testing.T) {
	tests := []struct {
		name    string
		raw     []byte
		wantErr bool
	}{
		{
			name:    "Valid Authorize.req",
			raw:     []byte(`[2, "uid123", "Authorize", {"idTag":"ABC123"}]`),
			wantErr: false,
		},
		{
			name:    "Missing idTag",
			raw:     []byte(`[2, "uid123", "Authorize", {}]`),
			wantErr: true,
		},
		{
			name:    "idTag too long",
			raw:     []byte(`[2, "uid123", "Authorize", {"idTag":"123456789012345678901"}]`), // 21 chars
			wantErr: true,
		},
		{
			name:    "Malformed JSON",
			raw:     []byte(`[2, "uid123", "Authorize", {"idTag":123}]`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ocpp16_messages.ValidateMessage(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestAuthorizeConfValidation checks various valid and invalid Authorize.conf responses.
func TestAuthorizeConfValidation(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		conf    authorize.AuthorizeConf
		wantErr bool
	}{
		{
			name: "Valid response",
			conf: authorize.AuthorizeConf{
				IdTagInfo: authorize.IdTagInfo{
					Status: core.AuthorizationAccepted,
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid status value",
			conf: authorize.AuthorizeConf{
				IdTagInfo: authorize.IdTagInfo{
					Status: core.AuthorizationStatus("UNKNOWN"),
				},
			},
			wantErr: true,
		},
		{
			name: "Valid with expiry and parent",
			conf: authorize.AuthorizeConf{
				IdTagInfo: authorize.IdTagInfo{
					Status:     core.AuthorizationAccepted,
					ExpiryDate: &now,
					ParentIdTag: &core.IdToken{
						IdTag: core.CiString20Type("PARENT"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid parentIdTag",
			conf: authorize.AuthorizeConf{
				IdTagInfo: authorize.IdTagInfo{
					Status: core.AuthorizationAccepted,
					ParentIdTag: &core.IdToken{
						IdTag: core.CiString20Type(""),
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := authorize.ValidateAuthorizeConf(tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAuthorizeConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
