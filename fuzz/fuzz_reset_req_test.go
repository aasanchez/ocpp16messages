//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/reset"
	rt "github.com/aasanchez/ocpp16messages/reset/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzResetReq(f *testing.F) {
	f.Add(rt.ResetTypeHard.String())
	f.Add(rt.ResetTypeSoft.String())
	f.Add("bad-type")

	f.Fuzz(func(t *testing.T, resetType string) {
		if len(resetType) > maxFuzzStringLen {
			t.Skip()
		}

		req, err := reset.Req(reset.ReqInput{
			Type: resetType,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !req.Type.IsValid() {
			t.Fatalf("Type = %q, want valid", req.Type.String())
		}
		if req.Type.String() != resetType {
			t.Fatalf("Type = %q, want %q", req.Type.String(), resetType)
		}
	})
}
