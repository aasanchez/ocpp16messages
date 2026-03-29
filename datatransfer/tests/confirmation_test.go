package datatransfer_test

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/datatransfer"
	types "github.com/aasanchez/ocpp16types"
)

const (
	statusAccepted         = "Accepted"
	statusRejected         = "Rejected"
	statusUnknownMessageId = "UnknownMessageId"
	statusUnknownVendor    = "UnknownVendor"
	confTestData           = `{"response": "data"}`
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: statusAccepted,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusAccepted {
		t.Errorf(types.ErrorMismatch, statusAccepted, conf.Status.String())
	}

	if conf.Data != nil {
		t.Error("Conf() Data != nil, want nil")
	}
}

func TestConf_ValidRejected(t *testing.T) {
	t.Parallel()

	conf, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: statusRejected,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusRejected {
		t.Errorf(types.ErrorMismatch, statusRejected, conf.Status.String())
	}
}

func TestConf_ValidUnknownMessageId(t *testing.T) {
	t.Parallel()

	conf, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: statusUnknownMessageId,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusUnknownMessageId {
		t.Errorf(
			types.ErrorMismatch,
			statusUnknownMessageId,
			conf.Status.String(),
		)
	}
}

func TestConf_ValidUnknownVendor(t *testing.T) {
	t.Parallel()

	conf, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: statusUnknownVendor,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusUnknownVendor {
		t.Errorf(types.ErrorMismatch, statusUnknownVendor, conf.Status.String())
	}
}

func TestConf_ValidWithData(t *testing.T) {
	t.Parallel()

	data := confTestData

	conf, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: statusAccepted,
		Data:   &data,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Data == nil {
		t.Error("Conf() Data = nil, want non-nil")

		return
	}

	if *conf.Data != confTestData {
		t.Errorf(types.ErrorMismatch, confTestData, *conf.Data)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: "Invalid",
		Data:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: "",
		Data:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestConf_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: "accepted",
		Data:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for lowercase status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}
