package authorize

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestConfirmation_validPayload(t *testing.T) {
	t.Parallel()

	expiry := "2027-04-12T10:03:04Z"
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
		t.Fatalf("expected no error, got: %v", err)
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

func TestConfirmation_invalidStatus(t *testing.T) {
	t.Parallel()

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: "UnknownStatus",
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for invalid status, got nil")
	}
}

func TestConfirmation_invalidExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "not-a-date"
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:     "Accepted",
			ExpiryDate: &invalidDate,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for invalid expiryDate, got nil")
	}
}

func TestConfirmation_invalidParentIdTag(t *testing.T) {
	t.Parallel()

	invalidTag := string(make([]byte, 100)) // exceeds 20 characters, invalid CiString20
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ParentIdTag: &invalidTag,
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for invalid parentIdTag, got nil")
	}
}

func TestConfirmation_payloadValidationFails_emptyStatus(t *testing.T) {
	t.Parallel()

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: "", // triggers input.Validate() failure
		},
	}

	_, err := Confirmation(payload)
	if err == nil {
		t.Fatal("expected error for empty status in payload, got nil")
	}
}
