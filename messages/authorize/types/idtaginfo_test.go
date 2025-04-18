package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestIdTagInfo_validOnlyStatus(t *testing.T) {
	t.Parallel()

	info, err := authorizetypes.IdTagInfo("Accepted")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	if err := info.Validate(); err != nil {
		t.Errorf("expected valid status-only IdTagInfo, got: %v", err)
	}
}

func TestIdTagInfo_withExpiryDate(t *testing.T) {
	t.Parallel()

	info, err := authorizetypes.IdTagInfo("Accepted")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	expiry, err := sharedtypes.DateTime("2027-04-12T14:03:04Z")
	if err != nil {
		t.Fatalf("failed to construct DateTimeType: %v", err)
	}

	info.ExpiryDate = &expiry

	if err := info.Validate(); err != nil {
		t.Errorf("expected valid expiryDate, got: %v", err)
	}
}

func TestIdTagInfo_withZeroExpiryDate_shouldFail(t *testing.T) {
	t.Parallel()

	info, err := authorizetypes.IdTagInfo("Accepted")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	// Create zero-value DateTimeType explicitly
	var zero sharedtypes.DateTimeType
	info.ExpiryDate = &zero

	if err := info.Validate(); err == nil {
		t.Error("expected error for zero ExpiryDate, got nil")
	}
}

func TestIdTagInfo_withValidParentIdTag(t *testing.T) {
	t.Parallel()

	info, err := authorizetypes.IdTagInfo("Accepted")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	parent, err := authorizetypes.IdToken("PARENT123456")
	if err != nil {
		t.Fatalf("failed to create parentIdTag: %v", err)
	}

	info.ParentIdTag = &parent

	if err := info.Validate(); err != nil {
		t.Errorf("expected valid ParentIdTag, got: %v", err)
	}
}

func TestIdTagInfo_withInvalidParentIdTag_shouldFail(t *testing.T) {
	t.Parallel()

	info, err := authorizetypes.IdTagInfo("Accepted")
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	info.ParentIdTag = &authorizetypes.IdTokenType{} // invalid zero value

	if err := info.Validate(); err == nil {
		t.Error("expected error for invalid ParentIdTag, got nil")
	}
}

func TestIdTagInfo_withInvalidStatus_shouldFail(t *testing.T) {
	t.Parallel()

	_, err := authorizetypes.IdTagInfo("WrongStatus")
	if err == nil {
		t.Error("expected constructor to fail with invalid status, got nil")
	}
}

func TestIdTagInfo_validateFailsForZeroStatus(t *testing.T) {
	t.Parallel()

	info := authorizetypes.IdTagInfoType{} // zero-value struct

	if err := info.Validate(); err == nil {
		t.Error("expected validation to fail for zero-value Status, got nil")
	}
}
