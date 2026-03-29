package changeconfiguration_test

import (
	"strings"
	"testing"

	cc "github.com/aasanchez/ocpp16messages/changeconfiguration"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ConfigurationStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.ConfigurationStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ConfigurationStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.ConfigurationStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_RebootRequired(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "RebootRequired"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ConfigurationStatusRebootRequired {
		t.Errorf(
			types.ErrorMismatch,
			types.ConfigurationStatusRebootRequired,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotSupported(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "NotSupported"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ConfigurationStatusNotSupported {
		t.Errorf(
			types.ErrorMismatch,
			types.ConfigurationStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Scheduled(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: "Scheduled"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Scheduled (invalid for ChangeConfiguration)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
