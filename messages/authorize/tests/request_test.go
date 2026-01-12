package authorize_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const testValidIdTag = "RFID-TAG-12345"

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := authorize.Req(authorize.Input{IdTag: testValidIdTag})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.IdTag.String() != testValidIdTag {
		t.Errorf(st.ErrorMismatch, testValidIdTag, req.IdTag.String())
	}
}

func TestReq_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := authorize.Req(authorize.Input{IdTag: ""})
	if err == nil {
		t.Error("Req() error = nil, want error for empty idTag")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestReq_IdTagTooLong(t *testing.T) {
	t.Parallel()

	// 23 chars, max is 20
	_, err := authorize.Req(authorize.Input{IdTag: "RFID-ABC123456789012345"})
	if err == nil {
		t.Error("Req() error = nil, want error for IdTag too long")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(
			"Req() error = %v, want 'exceeds maximum length'",
			err,
		)
	}
}

func TestReq_InvalidCharacters(t *testing.T) {
	t.Parallel()

	// Contains null byte
	_, err := authorize.Req(authorize.Input{IdTag: "RFID\x00ABC"})
	if err == nil {
		t.Error(
			"Req() error = nil, want error for non-printable chars",
		)
	}

	if !strings.Contains(err.Error(), "non-printable ASCII") {
		t.Errorf(
			"Req() error = %v, want 'non-printable ASCII'",
			err,
		)
	}
}
