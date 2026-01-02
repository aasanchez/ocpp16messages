package types

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestNewIdTagInfo(t *testing.T) {
	t.Parallel()

	t.Run("ValidStatus", func(t *testing.T) {
		t.Parallel()

		info, err := NewIdTagInfo(AuthorizationStatusAccepted)
		if err != nil {
			t.Fatalf("NewIdTagInfo() error = %v, want nil", err)
		}

		if info.Status != AuthorizationStatusAccepted {
			t.Errorf(
				"IdTagInfo.Status = %v, want %v",
				info.Status,
				AuthorizationStatusAccepted,
			)
		}

		if info.ExpiryDate != nil {
			t.Errorf("IdTagInfo.ExpiryDate = %v, want nil", info.ExpiryDate)
		}

		if info.ParentIdTag != nil {
			t.Errorf(
				"IdTagInfo.ParentIdTag = %v, want nil",
				info.ParentIdTag,
			)
		}
	})

	t.Run("InvalidStatus", func(t *testing.T) {
		t.Parallel()

		_, err := NewIdTagInfo(AuthorizationStatus("InvalidStatus"))
		if err == nil {
			t.Error("NewIdTagInfo() error = nil, want error")
		}
	})
}

func TestIdTagInfo_Chaining(t *testing.T) {
	t.Parallel()

	expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
	cistring, _ := st.NewCiString20Type("PARENT123")
	parentTag, _ := NewIdToken(cistring)

	info, _ := NewIdTagInfo(AuthorizationStatusAccepted)
	result := info.WithExpiryDate(expiryDate).WithParentIdTag(parentTag)

	if result.ExpiryDate == nil {
		t.Error("IdTagInfo.ExpiryDate = nil, want non-nil")
	}

	if result.ParentIdTag == nil {
		t.Error("IdTagInfo.ParentIdTag = nil, want non-nil")
	}

	if result.Status != AuthorizationStatusAccepted {
		t.Errorf(
			"IdTagInfo.Status = %v, want %v",
			result.Status,
			AuthorizationStatusAccepted,
		)
	}
}
