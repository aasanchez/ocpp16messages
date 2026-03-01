//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewIdTagInfo(f *testing.F) {
	f.Add(st.AuthorizationStatusAccepted.String())
	f.Add(st.AuthorizationStatusBlocked.String())
	f.Add(st.AuthorizationStatusExpired.String())
	f.Add(st.AuthorizationStatusInvalid.String())
	f.Add(st.AuthorizationStatusConcurrentTx.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		idTagInfo, err := st.NewIdTagInfo(st.AuthorizationStatus(status))
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
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
	})
}
