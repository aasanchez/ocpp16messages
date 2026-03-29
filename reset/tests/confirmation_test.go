package reset_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/reset"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := reset.Conf(reset.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ResetStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.ResetStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := reset.Conf(reset.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ResetStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.ResetStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := reset.Conf(reset.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := reset.Conf(reset.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := reset.Conf(reset.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := reset.Conf(reset.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Pending (invalid for Reset)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
