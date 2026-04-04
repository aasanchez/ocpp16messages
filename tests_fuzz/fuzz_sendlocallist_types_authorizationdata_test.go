//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewAuthorizationData(f *testing.F) {
	f.Add("", false, "", false, "", false, "")
	f.Add("RFID-ABC123", false, "", false, "", false, "")
	f.Add(
		"RFID-ABC123",
		true,
		types.AuthorizationStatusAccepted.String(),
		true,
		"2025-01-15T10:30:00Z",
		true,
		"RFID-PARENT",
	)
	f.Add("RFID-ABC123", true, "invalid-status", false, "", false, "")
	f.Add("bad\x01", false, "", false, "", false, "")

	f.Fuzz(func(
		t *testing.T,
		idTag string,
		hasIdTagInfo bool,
		status string,
		hasExpiryDate bool,
		expiryDate string,
		hasParentIdTag bool,
		parentIdTag string,
	) {
		if len(idTag) > maxFuzzStringLen ||
			len(status) > maxFuzzStringLen ||
			len(expiryDate) > maxFuzzStringLen ||
			len(parentIdTag) > maxFuzzStringLen {
			t.Skip()
		}

		var idTagInfoPtr *types.IdTagInfoInput
		if hasIdTagInfo {
			var expiryDatePtr *string
			if hasExpiryDate {
				expiryDatePtr = &expiryDate
			}

			var parentIdTagPtr *string
			if hasParentIdTag {
				parentIdTagPtr = &parentIdTag
			}

			idTagInfoPtr = &types.IdTagInfoInput{
				Status:      status,
				ExpiryDate:  expiryDatePtr,
				ParentIdTag: parentIdTagPtr,
			}
		}

		authData, err := types.NewAuthorizationData(types.AuthorizationDataInput{
			IdTag:     idTag,
			IdTagInfo: idTagInfoPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrEmptyValue) && !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if idTag == "" {
			t.Fatal("NewAuthorizationData succeeded with empty IdTag")
		}

		if authData.IdTag().String() != idTag {
			t.Fatalf("IdTag = %q, want %q", authData.IdTag().String(), idTag)
		}

		if !hasIdTagInfo {
			if authData.IdTagInfo() != nil {
				t.Fatal("IdTagInfo != nil, want nil")
			}

			return
		}

		if authData.IdTagInfo() == nil {
			t.Fatal("IdTagInfo = nil, want non-nil")
		}

		if !authData.IdTagInfo().Status().IsValid() {
			t.Fatalf("Status = %q, want valid", authData.IdTagInfo().Status().String())
		}

		if hasExpiryDate {
			if authData.IdTagInfo().ExpiryDate() == nil {
				t.Fatal("ExpiryDate = nil, want non-nil")
			}
			if authData.IdTagInfo().ExpiryDate().Value().Location() != time.UTC {
				t.Fatalf(
					"ExpiryDate location = %v, want UTC",
					authData.IdTagInfo().ExpiryDate().Value().Location(),
				)
			}
		} else if authData.IdTagInfo().ExpiryDate() != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if hasParentIdTag {
			if authData.IdTagInfo().ParentIdTag() == nil {
				t.Fatal("ParentIdTag = nil, want non-nil")
			}
			if authData.IdTagInfo().ParentIdTag().String() != parentIdTag {
				t.Fatalf(
					"ParentIdTag = %q, want %q",
					authData.IdTagInfo().ParentIdTag().String(),
					parentIdTag,
				)
			}
		} else if authData.IdTagInfo().ParentIdTag() != nil {
			t.Fatal("ParentIdTag != nil, want nil")
		}
	})
}
