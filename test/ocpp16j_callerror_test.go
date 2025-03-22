// Package test contains integration tests for OCPP 1.6J message handling.
// This file tests the parsing of CALLERROR messages (MessageTypeId = 4).
package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages"
)

// TestCallErrorMessageParsing ensures that MessageTypeId 4 messages (CALLERROR)
// are correctly parsed and wrapped in a RawMessage.
func TestCallErrorMessageParsing(t *testing.T) {
	tests := []struct {
		name       string
		raw        []byte
		expectCode string
		expectErr  bool
	}{
		{
			name:       "Valid CALLERROR message",
			raw:        []byte(`[4, "uid123", "ProtocolError", "Invalid payload", {"errorDetail":"something"}]`),
			expectCode: "ProtocolError",
			expectErr:  false,
		},
		{
			name:      "Missing elements",
			raw:       []byte(`[4, "uid123", "ProtocolError"]`),
			expectErr: true,
		},
		{
			name:      "Non-string error code",
			raw:       []byte(`[4, "uid123", 123, "Error", {}]`),
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ocpp16_messages.ValidateMessage(tt.raw)
			if (err != nil) != tt.expectErr {
				t.Errorf("ValidateMessage() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if err == nil && result.Action != tt.expectCode {
				t.Errorf("Expected error code %q, got %q", tt.expectCode, result.Action)
			}
		})
	}
}
