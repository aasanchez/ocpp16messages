// Package test provides external test coverage for OCPP 1.6 message validation logic.
//
// These tests validate AuthorizeReq and AuthorizeConf message types
// against the OCPP 1.6J specification using the validators package.
package test

import (
	"testing"
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
	"github.com/aasanchez/ocpp16_messages/models"
	"github.com/aasanchez/ocpp16_messages/validators"
)

func TestAuthorizeReq(t *testing.T) {
	tests := []struct {
		name    string
		req     chargePoint.AuthorizeReq
		wantErr bool
	}{
		{
			name: "Valid IdTag",
			req:  chargePoint.AuthorizeReq{IdTag: models.CiString20Type("ABC123")},
		},
		{
			name:    "Empty IdTag",
			req:     chargePoint.AuthorizeReq{IdTag: models.CiString20Type("")},
			wantErr: true,
		},
		{
			name:    "IdTag exceeds 20 characters",
			req:     chargePoint.AuthorizeReq{IdTag: models.CiString20Type("123456789012345678901")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validators.ValidateAuthorizeReq(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAuthorizeReq() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizeConf(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name    string
		conf    chargePoint.AuthorizeConf
		wantErr bool
	}{
		{
			name: "Valid Accepted",
			conf: chargePoint.AuthorizeConf{
				IdTagInfo: chargePoint.IdTagInfo{
					Status: enums.AuthorizationAccepted,
				},
			},
		},
		{
			name: "Valid with ExpiryDate and ParentIdTag",
			conf: chargePoint.AuthorizeConf{
				IdTagInfo: chargePoint.IdTagInfo{
					Status:      enums.AuthorizationAccepted,
					ExpiryDate:  &now,
					ParentIdTag: ptr(models.CiString20Type("Parent123")),
				},
			},
		},
		{
			name: "Invalid AuthorizationStatus",
			conf: chargePoint.AuthorizeConf{
				IdTagInfo: chargePoint.IdTagInfo{
					Status: enums.AuthorizationStatus("Unknown"),
				},
			},
			wantErr: true,
		},
		{
			name: "ParentIdTag too long",
			conf: chargePoint.AuthorizeConf{
				IdTagInfo: chargePoint.IdTagInfo{
					Status:      enums.AuthorizationAccepted,
					ParentIdTag: ptr(models.CiString20Type("123456789012345678901")),
				},
			},
			wantErr: true,
		},
		{
			name: "Empty ParentIdTag",
			conf: chargePoint.AuthorizeConf{
				IdTagInfo: chargePoint.IdTagInfo{
					Status:      enums.AuthorizationAccepted,
					ParentIdTag: ptr(models.CiString20Type("")),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validators.ValidateAuthorizeConf(tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAuthorizeConf() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
