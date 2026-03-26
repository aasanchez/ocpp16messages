//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"unicode/utf8"

	"github.com/aasanchez/ocpp16messages/getconfiguration"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzGetConfigurationReq(f *testing.F) {
	f.Add(uint8(0), "", "")                                          // nil keys
	f.Add(uint8(1), "", "")                                          // empty keys
	f.Add(uint8(2), "MeterValueSampleInterval", "")                  // one key
	f.Add(uint8(3), "MeterValueSampleInterval", "HeartbeatInterval") // two keys
	f.Add(uint8(2), "", "")                                          // empty key
	f.Add(uint8(2), "\x00", "")                                      // bad ASCII

	f.Fuzz(func(t *testing.T, keysMode uint8, key0 string, key1 string) {
		if len(key0) > maxFuzzStringLen || len(key1) > maxFuzzStringLen {
			t.Skip()
		}

		if (!utf8.ValidString(key0) && len(key0) > 256) ||
			(!utf8.ValidString(key1) && len(key1) > 256) {
			t.Skip()
		}

		var keys []string

		switch keysMode % 4 {
		case 0:
			keys = nil
		case 1:
			keys = []string{}
		case 2:
			keys = []string{key0}
		default:
			keys = []string{key0, key1}
		}

		req, err := getconfiguration.Req(getconfiguration.ReqInput{
			Key: keys,
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

		switch keysMode % 4 {
		case 0, 1:
			if req.Key != nil {
				t.Fatal("Key != nil, want nil")
			}
		case 2:
			if req.Key == nil {
				t.Fatal("Key = nil, want non-nil")
			}
			if len(req.Key) != 1 {
				t.Fatalf("len(Key) = %d, want 1", len(req.Key))
			}
			if req.Key[0].String() != key0 {
				t.Fatalf("Key[0] = %q, want %q", req.Key[0].String(), key0)
			}
		default:
			if req.Key == nil {
				t.Fatal("Key = nil, want non-nil")
			}
			if len(req.Key) != 2 {
				t.Fatalf("len(Key) = %d, want 2", len(req.Key))
			}
			if req.Key[0].String() != key0 {
				t.Fatalf("Key[0] = %q, want %q", req.Key[0].String(), key0)
			}
			if req.Key[1].String() != key1 {
				t.Fatalf("Key[1] = %q, want %q", req.Key[1].String(), key1)
			}
		}
	})
}
