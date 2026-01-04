package authorize

import (
	"testing"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestNewResponse_BasicResponse(t *testing.T) {
	t.Parallel()

	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.Status != mat.AuthorizationStatusAccepted {
		t.Errorf(
			ma.ErrorResponseStatusMismatch,
			resp.IdTagInfo.Status,
			mat.AuthorizationStatusAccepted,
		)
	}
}

func TestNewResponse_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate, _ := st.NewDateTime("2025-12-31T23:59:59Z")
	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	idTagInfo = idTagInfo.WithExpiryDate(expiryDate)

	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.ExpiryDate == nil {
		t.Error("Response.IdTagInfo.ExpiryDate = nil, want non-nil")
	}
}

func TestNewResponse_WithParentIdTag(t *testing.T) {
	t.Parallel()

	cistring, _ := st.NewCiString20Type("PARENT123")
	parentTag, _ := mat.NewIdToken(cistring)
	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	idTagInfo = idTagInfo.WithParentIdTag(parentTag)

	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.ParentIdTag == nil {
		t.Error("Response.IdTagInfo.ParentIdTag = nil, want non-nil")
	}
}

func TestNewResponse_StatusAccepted(t *testing.T) {
	t.Parallel()

	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusAccepted)
	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.Status != mat.AuthorizationStatusAccepted {
		t.Errorf(
			ma.ErrorResponseStatusMismatch,
			resp.IdTagInfo.Status,
			mat.AuthorizationStatusAccepted,
		)
	}
}

func TestNewResponse_StatusBlocked(t *testing.T) {
	t.Parallel()

	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusBlocked)
	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.Status != mat.AuthorizationStatusBlocked {
		t.Errorf(
			ma.ErrorResponseStatusMismatch,
			resp.IdTagInfo.Status,
			mat.AuthorizationStatusBlocked,
		)
	}
}

func TestNewResponse_StatusExpired(t *testing.T) {
	t.Parallel()

	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusExpired)
	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.Status != mat.AuthorizationStatusExpired {
		t.Errorf(
			ma.ErrorResponseStatusMismatch,
			resp.IdTagInfo.Status,
			mat.AuthorizationStatusExpired,
		)
	}
}

func TestNewResponse_StatusInvalid(t *testing.T) {
	t.Parallel()

	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusInvalid)
	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.Status != mat.AuthorizationStatusInvalid {
		t.Errorf(
			ma.ErrorResponseStatusMismatch,
			resp.IdTagInfo.Status,
			mat.AuthorizationStatusInvalid,
		)
	}
}

func TestNewResponse_StatusConcurrentTx(t *testing.T) {
	t.Parallel()

	idTagInfo, _ := mat.NewIdTagInfo(mat.AuthorizationStatusConcurrentTx)
	resp := ma.NewResponse(idTagInfo)

	if resp.IdTagInfo.Status != mat.AuthorizationStatusConcurrentTx {
		t.Errorf(
			ma.ErrorResponseStatusMismatch,
			resp.IdTagInfo.Status,
			mat.AuthorizationStatusConcurrentTx,
		)
	}
}
