//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"unicode/utf8"

	"github.com/aasanchez/ocpp16messages/authorize"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzAuthorizeReq(f *testing.F) {
	f.Add("RFID-ABC123")
	f.Add("")
	f.Add("123456789012345678901") // 21 chars
	f.Add("\x00")

	f.Fuzz(func(t *testing.T, idTag string) {
		if len(idTag) > maxFuzzStringLen {
			t.Skip()
		}

		// Avoid extremely expensive rune iteration on invalid UTF-8.
		if !utf8.ValidString(idTag) && len(idTag) > 256 {
			t.Skip()
		}

		req, err := authorize.Req(authorize.ReqInput{
			IdTag: idTag,
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
			t.Fatal("Req succeeded with empty IdTag")
		}
		if len(idTag) > types.CiString20Max {
			t.Fatalf("Req succeeded with IdTag len=%d", len(idTag))
		}

		for _, r := range idTag {
			if r < 32 || r > 126 {
				t.Fatalf("Req succeeded with non-printable ASCII rune=%U", r)
			}
		}

		if req.IdTag.String() != idTag {
			t.Fatalf("IdTag = %q, want %q", req.IdTag.String(), idTag)
		}
	})
}
