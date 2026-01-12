package cancelReservation_test

import (
	"strings"
	"testing"

	cr "github.com/aasanchez/ocpp16messages/messages/cancelReservation"
	mct "github.com/aasanchez/ocpp16messages/messages/cancelReservation/types"
)

const (
	errStatus       = "status"
	errConfUnexpect = "Unexpected Error: %v"
	errConfMismatch = "Expected %q, got %q"
	errConfWantNil  = "Conf() error = nil, want error for %s"
	errConfContains = "Conf() error = %v, want error containing '%s'"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(errConfUnexpect, err)
	}

	if conf.Status != mct.CancelReservationStatusAccepted {
		t.Errorf(
			errConfMismatch,
			mct.CancelReservationStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(errConfUnexpect, err)
	}

	if conf.Status != mct.CancelReservationStatusRejected {
		t.Errorf(
			errConfMismatch,
			mct.CancelReservationStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(errConfWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(errConfWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(errConfWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(errConfWantNil, "Pending (invalid for CancelReservation)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}
