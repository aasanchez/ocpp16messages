package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestAuthorizeReqValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid Authorize.req",
			input:   `[2,"123456","Authorize",{"idTag":{"idTag":"D0431F35"}}]`,
			wantErr: false,
		},
		{
			name:    "Missing idTag",
			input:   `[2,"123457","Authorize",{}]`,
			wantErr: true,
		},
		{
			name:    "idTag too long",
			input:   `[2,"123458","Authorize",{"idTag":{"idTag":"012345678901234567890"}}]`,
			wantErr: true,
		},
		{
			name:    "Malformed JSON",
			input:   `[2,"123459","Authorize",{"idTag":}]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := []byte(tt.input)
			parsed, err := core.ValidateRawMessage(msg)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected parse error: %v", err)
				}
				return
			}
			err = core.ValidateMessage(parsed.Action, parsed.Payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizeConfValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid response",
			input:   `[3,"123456","AuthorizeResponse",{"idTagInfo":{"status":"Accepted"}}]`,
			wantErr: false,
		},
		{
			name:    "Invalid status value",
			input:   `[3,"123457","AuthorizeResponse",{"idTagInfo":{"status":"INVALID_STATUS"}}]`,
			wantErr: true,
		},
		{
			name:    "Valid with expiry and parent",
			input:   `[3,"123458","AuthorizeResponse",{"idTagInfo":{"status":"Accepted","expiryDate":"2024-03-22T12:00:00Z","parentIdTag":"ABC123"}}]`,
			wantErr: false,
		},
		{
			name:    "Invalid parentIdTag",
			input:   `[3,"123459","AuthorizeResponse",{"idTagInfo":{"status":"Accepted","parentIdTag":"0123456789012345678901"}}]`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := []byte(tt.input)
			parsed, err := core.ValidateRawMessage(msg)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected parse error: %v", err)
				}
				return
			}
			err = core.ValidateMessage(parsed.Action, parsed.Payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
