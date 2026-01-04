package types

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestNewIdTagInfo_ValidStatus(t *testing.T) {
	t.Parallel()

	_, err := NewIdTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf("NewIdTagInfo() error = %v, want nil", err)
	}
}

func TestNewIdTagInfo_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := NewIdTagInfo(AuthorizationStatus("InvalidStatus"))
	if err == nil {
		t.Error("NewIdTagInfo() error = nil, want error")
	}
}

func TestNewIdTagInfo_SetsCorrectStatus(t *testing.T) {
	t.Parallel()

	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	if info.Status != AuthorizationStatusAccepted {
		t.Errorf("IdTagInfo.Status = %v, want %v", info.Status, AuthorizationStatusAccepted)
	}
}

func TestNewIdTagInfo_DefaultExpiryDateIsNil(t *testing.T) {
	t.Parallel()

	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	if info.ExpiryDate != nil {
		t.Errorf("IdTagInfo.ExpiryDate = %v, want nil", info.ExpiryDate)
	}
}

func TestNewIdTagInfo_DefaultParentIdTagIsNil(t *testing.T) {
	t.Parallel()

	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	if info.ParentIdTag != nil {
		t.Errorf("IdTagInfo.ParentIdTag = %v, want nil", info.ParentIdTag)
	}
}

func TestIdTagInfo_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate)

	if result.ExpiryDate == nil {
		t.Error("IdTagInfo.ExpiryDate = nil, want non-nil")
	}
}

func TestIdTagInfo_WithParentIdTag(t *testing.T) {
	t.Parallel()

	cistring, _ := st.NewCiString20Type("PARENT123")
	parentTag, _ := NewIdToken(cistring)
	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	result := info.WithParentIdTag(parentTag)

	if result.ParentIdTag == nil {
		t.Error("IdTagInfo.ParentIdTag = nil, want non-nil")
	}
}

func TestIdTagInfo_WithExpiryDate_PreservesStatus(t *testing.T) {
	t.Parallel()

	expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate)

	if result.Status != AuthorizationStatusAccepted {
		t.Errorf("IdTagInfo.Status = %v, want %v", result.Status, AuthorizationStatusAccepted)
	}
}

func TestIdTagInfo_WithParentIdTag_PreservesStatus(t *testing.T) {
	t.Parallel()

	cistring, _ := st.NewCiString20Type("PARENT123")
	parentTag, _ := NewIdToken(cistring)
	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	result := info.WithParentIdTag(parentTag)

	if result.Status != AuthorizationStatusAccepted {
		t.Errorf("IdTagInfo.Status = %v, want %v", result.Status, AuthorizationStatusAccepted)
	}
}
