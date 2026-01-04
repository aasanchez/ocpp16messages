package authorize

import (
	"testing"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestNewResponse(t *testing.T) {
	t.Parallel()

	t.Run("BasicResponse", func(t *testing.T) {
		t.Parallel()

		idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
		resp := ma.NewResponse(idTagInfo)

		if resp.IdTagInfo.Status != mat.AuthorizationStatusAccepted {
			t.Errorf(
				"Response.IdTagInfo.Status = %v, want %v",
				resp.IdTagInfo.Status,
				mat.AuthorizationStatusAccepted,
			)
		}
	})

	t.Run("ResponseWithExpiryDate", func(t *testing.T) {
		t.Parallel()

		expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
		idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
		idTagInfo = idTagInfo.WithExpiryDate(expiryDate)

		resp := ma.NewResponse(idTagInfo)

		if resp.IdTagInfo.ExpiryDate == nil {
			t.Error("Response.IdTagInfo.ExpiryDate = nil, want non-nil")
		}
	})

	t.Run("ResponseWithParentIdTag", func(t *testing.T) {
		t.Parallel()

		cistring, _ := st.NewCiString20Type("PARENT123")
		parentTag, _ := mat.NewIdToken(cistring)
		idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
		idTagInfo = idTagInfo.WithParentIdTag(parentTag)

		resp := ma.NewResponse(idTagInfo)

		if resp.IdTagInfo.ParentIdTag == nil {
			t.Error("Response.IdTagInfo.ParentIdTag = nil, want non-nil")
		}
	})

	t.Run("AllStatuses", func(t *testing.T) {
		t.Parallel()

		statuses := []mat.AuthorizationStatus{
			mat.AuthorizationStatusAccepted,
			mat.AuthorizationStatusBlocked,
			mat.AuthorizationStatusExpired,
			mat.AuthorizationStatusInvalid,
			mat.AuthorizationStatusConcurrentTx,
		}

		for _, status := range statuses {
			idTagInfo, err := mat.NewIdTagInfo(status)
			if err != nil {
				t.Fatalf("NewIdTagInfo(%v) error = %v", status, err)
			}

			resp := ma.NewResponse(idTagInfo)
			if resp.IdTagInfo.Status != status {
				t.Errorf(
					"Response.IdTagInfo.Status = %v, want %v",
					resp.IdTagInfo.Status,
					status,
				)
			}
		}
	})
}
