package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/types"
)

const fieldStatus = "IdTagInfo.Status"

func TestNewIdTagInfo_ValidStatus(t *testing.T) {
	t.Parallel()

	_, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if err != nil {
		t.Errorf("types.NewIdTagInfo() error = %v, want nil", err)
	}
}

func TestNewIdTagInfo_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := types.NewIdTagInfo(types.AuthorizationStatus("InvalidStatus"))
	if err == nil {
		t.Error("types.NewIdTagInfo() error = nil, want error")
	}
}

func TestNewIdTagInfo_SetsCorrectStatus(t *testing.T) {
	t.Parallel()

	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if info.Status != types.AuthorizationStatusAccepted {
		t.Errorf(
			types.ErrorMethodMismatch,
			fieldStatus,
			info.Status,
			types.AuthorizationStatusAccepted,
		)
	}
}

func TestNewIdTagInfo_DefaultExpiryDateIsNil(t *testing.T) {
	t.Parallel()

	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if info.ExpiryDate != nil {
		t.Errorf("IdTagInfo.ExpiryDate = %v, want nil", info.ExpiryDate)
	}
}

func TestNewIdTagInfo_DefaultParentIdTagIsNil(t *testing.T) {
	t.Parallel()

	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	if info.ParentIdTag != nil {
		t.Errorf("IdTagInfo.ParentIdTag = %v, want nil", info.ParentIdTag)
	}
}

func TestIdTagInfo_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate, _ := types.NewDateTime("2025-12-31T23:59:59Z")
	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate)

	if result.ExpiryDate == nil {
		t.Error("IdTagInfo.ExpiryDate = nil, want non-nil")
	}
}

func TestIdTagInfo_WithParentIdTag(t *testing.T) {
	t.Parallel()

	cistring, _ := types.NewCiString20Type("PARENT123")
	parentTag := types.NewIdToken(cistring)
	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	result := info.WithParentIdTag(parentTag)

	if result.ParentIdTag == nil {
		t.Error("IdTagInfo.ParentIdTag = nil, want non-nil")
	}
}

func TestIdTagInfo_WithExpiryDate_PreservesStatus(t *testing.T) {
	t.Parallel()

	expiryDate, _ := types.NewDateTime("2025-12-31T23:59:59Z")
	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate)

	if result.Status != types.AuthorizationStatusAccepted {
		t.Errorf(
			types.ErrorMethodMismatch,
			fieldStatus,
			result.Status,
			types.AuthorizationStatusAccepted,
		)
	}
}

func TestIdTagInfo_WithParentIdTag_PreservesStatus(t *testing.T) {
	t.Parallel()

	cistring, _ := types.NewCiString20Type("PARENT123")
	parentTag := types.NewIdToken(cistring)
	info, _ := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
	result := info.WithParentIdTag(parentTag)

	if result.Status != types.AuthorizationStatusAccepted {
		t.Errorf(
			types.ErrorMethodMismatch,
			fieldStatus,
			result.Status,
			types.AuthorizationStatusAccepted,
		)
	}
}
