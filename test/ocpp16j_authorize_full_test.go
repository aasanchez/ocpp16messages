package test

import (
	"testing"
	"time"

	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

func TestAuthorizeReqValidation(t *testing.T) {
	tests := []struct {
		name    string
		rawJSON string
		wantErr bool
	}{
		{
			name:    "Valid Authorize.req",
			rawJSON: `[2,"123456","Authorize",{"idTag":"ABC123"}]`,
			wantErr: false,
		},
		{
			name:    "Missing idTag",
			rawJSON: `[2,"123456","Authorize",{}]`,
			wantErr: true,
		},
		{
			name:    "idTag too long",
			rawJSON: `[2,"123456","Authorize",{"idTag":"123456789012345678901"}]`,
			wantErr: true,
		},
		{
			name:    "Malformed JSON",
			rawJSON: `[2,"123456","Authorize",{"idTag":12345}]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := core.ValidateRawMessage([]byte(tt.rawJSON))
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateRawMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizeConfValidation(t *testing.T) {
	tests := []struct {
		name    string
		conf    authorize.Conf
		wantErr bool
	}{
		{
			name: "Valid response",
			conf: authorize.Conf{
				IdTagInfo: authorize.IdTagInfo{
					Status: "Accepted",
				},
			},
		},
		{
			name: "Invalid status value",
			conf: authorize.Conf{
				IdTagInfo: authorize.IdTagInfo{
					Status: "InvalidStatus",
				},
			},
			wantErr: true,
		},
		{
			name: "Valid with expiry and parent",
			conf: authorize.Conf{
				IdTagInfo: authorize.IdTagInfo{
					Status:     "Accepted",
					ExpiryDate: func() *time.Time { t := time.Now().Add(time.Hour); return &t }(),
					ParentIdTag: func() *core.CiString20 {
						s := core.CiString20("parent123")
						return &s
					}(),
				},
			},
		},
		{
			name: "Invalid parentIdTag",
			conf: authorize.Conf{
				IdTagInfo: authorize.IdTagInfo{
					Status: "Accepted",
					ParentIdTag: func() *core.CiString20 {
						s := core.CiString20("this_string_is_way_too_long_for_limit")
						return &s
					}(),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := authorize.ValidateConf(tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCallErrorMessageParsing(t *testing.T) {
	tests := []struct {
		name      string
		raw       string
		expectErr bool
	}{
		{
			name:      "Valid CALLERROR message",
			raw:       `[4,"123","ProtocolError","Failed to parse message",{}]`,
			expectErr: false,
		},
		{
			name:      "Missing elements",
			raw:       `[4,"123","ProtocolError"]`,
			expectErr: true,
		},
		{
			name:      "Non-string error code",
			raw:       `[4,"123",1234,"Message",{}]`,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := core.ValidateRawMessage([]byte(tt.raw))
			if (err != nil) != tt.expectErr {
				t.Errorf("ValidateRawMessage() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
