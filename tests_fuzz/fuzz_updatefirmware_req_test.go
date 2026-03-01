//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/types"
	uf "github.com/aasanchez/ocpp16messages/updateFirmware"
)

func FuzzUpdateFirmwareReq(f *testing.F) {
	f.Add("https://example.com/fw", "2025-01-15T10:30:00Z", false, 0, false, 0)
	f.Add("", "2025-01-15T10:30:00Z", false, 0, false, 0)
	f.Add("https://example.com/fw", "bad-time", false, 0, false, 0)
	f.Add("https://example.com/fw", "2025-01-15T10:30:00Z", true, -1, false, 0)

	f.Fuzz(func(
		t *testing.T,
		location string,
		retrieveDate string,
		hasRetries bool,
		retries int,
		hasRetryInterval bool,
		retryInterval int,
	) {
		if len(location) > maxFuzzStringLen || len(retrieveDate) > maxFuzzStringLen {
			t.Skip()
		}

		var retriesPtr *int
		if hasRetries {
			retriesPtr = &retries
		}

		var retryIntervalPtr *int
		if hasRetryInterval {
			retryIntervalPtr = &retryInterval
		}

		req, err := uf.Req(uf.ReqInput{
			Location:      location,
			RetrieveDate:  retrieveDate,
			Retries:       retriesPtr,
			RetryInterval: retryIntervalPtr,
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

		if location == "" {
			t.Fatal("Req succeeded with empty Location")
		}

		if req.Location.String() != location {
			t.Fatalf("Location = %q, want %q", req.Location.String(), location)
		}

		if req.RetrieveDate.Value().Location() != time.UTC {
			t.Fatalf(
				"RetrieveDate location = %v, want UTC",
				req.RetrieveDate.Value().Location(),
			)
		}

		if hasRetries {
			if req.Retries == nil {
				t.Fatal("Retries = nil, want non-nil")
			}
			if retries < 0 || retries > math.MaxUint16 {
				t.Fatalf("Req succeeded with retries=%d", retries)
			}
		} else if req.Retries != nil {
			t.Fatal("Retries != nil, want nil")
		}

		if hasRetryInterval {
			if req.RetryInterval == nil {
				t.Fatal("RetryInterval = nil, want non-nil")
			}
			if retryInterval < 0 || retryInterval > math.MaxUint16 {
				t.Fatalf("Req succeeded with retryInterval=%d", retryInterval)
			}
		} else if req.RetryInterval != nil {
			t.Fatal("RetryInterval != nil, want nil")
		}
	})
}
