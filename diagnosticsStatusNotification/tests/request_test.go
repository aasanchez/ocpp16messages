package diagnosticsStatusNotification_test

import (
	"errors"
	"testing"

	dsn "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	statusIdle         = "Idle"
	statusUploaded     = "Uploaded"
	statusUploadFailed = "UploadFailed"
	statusUploading    = "Uploading"
)

func TestReq_ValidIdle(t *testing.T) {
	t.Parallel()

	req, err := dsn.Req(dsn.ReqInput{Status: statusIdle})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusIdle {
		t.Errorf(st.ErrorMismatch, statusIdle, req.Status.String())
	}
}

func TestReq_ValidUploaded(t *testing.T) {
	t.Parallel()

	req, err := dsn.Req(dsn.ReqInput{Status: statusUploaded})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusUploaded {
		t.Errorf(st.ErrorMismatch, statusUploaded, req.Status.String())
	}
}

func TestReq_ValidUploadFailed(t *testing.T) {
	t.Parallel()

	req, err := dsn.Req(dsn.ReqInput{Status: statusUploadFailed})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusUploadFailed {
		t.Errorf(st.ErrorMismatch, statusUploadFailed, req.Status.String())
	}
}

func TestReq_ValidUploading(t *testing.T) {
	t.Parallel()

	req, err := dsn.Req(dsn.ReqInput{Status: statusUploading})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusUploading {
		t.Errorf(st.ErrorMismatch, statusUploading, req.Status.String())
	}
}

func TestReq_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := dsn.Req(dsn.ReqInput{Status: "Invalid"})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestReq_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := dsn.Req(dsn.ReqInput{Status: ""})
	if err == nil {
		t.Error("Req() error = nil, want error for empty status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestReq_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := dsn.Req(dsn.ReqInput{Status: "idle"})
	if err == nil {
		t.Error("Req() error = nil, want error for lowercase status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}
