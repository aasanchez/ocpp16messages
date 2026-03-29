//go:build fuzz

package fuzz

import (
	"errors"
	"strings"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewIdTagInfo(f *testing.F) {
	f.Add(types.AuthorizationStatusAccepted.String())
	f.Add(types.AuthorizationStatusBlocked.String())
	f.Add(types.AuthorizationStatusExpired.String())
	f.Add(types.AuthorizationStatusInvalid.String())
	f.Add(types.AuthorizationStatusConcurrentTx.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		idTagInfo, err := types.NewIdTagInfo(types.AuthorizationStatus(status))
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !idTagInfo.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", idTagInfo.Status.String())
		}

		if idTagInfo.ExpiryDate != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if idTagInfo.ParentIdTag != nil {
			t.Fatal("ParentIdTag != nil, want nil")
		}

		// String() determinism
		s1 := idTagInfo.String()
		s2 := idTagInfo.String()

		if s1 != s2 {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				s1, s2,
			)
		}

		// String() must contain status
		if !strings.Contains(s1, status) {
			t.Fatalf(
				"String() %q missing status %q",
				s1, status,
			)
		}
	})
}
