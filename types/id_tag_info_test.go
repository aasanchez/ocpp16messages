package types

import (
	"fmt"
	"testing"
	"time"
)

func TestIdTagInfoValidMinimal(t *testing.T) {
	t.Parallel()

	info, err := IdTagInfo(Accepted)
	if err != nil {
		t.Fatalf("unexpected error creating IdTagInfo: %v", err)
	}

	if err := info.Validate(); err != nil {
		t.Errorf("expected valid IdTagInfo, got validation error: %v", err)
	}
}

func TestIdTagInfoWithExpiryDate(t *testing.T) {
	t.Parallel()

	expiry := time.Now().Add(24 * time.Hour).UTC()

	info, err := IdTagInfo(Accepted)
	if err != nil {
		t.Fatalf("unexpected error creating IdTagInfo: %v", err)
	}

	info.ExpiryDate = &expiry

	if err := info.Validate(); err != nil {
		t.Errorf("expected IdTagInfo with expiry date to be valid, got error: %v", err)
	}
}

func TestIdTagInfoWithInvalidStatus(t *testing.T) {
	t.Parallel()

	info := IdTagInfoType{
		Status:      "UNKNOWN_STATUS",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	err := info.Validate()
	if err == nil {
		t.Error("expected error for invalid status, got nil")
	}
}

func TestIdTagInfoWithEmptyExpiryDate(t *testing.T) {
	t.Parallel()

	var zeroTime time.Time

	info, err := IdTagInfo(Accepted)
	if err != nil {
		t.Fatalf("unexpected error creating IdTagInfo: %v", err)
	}

	info.ExpiryDate = &zeroTime

	if err := info.Validate(); err == nil {
		t.Error("expected error for empty expiry date, got nil")
	}
}

func TestIdTagInfoWithValidParentIdTag(t *testing.T) {
	t.Parallel()

	parent, err := IdToken("PARENT123456")
	if err != nil {
		t.Fatalf("unexpected error creating parentIdTag: %v", err)
	}

	info, err := IdTagInfo(Accepted)
	if err != nil {
		t.Fatalf("unexpected error creating IdTagInfo: %v", err)
	}

	info.ParentIdTag = &parent

	if err := info.Validate(); err != nil {
		t.Errorf("expected IdTagInfo with valid parentIdTag to pass, got error: %v", err)
	}
}

func TestIdTagInfoWithInvalidParentIdTag(t *testing.T) {
	t.Parallel()

	invalidParent := IdTokenType{
		value: CiString20Type{
			inner: ciString{
				Value:  "",
				MaxLen: maxLenCiString20,
			},
		},
	}

	info := IdTagInfoType{
		Status:      Accepted,
		ParentIdTag: &invalidParent,
		ExpiryDate:  nil,
	}

	if err := info.Validate(); err == nil {
		t.Error("expected validation to fail due to invalid ParentIdTag, got nil")
	}
}

func TestIdTagInfoInvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := IdTagInfo("INVALID")
	if err == nil {
		t.Error("expected error for invalid AuthorizationStatus, got nil")
	}
}

func TestIdTagInfoStringMethod(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	parent, _ := IdToken("GROUPID123")

	info := IdTagInfoType{
		Status:      Accepted,
		ExpiryDate:  &now,
		ParentIdTag: &parent,
	}

	expected := fmt.Sprintf(
		"{status=%s, expiryDate=%s, parentIdTag=%s}",
		info.Status,
		now.Format(time.RFC3339),
		parent.String(),
	)

	if info.String() != expected {
		t.Errorf("unexpected String() output:\nwant: %s\ngot : %s", expected, info.String())
	}
}
