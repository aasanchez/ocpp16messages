package dataTransfer_test

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/dataTransfer"
	st "github.com/aasanchez/ocpp16messages/types"
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

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: statusAccepted,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusAccepted {
		t.Errorf(st.ErrorMismatch, statusAccepted, conf.Status.String())
	}

	if conf.Data != nil {
		t.Error("Conf() Data != nil, want nil")
	}
}

func TestConf_ValidRejected(t *testing.T) {
	t.Parallel()

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: statusRejected,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusRejected {
		t.Errorf(st.ErrorMismatch, statusRejected, conf.Status.String())
	}
}

func TestConf_ValidUnknownMessageId(t *testing.T) {
	t.Parallel()

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: statusUnknownMessageId,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusUnknownMessageId {
		t.Errorf(
			st.ErrorMismatch,
			statusUnknownMessageId,
			conf.Status.String(),
		)
	}
}

func TestConf_ValidUnknownVendor(t *testing.T) {
	t.Parallel()

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: statusUnknownVendor,
		Data:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status.String() != statusUnknownVendor {
		t.Errorf(st.ErrorMismatch, statusUnknownVendor, conf.Status.String())
	}
}

func TestConf_ValidWithData(t *testing.T) {
	t.Parallel()

	data := confTestData

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: statusAccepted,
		Data:   &data,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Data == nil {
		t.Error("Conf() Data = nil, want non-nil")

		return
	}

	if *conf.Data != confTestData {
		t.Errorf(st.ErrorMismatch, confTestData, *conf.Data)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "Invalid",
		Data:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "",
		Data:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestConf_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "accepted",
		Data:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for lowercase status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}
