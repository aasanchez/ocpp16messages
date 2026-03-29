package firmwarestatusnotification_test

import (
	"errors"
	"testing"

	fsn "github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
	types "github.com/aasanchez/ocpp16types"
)

const (
	statusIdle               = "Idle"
	statusDownloading        = "Downloading"
	statusDownloaded         = "Downloaded"
	statusDownloadFailed     = "DownloadFailed"
	statusInstalling         = "Installing"
	statusInstalled          = "Installed"
	statusInstallationFailed = "InstallationFailed"
)

func TestReq_ValidIdle(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusIdle})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusIdle {
		t.Errorf(types.ErrorMismatch, statusIdle, req.Status.String())
	}
}

func TestReq_ValidDownloading(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusDownloading})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusDownloading {
		t.Errorf(types.ErrorMismatch, statusDownloading, req.Status.String())
	}
}

func TestReq_ValidDownloaded(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusDownloaded})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusDownloaded {
		t.Errorf(types.ErrorMismatch, statusDownloaded, req.Status.String())
	}
}

func TestReq_ValidDownloadFailed(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusDownloadFailed})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusDownloadFailed {
		t.Errorf(types.ErrorMismatch, statusDownloadFailed, req.Status.String())
	}
}

func TestReq_ValidInstalling(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusInstalling})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusInstalling {
		t.Errorf(types.ErrorMismatch, statusInstalling, req.Status.String())
	}
}

func TestReq_ValidInstalled(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusInstalled})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusInstalled {
		t.Errorf(types.ErrorMismatch, statusInstalled, req.Status.String())
	}
}

func TestReq_ValidInstallationFailed(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: statusInstallationFailed})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != statusInstallationFailed {
		t.Errorf(
			types.ErrorMismatch,
			statusInstallationFailed,
			req.Status.String(),
		)
	}
}

func TestReq_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := fsn.Req(fsn.ReqInput{Status: "Invalid"})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestReq_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := fsn.Req(fsn.ReqInput{Status: ""})
	if err == nil {
		t.Error("Req() error = nil, want error for empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestReq_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := fsn.Req(fsn.ReqInput{Status: "idle"})
	if err == nil {
		t.Error("Req() error = nil, want error for lowercase status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}
