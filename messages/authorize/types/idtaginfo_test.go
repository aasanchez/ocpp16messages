package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestIdTagInfo_statusOnly(t *testing.T) {
	t.Parallel()

	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	info, err := authorizetypes.IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error statusOnly: %v", err)
	}

	val := info.Value()

	if val.Status != authorizetypes.Accepted {
		t.Errorf("expected status %q, got %q", authorizetypes.Accepted, val.Status)
	}

	if val.ExpiryDate != nil {
		t.Errorf("expected expiryDate nil, got %v", *val.ExpiryDate)
	}

	if val.ParentIdTag != nil {
		t.Errorf("expected parentIdTag nil, got %v", *val.ParentIdTag)
	}
}

func TestIdTagInfo_withExpiryDateOnly(t *testing.T) {
	t.Parallel()

	exp := "2023-04-12T10:03:04Z"

	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  &exp,
		ParentIdTag: nil,
	}

	info, err := authorizetypes.IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error withExpiryDateOnly: %v", err)
	}

	val := info.Value()
	if val.ExpiryDate == nil || *val.ExpiryDate != exp {
		t.Errorf("expiryDate mismatch: want %s, got %v", exp, val.ExpiryDate)
	}
}

func TestIdTagInfo_withParentIdTagOnly(t *testing.T) {
	t.Parallel()

	parent := "ABC123"

	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &parent,
	}

	info, err := authorizetypes.IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	val := info.Value()
	if val.ParentIdTag == nil || *val.ParentIdTag != parent {
		t.Errorf("parentIdTag mismatch: want %s, got %v", parent, val.ParentIdTag)
	}
}

func TestIdTagInfo_allFieldsPresent(t *testing.T) {
	t.Parallel()

	exp := "2027-04-12T10:01:04Z"
	parent := "XYZ-987"

	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  &exp,
		ParentIdTag: &parent,
	}

	info, err := authorizetypes.IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	val := info.Value()

	if val.Status != authorizetypes.Accepted {
		t.Errorf("status mismatch: got %s", val.Status)
	}

	if val.ExpiryDate == nil || *val.ExpiryDate != exp {
		t.Errorf("expiryDate mismatch: want %s, got %v", exp, val.ExpiryDate)
	}

	if val.ParentIdTag == nil || *val.ParentIdTag != parent {
		t.Errorf("parentIdTag mismatch: want %s, got %v", parent, val.ParentIdTag)
	}
}

func TestIdTagInfo_invalidStatus(t *testing.T) {
	t.Parallel()

	input := authorizetypes.IdTagInfoPayload{
		Status:      "InvalidStatus",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	_, err := authorizetypes.IdTagInfo(input)
	if err == nil {
		t.Fatal("expected error for invalid status, got nil")
	}
}

func TestIdTagInfo_invalidExpiryDate(t *testing.T) {
	t.Parallel()

	bad := "not-a-date"

	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  &bad,
		ParentIdTag: nil,
	}

	_, err := authorizetypes.IdTagInfo(input)
	if err == nil {
		t.Fatal("expected error for invalid expiryDate, got nil")
	}
}

func TestIdTagInfo_invalidParentIdTag_tooLong(t *testing.T) {
	t.Parallel()

	tooLong := "TOO-LONG-ID-TAG-FOR-CISTRING20-XXXXX"

	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &tooLong,
	}

	_, err := authorizetypes.IdTagInfo(input)
	if err == nil {
		t.Fatal("expected error for invalid parentIdTag, got nil")
	}
}
