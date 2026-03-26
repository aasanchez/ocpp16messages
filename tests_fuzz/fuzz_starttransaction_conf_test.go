//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	stt "github.com/aasanchez/ocpp16messages/starttransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzStartTransactionConf(f *testing.F) {
	f.Add(1, st.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(-1, st.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(1, "invalid-status", false, "", false, "")
	f.Add(1, st.AuthorizationStatusAccepted.String(), true, "bad-timestamp", false, "")
	f.Add(1, st.AuthorizationStatusAccepted.String(), false, "", true, "")

	f.Fuzz(func(
		t *testing.T,
		transactionId int,
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

		conf, err := stt.Conf(stt.ConfInput{
			TransactionId: transactionId,
			Status:        status,
			ExpiryDate:    expiryDatePtr,
			ParentIdTag:   parentIdTagPtr,
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

		if transactionId < 0 || transactionId > math.MaxUint16 {
			t.Fatalf("Conf succeeded with transactionId=%d", transactionId)
		}

		if got := conf.TransactionId.Value(); got != uint16(transactionId) {
			t.Fatalf("TransactionId = %d, want %d", got, transactionId)
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
		} else if conf.IdTagInfo.ExpiryDate != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if hasParentIdTag {
			if conf.IdTagInfo.ParentIdTag == nil {
				t.Fatal("ParentIdTag = nil, want non-nil")
			}
		} else if conf.IdTagInfo.ParentIdTag != nil {
			t.Fatal("ParentIdTag != nil, want nil")
		}
	})
}
