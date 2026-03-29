//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	dt "github.com/aasanchez/ocpp16messages/datatransfer"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzDataTransferConf(f *testing.F) {
	f.Add(types.DataTransferStatusAccepted.String(), false, "")
	f.Add("invalid-status", false, "")
	f.Add(types.DataTransferStatusRejected.String(), true, "payload")

	f.Fuzz(func(t *testing.T, status string, hasData bool, data string) {
		if len(status) > maxFuzzStringLen || len(data) > maxFuzzStringLen {
			t.Skip()
		}

		var dataPtr *string
		if hasData {
			dataPtr = &data
		}

		conf, err := dt.Conf(dt.ConfInput{
			Status: status,
			Data:   dataPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !conf.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.Status.String())
		}

		if hasData {
			if conf.Data == nil {
				t.Fatal("Data = nil, want non-nil")
			}
			if *conf.Data != data {
				t.Fatalf("Data = %q, want %q", *conf.Data, data)
			}
		} else if conf.Data != nil {
			t.Fatal("Data != nil, want nil")
		}
	})
}
