package authorize

import (
	"strings"
	"testing"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestAuthorizeConfirmation_validPayload(t *testing.T) {
	t.Parallel()
	t.Helper()

	parent := "A632-E2BB0231072C"
	expectedExpiry := "2025-04-12T10:03:04Z"

	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  expectedExpiry,
			ParentIdTag: &parent,
		},
	}

	msg, err := Confirmation(payload)
	if err != nil {
		t.Fatalf(st.ErrExpectedNoError, err)
	}

	val := msg.IdTagInfo.Value()

	if val.Status != "Accepted" {
		t.Errorf("unexpected status: got %s", val.Status)
	}

	if val.ExpiryDate != expectedExpiry {
		t.Errorf("unexpected expiry date: got %s, want %s", val.ExpiryDate, expectedExpiry)
	}

	if val.ParentIdTag == nil || *val.ParentIdTag != parent {
		t.Errorf("unexpected parentIdTag: got %+v, want %s", val.ParentIdTag, parent)
	}
}

func TestAuthorizeConfirmation_invalidStatus(t *testing.T) {
	t.Parallel()

	parent := "PARENT"

	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "UnknownStatus",
			ExpiryDate:  "2026-01-01T00:00:00Z",
			ParentIdTag: &parent,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for invalid status, got nil")
	}
}

func TestAuthorizeConfirmation_invalidExpiryDate(t *testing.T) {
	t.Parallel()

	parent := "PARENT1"

	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  "not-a-date",
			ParentIdTag: &parent,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for invalid expiryDate, got nil")
	}
}

func TestAuthorizeConfirmation_invalidParentIdTag(t *testing.T) {
	t.Parallel()

	invalidTag := strings.Repeat("X", 100) // invalid CiString20

	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  "2027-05-01T00:00:00Z",
			ParentIdTag: &invalidTag,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for invalid parentIdTag, got nil")
	}
}

func TestAuthorizeConfirmation_payloadValidationFails_emptyStatus(t *testing.T) {
	t.Parallel()

	parent := "PARENT2"

	payload := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "",
			ExpiryDate:  "2027-01-03T00:00:00Z",
			ParentIdTag: &parent,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for empty status in payload, got nil")
	}
}
