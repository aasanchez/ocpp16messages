package authorize

import (
	"strings"
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const ()

func TestAuthorizeConfirmation_validPayload(t *testing.T) {
	t.Parallel()

	expiry := "2025-04-12T10:03:04Z"
	parent := "A632-E2BB0231072C"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	msg, err := Confirmation(payload)
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}

	val := msg.IdTagInfo.Value()

	if val.Status != "Accepted" {
		t.Errorf("unexpected status: got %s", val.Status)
	}

	if val.ExpiryDate == nil || *val.ExpiryDate != expiry {
		t.Errorf("unexpected expiry date: got %+v", val.ExpiryDate)
	}

	if val.ParentIdTag == nil || *val.ParentIdTag != parent {
		t.Errorf("unexpected parentIdTag: got %+v", val.ParentIdTag)
	}
}

func TestAuthorizeConfirmation_invalidStatus(t *testing.T) {
	t.Parallel()

	expiry := "2026-01-01T00:00:00Z"
	parent := "PARENT"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "UnknownStatus",
			ExpiryDate:  &expiry,
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

	invalidDate := "not-a-date"
	parent := "PARENT1"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &invalidDate,
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

	expiry := "2027-05-01T00:00:00Z"
	invalidTag := strings.Repeat("X", 100) // invalid CiString20

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &expiry,
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

	expiry := "2027-01-03T00:00:00Z"
	parent := "PARENT2"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "",
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for empty status in payload, got nil")
	}
}
