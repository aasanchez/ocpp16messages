package authorize

import (
	"testing"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestNewRequest_EmptyIdTag(t *testing.T) {
	t.Parallel()

	payload := mat.Request{
		IdTag: "",
	}

	_, err := NewRequest(payload)
	if err == nil {
		t.Error("NewRequest() error = nil, want error for empty IdTag")
	}
}

func TestNewRequest_IdTagTooLong(t *testing.T) {
	t.Parallel()

	payload := mat.Request{
		IdTag: "RFID-ABC123456789012345", // 25 chars, max is 20
	}

	_, err := NewRequest(payload)
	if err == nil {
		t.Error("NewRequest() error = nil, want error for IdTag too long")
	}
}

func TestNewRequest_InvalidCharacters(t *testing.T) {
	t.Parallel()

	payload := mat.Request{
		IdTag: "RFID\x00ABC", // Contains null byte (non-printable ASCII)
	}

	_, err := NewRequest(payload)
	if err == nil {
		t.Error(
			"NewRequest() error = nil, want error for non-printable chars",
		)
	}
}
