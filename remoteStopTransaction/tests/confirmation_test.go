package remoteStopTransaction_test

import (
	"strings"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remoteStopTransaction"
	rt "github.com/aasanchez/ocpp16messages/remoteStopTransaction/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.RemoteStopTransactionStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			rt.RemoteStopTransactionStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.RemoteStopTransactionStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			rt.RemoteStopTransactionStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := rst.Conf(rst.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Pending (invalid for RemoteStop)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
