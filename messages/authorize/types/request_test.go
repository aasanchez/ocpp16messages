package types

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestRequest_Validate_Valid(t *testing.T) {
	t.Parallel()

	req := Request{
		IdTag: "RFID-TAG-123",
	}

	err := req.Validate()
	if err != nil {
		t.Errorf("Request.Validate() error = %v, want nil", err)
	}
}

func TestRequest_Validate_Empty(t *testing.T) {
	t.Parallel()

	req := Request{
		IdTag: "",
	}

	err := req.Validate()
	if err == nil {
		t.Error("Request.Validate() error = nil, want error for empty IdTag")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(
			"Request.Validate() error = %v, want error wrapping %v",
			err,
			st.ErrInvalidValue,
		)
	}
}

func TestRequest_Validate_TooLong(t *testing.T) {
	t.Parallel()

	req := Request{
		IdTag: "RFID-ABC123456789012345", // 25 characters, max is 20
	}

	err := req.Validate()
	if err == nil {
		t.Error("Request.Validate() error = nil, want error for IdTag too long")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(
			"Request.Validate() error = %v, want 'exceeds maximum length'",
			err,
		)
	}
}

func TestRequest_Validate_InvalidCharacters(t *testing.T) {
	t.Parallel()

	req := Request{
		IdTag: "RFID\x00TAG", // Contains null byte (non-printable ASCII)
	}

	err := req.Validate()
	if err == nil {
		t.Error("Request.Validate() error = nil, want error for invalid chars")
	}

	if !strings.Contains(err.Error(), "non-printable ASCII") {
		t.Errorf(
			"Request.Validate() error = %v, want 'non-printable ASCII'",
			err,
		)
	}
}

func TestRequest_Value(t *testing.T) {
	t.Parallel()

	expectedTag := "RFID-TAG-456"
	req := Request{
		IdTag: expectedTag,
	}

	if req.Value() != expectedTag {
		t.Errorf(st.ErrorMismatch, expectedTag, req.Value())
	}
}
