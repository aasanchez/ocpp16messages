//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	gd "github.com/aasanchez/ocpp16messages/getdiagnostics"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzGetDiagnosticsReq(f *testing.F) {
	f.Add("https://example.com/diag", false, 0, false, 0, false, "", false, "")
	f.Add("", false, 0, false, 0, false, "", false, "")
	f.Add("https://example.com/diag", true, -1, false, 0, false, "", false, "")
	f.Add("https://example.com/diag", false, 0, false, 0, true, "bad-time", false, "")
	f.Add("https://example.com/diag", false, 0, false, 0, false, "", true, "2025-01-15T10:30:00Z")

	f.Fuzz(func(
		t *testing.T,
		location string,
		hasRetries bool,
		retries int,
		hasRetryInterval bool,
		retryInterval int,
		hasStartTime bool,
		startTime string,
		hasStopTime bool,
		stopTime string,
	) {
		if len(location) > maxFuzzStringLen ||
			len(startTime) > maxFuzzStringLen ||
			len(stopTime) > maxFuzzStringLen {
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

		var startTimePtr *string
		if hasStartTime {
			startTimePtr = &startTime
		}

		var stopTimePtr *string
		if hasStopTime {
			stopTimePtr = &stopTime
		}

		req, err := gd.Req(gd.ReqInput{
			Location:      location,
			Retries:       retriesPtr,
			RetryInterval: retryIntervalPtr,
			StartTime:     startTimePtr,
			StopTime:      stopTimePtr,
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

		if hasRetries {
			if req.Retries == nil {
				t.Fatal("Retries = nil, want non-nil")
			}
			if retries < 0 || retries > math.MaxUint16 {
				t.Fatalf("Req succeeded with retries=%d", retries)
			}
			if got := req.Retries.Value(); got != uint16(retries) {
				t.Fatalf("Retries = %d, want %d", got, retries)
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
			if got := req.RetryInterval.Value(); got != uint16(retryInterval) {
				t.Fatalf("RetryInterval = %d, want %d", got, retryInterval)
			}
		} else if req.RetryInterval != nil {
			t.Fatal("RetryInterval != nil, want nil")
		}

		if hasStartTime {
			if req.StartTime == nil {
				t.Fatal("StartTime = nil, want non-nil")
			}
			if req.StartTime.Value().Location() != time.UTC {
				t.Fatalf(
					"StartTime location = %v, want UTC",
					req.StartTime.Value().Location(),
				)
			}
		} else if req.StartTime != nil {
			t.Fatal("StartTime != nil, want nil")
		}

		if hasStopTime {
			if req.StopTime == nil {
				t.Fatal("StopTime = nil, want non-nil")
			}
			if req.StopTime.Value().Location() != time.UTC {
				t.Fatalf(
					"StopTime location = %v, want UTC",
					req.StopTime.Value().Location(),
				)
			}
		} else if req.StopTime != nil {
			t.Fatal("StopTime != nil, want nil")
		}
	})
}
