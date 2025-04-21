package authorizetypes_test

import (
	"fmt"
	"testing"
	"time"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

const errCreateIdTagInfo = "unexpected error creating IdTagInfo: %v"

func TestIdTagInfoValidMinimal(t *testing.T) {
	t.Parallel()

	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
	if err != nil {
		t.Fatalf(errCreateIdTagInfo, err)
	}

	if err := info.Validate(); err != nil {
		t.Errorf("expected valid IdTagInfo, got validation error: %v", err)
	}
}

func TestIdTagInfoWithExpiryDate(t *testing.T) {
	t.Parallel()

	expiry := time.Now().Add(24 * time.Hour).UTC()

	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
	if err != nil {
		t.Fatalf(errCreateIdTagInfo, err)
	}

	info.ExpiryDate = &expiry

	if err := info.Validate(); err != nil {
		t.Errorf("expected IdTagInfo with expiry date to be valid, got error: %v", err)
	}
}

func TestIdTagInfoWithInvalidStatus(t *testing.T) {
	t.Parallel()

	info := authorizetypes.IdTagInfoType{
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

	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
	if err != nil {
		t.Fatalf(errCreateIdTagInfo, err)
	}

	info.ExpiryDate = &zeroTime

	if err := info.Validate(); err == nil {
		t.Error("expected error for empty expiry date, got nil")
	}
}

func TestIdTagInfoWithValidParentIdTag(t *testing.T) {
	t.Parallel()

	parent, err := authorizetypes.IdToken("PARENT123456")
	if err != nil {
		t.Fatalf("unexpected error creating parentIdTag: %v", err)
	}

	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
	if err != nil {
		t.Fatalf(errCreateIdTagInfo, err)
	}

	info.ParentIdTag = &parent

	if err := info.Validate(); err != nil {
		t.Errorf("expected IdTagInfo with valid parentIdTag to pass, got error: %v", err)
	}
}

func TestIdTagInfoInvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := authorizetypes.IdTagInfo("INVALID")
	if err == nil {
		t.Error("expected error for invalid AuthorizationStatus, got nil")
	}
}

func TestIdTagInfoStringMethod(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	parent, _ := authorizetypes.IdToken("GROUPID123")

	info := authorizetypes.IdTagInfoType{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  &now,
		ParentIdTag: &parent,
	}

	expected := fmt.Sprintf(
		"{status:%s, expiryDate:%s, parentIdTag:%s}",
		info.Status,
		now.Format(time.RFC3339),
		parent.String(),
	)

	if info.String() != expected {
		t.Errorf("unexpected String() output:\nwant: %s\ngot : %s", expected, info.String())
	}
}
