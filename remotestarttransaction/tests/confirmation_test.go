package remotestarttransaction_test

import (
	"strings"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remotestarttransaction"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.RemoteStartTransactionStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.RemoteStartTransactionStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.RemoteStartTransactionStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.RemoteStartTransactionStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Pending (invalid for RemoteStart)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
