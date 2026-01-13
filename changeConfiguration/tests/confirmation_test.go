package changeConfiguration_test

import (
	"strings"
	"testing"

	cc "github.com/aasanchez/ocpp16messages/changeConfiguration"
	cct "github.com/aasanchez/ocpp16messages/changeConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != cct.ConfigurationStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			cct.ConfigurationStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != cct.ConfigurationStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			cct.ConfigurationStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_RebootRequired(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "RebootRequired"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != cct.ConfigurationStatusRebootRequired {
		t.Errorf(
			st.ErrorMismatch,
			cct.ConfigurationStatusRebootRequired,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotSupported(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "NotSupported"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != cct.ConfigurationStatusNotSupported {
		t.Errorf(
			st.ErrorMismatch,
			cct.ConfigurationStatusNotSupported,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Scheduled(t *testing.T) {
	t.Parallel()

	_, err := cc.Conf(cc.ConfInput{Status: "Scheduled"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Scheduled (invalid for ChangeConfiguration)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
