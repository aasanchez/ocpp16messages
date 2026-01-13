package types_test

import (
	"testing"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const fieldStatus = "IdTagInfo.Status"

func TestNewIdTagInfo_ValidStatus(t *testing.T) {
	t.Parallel()

	_, err := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	if err != nil {
		t.Errorf("mat.NewIdTagInfo() error = %v, want nil", err)
	}
}

func TestNewIdTagInfo_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := mat.NewIdTagInfo(mat.AuthorizationStatus("InvalidStatus"))
	if err == nil {
		t.Error("mat.NewIdTagInfo() error = nil, want error")
	}
}

func TestNewIdTagInfo_SetsCorrectStatus(t *testing.T) {
	t.Parallel()

	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	if info.Status != mat.AuthorizationStatusAccepted {
		t.Errorf(
			st.ErrorMethodMismatch,
			fieldStatus,
			info.Status,
			mat.AuthorizationStatusAccepted,
		)
	}
}

func TestNewIdTagInfo_DefaultExpiryDateIsNil(t *testing.T) {
	t.Parallel()

	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	if info.ExpiryDate != nil {
		t.Errorf("IdTagInfo.ExpiryDate = %v, want nil", info.ExpiryDate)
	}
}

func TestNewIdTagInfo_DefaultParentIdTagIsNil(t *testing.T) {
	t.Parallel()

	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	if info.ParentIdTag != nil {
		t.Errorf("IdTagInfo.ParentIdTag = %v, want nil", info.ParentIdTag)
	}
}

func TestIdTagInfo_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate)

	if result.ExpiryDate == nil {
		t.Error("IdTagInfo.ExpiryDate = nil, want non-nil")
	}
}

func TestIdTagInfo_WithParentIdTag(t *testing.T) {
	t.Parallel()

	cistring, _ := st.NewCiString20Type("PARENT123")
	parentTag := mat.NewIdToken(cistring)
	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	result := info.WithParentIdTag(parentTag)

	if result.ParentIdTag == nil {
		t.Error("IdTagInfo.ParentIdTag = nil, want non-nil")
	}
}

func TestIdTagInfo_WithExpiryDate_PreservesStatus(t *testing.T) {
	t.Parallel()

	expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate)

	if result.Status != mat.AuthorizationStatusAccepted {
		t.Errorf(
			st.ErrorMethodMismatch,
			fieldStatus,
			result.Status,
			mat.AuthorizationStatusAccepted,
		)
	}
}

func TestIdTagInfo_WithParentIdTag_PreservesStatus(t *testing.T) {
	t.Parallel()

	cistring, _ := st.NewCiString20Type("PARENT123")
	parentTag := mat.NewIdToken(cistring)
	info, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	result := info.WithParentIdTag(parentTag)

	if result.Status != mat.AuthorizationStatusAccepted {
		t.Errorf(
			st.ErrorMethodMismatch,
			fieldStatus,
			result.Status,
			mat.AuthorizationStatusAccepted,
		)
	}
}
