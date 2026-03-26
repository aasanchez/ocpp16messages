//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"unicode/utf8"

	"github.com/aasanchez/ocpp16messages/changeconfiguration"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzChangeConfigurationReq(f *testing.F) {
	f.Add("MeterValueSampleInterval", "10")
	f.Add("", "10")
	f.Add("Key", "")
	f.Add("123456789012345678901234567890123456789012345678901", "10") // 51 chars
	f.Add("\x00", "10")

	f.Fuzz(func(t *testing.T, key string, value string) {
		if len(key) > maxFuzzStringLen || len(value) > maxFuzzStringLen {
			t.Skip()
		}

		if (!utf8.ValidString(key) && len(key) > 256) ||
			(!utf8.ValidString(value) && len(value) > 256) {
			t.Skip()
		}

		req, err := changeconfiguration.Req(changeconfiguration.ReqInput{
			Key:   key,
			Value: value,
		})
		if err != nil {
			if !errors.Is(err, st.ErrEmptyValue) && !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if key == "" {
			t.Fatal("Req succeeded with empty Key")
		}
		if value == "" {
			t.Fatal("Req succeeded with empty Value")
		}
		if len(key) > st.CiString50Max {
			t.Fatalf("Req succeeded with Key len=%d", len(key))
		}
		if len(value) > st.CiString500Max {
			t.Fatalf("Req succeeded with Value len=%d", len(value))
		}

		if req.Key.String() != key {
			t.Fatalf("Key = %q, want %q", req.Key.String(), key)
		}
		if req.Value.String() != value {
			t.Fatalf("Value = %q, want %q", req.Value.String(), value)
		}
	})
}
