package triggermessage_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/triggermessage"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := triggermessage.Conf(triggermessage.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.TriggerMessageStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.TriggerMessageStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := triggermessage.Conf(triggermessage.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.TriggerMessageStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.TriggerMessageStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotImplemented(t *testing.T) {
	t.Parallel()

	conf, err := triggermessage.Conf(triggermessage.ConfInput{
		Status: "NotImplemented",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.TriggerMessageStatusNotImplemented {
		t.Errorf(
			types.ErrorMismatch,
			types.TriggerMessageStatusNotImplemented,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Conf(triggermessage.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Conf(triggermessage.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Conf(triggermessage.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Conf(triggermessage.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Pending (invalid for TriggerMessage)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
