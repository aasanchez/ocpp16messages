//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	"github.com/aasanchez/ocpp16messages/authorize"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzAuthorizeConf(f *testing.F) {
	f.Add(types.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), true, "2025-01-15T10:30:00Z", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), false, "", true, "PARENT-123")
	f.Add(types.AuthorizationStatusInvalid.String(), false, "", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), true, "bad-time", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), false, "", true, "")

	f.Fuzz(func(
		t *testing.T,
		status string,
		hasExpiryDate bool,
		expiryDate string,
		hasParentIdTag bool,
		parentIdTag string,
	) {
		if len(status) > maxFuzzStringLen ||
			len(expiryDate) > maxFuzzStringLen ||
			len(parentIdTag) > maxFuzzStringLen {
			t.Skip()
		}

		var expiryDatePtr *string
		if hasExpiryDate {
			expiryDatePtr = &expiryDate
		}

		var parentIdTagPtr *string
		if hasParentIdTag {
			parentIdTagPtr = &parentIdTag
		}

		conf, err := authorize.Conf(authorize.ConfInput{
			Status:      status,
			ExpiryDate:  expiryDatePtr,
			ParentIdTag: parentIdTagPtr,
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

		if !conf.IdTagInfo.Status().IsValid() {
			t.Fatalf("Status = %q, want valid", conf.IdTagInfo.Status().String())
		}
		if conf.IdTagInfo.Status().String() != status {
			t.Fatalf("Status = %q, want %q", conf.IdTagInfo.Status().String(), status)
		}

		if hasExpiryDate {
			if conf.IdTagInfo.ExpiryDate() == nil {
				t.Fatal("ExpiryDate = nil, want non-nil")
			}
			if conf.IdTagInfo.ExpiryDate().Value().Location() != time.UTC {
				t.Fatalf(
					"ExpiryDate location = %v, want UTC",
					conf.IdTagInfo.ExpiryDate().Value().Location(),
				)
			}
		} else if conf.IdTagInfo.ExpiryDate() != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if hasParentIdTag {
			if conf.IdTagInfo.ParentIdTag() == nil {
				t.Fatal("ParentIdTag = nil, want non-nil")
			}
			if parentIdTag == "" {
				t.Fatal("Conf succeeded with empty ParentIdTag")
			}
			if conf.IdTagInfo.ParentIdTag().String() != parentIdTag {
				t.Fatalf(
					"ParentIdTag = %q, want %q",
					conf.IdTagInfo.ParentIdTag().String(),
					parentIdTag,
				)
			}
		} else if conf.IdTagInfo.ParentIdTag() != nil {
			t.Fatal("ParentIdTag != nil, want nil")
		}
	})
}
