package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestConfirmationPayload_validStatusOnly(t *testing.T) {
	t.Parallel()

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: authorizetypes.Accepted,
		},
	}

	if err := payload.Validate(); err != nil {
		t.Fatalf("expected valid payload, got error: %v", err)
	}

	got := payload.Value()

	if got.IdTagInfo.Status != authorizetypes.Accepted {
		t.Errorf("expected status %q, got %q", authorizetypes.Accepted, got.IdTagInfo.Status)
	}

	if got.IdTagInfo.ExpiryDate != nil {
		t.Errorf("expected expiryDate nil, got %v", *got.IdTagInfo.ExpiryDate)
	}

	if got.IdTagInfo.ParentIdTag != nil {
		t.Errorf("expected parentIdTag nil, got %v", *got.IdTagInfo.ParentIdTag)
	}
}

func TestConfirmationPayload_withAllFields(t *testing.T) {
	t.Parallel()

	expiry := "2027-04-12T10:03:04Z"
	parent := "ABC123"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	if err := payload.Validate(); err != nil {
		t.Fatalf("expected valid payload, got error: %v", err)
	}

	got := payload.Value()

	if got.IdTagInfo.Status != authorizetypes.Accepted {
		t.Errorf("status mismatch: expected %q, got %q", authorizetypes.Accepted, got.IdTagInfo.Status)
	}
	if got.IdTagInfo.ExpiryDate == nil || *got.IdTagInfo.ExpiryDate != expiry {
		t.Errorf("expiryDate mismatch: expected %q, got %v", expiry, got.IdTagInfo.ExpiryDate)
	}
	if got.IdTagInfo.ParentIdTag == nil || *got.IdTagInfo.ParentIdTag != parent {
		t.Errorf("parentIdTag mismatch: expected %q, got %v", parent, got.IdTagInfo.ParentIdTag)
	}
}

func TestConfirmationPayload_missingStatus(t *testing.T) {
	t.Parallel()

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: "", // Missing required status
		},
	}

	err := payload.Validate()
	if err == nil {
		t.Fatal("expected validation error for missing status, got nil")
	}
}
