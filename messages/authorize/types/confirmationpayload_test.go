package authorizetypes

import (
	"testing"
)

func TestConfirmationPayload_validStatusOnly(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{
		IdTagInfo: IdTagInfoPayload{
			Status:      Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	err := payload.Validate()
	if err != nil {
		t.Fatalf("expected valid payload, got error -> %v", err)
	}

	got := payload.Value()

	if got.IdTagInfo.Status != Accepted {
		t.Errorf("expected status %q, got %q", Accepted, got.IdTagInfo.Status)
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

	payload := ConfirmationPayload{
		IdTagInfo: IdTagInfoPayload{
			Status:      Accepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	err := payload.Validate()
	if err != nil {
		t.Fatalf("expected valid payload, got error -> %v", err)
	}

	got := payload.Value()

	if got.IdTagInfo.Status != Accepted {
		t.Errorf("status mismatch: expected %q, got %q", Accepted, got.IdTagInfo.Status)
	}

	if got.IdTagInfo.ExpiryDate == nil || *got.IdTagInfo.ExpiryDate != expiry {
		t.Errorf("expiryDate mismatch -> expected %q, got %v", expiry, got.IdTagInfo.ExpiryDate)
	}

	if got.IdTagInfo.ParentIdTag == nil || *got.IdTagInfo.ParentIdTag != parent {
		t.Errorf("parentIdTag mismatch -> expected %q, got %v", parent, got.IdTagInfo.ParentIdTag)
	}
}

func TestConfirmationPayload_missingStatus(t *testing.T) {
	t.Parallel()

	payload := ConfirmationPayload{
		IdTagInfo: IdTagInfoPayload{
			Status:      "",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	err := payload.Validate()
	if err == nil {
		t.Fatal("expected validation error for missing status, got nil")
	}
}
