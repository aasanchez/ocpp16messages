//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	stp "github.com/aasanchez/ocpp16messages/stopTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzStopTransactionConf(f *testing.F) {
	f.Add(false, "", false, "", false, "")
	f.Add(true, st.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(true, "invalid-status", false, "", false, "")
	f.Add(false, "", true, "bad-timestamp", false, "")
	f.Add(false, "", false, "", true, "bad\x01")

	f.Fuzz(func(
		t *testing.T,
		hasStatus bool,
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

		var statusPtr *string
		if hasStatus {
			statusPtr = &status
		}

		var expiryDatePtr *string
		if hasExpiryDate {
			expiryDatePtr = &expiryDate
		}

		var parentIdTagPtr *string
		if hasParentIdTag {
			parentIdTagPtr = &parentIdTag
		}

		conf, err := stp.Conf(stp.ConfInput{
			Status:      statusPtr,
			ExpiryDate:  expiryDatePtr,
			ParentIdTag: parentIdTagPtr,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if !hasStatus {
			if conf.IdTagInfo != nil {
				t.Fatal("IdTagInfo != nil, want nil when Status is omitted")
			}

			return
		}

		if conf.IdTagInfo == nil {
			t.Fatal("IdTagInfo = nil, want non-nil when Status is provided")
		}

		if !conf.IdTagInfo.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.IdTagInfo.Status.String())
		}

		if hasExpiryDate {
			if conf.IdTagInfo.ExpiryDate == nil {
				t.Fatal("ExpiryDate = nil, want non-nil")
			}
			if conf.IdTagInfo.ExpiryDate.Value().Location() != time.UTC {
				t.Fatalf(
					"ExpiryDate location = %v, want UTC",
					conf.IdTagInfo.ExpiryDate.Value().Location(),
				)
			}
		}

		if hasParentIdTag {
			if conf.IdTagInfo.ParentIdTag == nil {
				t.Fatal("ParentIdTag = nil, want non-nil")
			}
		}
	})
}
