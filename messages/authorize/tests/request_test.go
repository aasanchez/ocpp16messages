package authorize_test

import (
	"errors"
	"strings"
	"testing"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestNewRequest_EmptyIdTag(t *testing.T) {
	t.Parallel()

	payload := mat.Request{
		IdTag: "",
	}

	_, err := ma.NewRequest(payload)
	if err == nil {
		t.Error(err)
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestNewRequest_IdTagTooLong(t *testing.T) {
	t.Parallel()

	payload := mat.Request{
		IdTag: "RFID-ABC123456789012345", // 25 chars, max is 20
	}

	_, err := ma.NewRequest(payload)
	if err == nil {
		t.Error("NewRequest() error = nil, want error for IdTag too long")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(
			"NewRequest() error = %v, want 'exceeds maximum length'",
			err,
		)
	}
}

func TestNewRequest_InvalidCharacters(t *testing.T) {
	t.Parallel()

	payload := mat.Request{
		IdTag: "RFID\x00ABC", // Contains null byte (non-printable ASCII)
	}

	_, err := ma.NewRequest(payload)
	if err == nil {
		t.Error(
			"NewRequest() error = nil, want error for non-printable chars",
		)
	}

	if !strings.Contains(err.Error(), "non-printable ASCII") {
		t.Errorf(
			"NewRequest() error = %v, want 'non-printable ASCII'",
			err,
		)
	}
}
