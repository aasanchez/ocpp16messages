package triggerMessage_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/triggerMessage"
	tt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != tt.TriggerMessageStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			tt.TriggerMessageStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "Rejected",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != tt.TriggerMessageStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			tt.TriggerMessageStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_NotImplemented(t *testing.T) {
	t.Parallel()

	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "NotImplemented",
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != tt.TriggerMessageStatusNotImplemented {
		t.Errorf(
			st.ErrorMismatch,
			tt.TriggerMessageStatusNotImplemented,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := triggerMessage.Conf(triggerMessage.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := triggerMessage.Conf(triggerMessage.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := triggerMessage.Conf(triggerMessage.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := triggerMessage.Conf(triggerMessage.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Pending (invalid for TriggerMessage)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
