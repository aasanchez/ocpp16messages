package changeAvailability_test

import (
	"strings"
	"testing"

	ca "github.com/aasanchez/ocpp16messages/messages/changeAvailability"
	mcat "github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
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

	conf, err := ca.Conf(ca.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(errConfUnexpect, err)
	}

	if conf.Status != mcat.AvailabilityStatusAccepted {
		t.Errorf(
			errConfMismatch,
			mcat.AvailabilityStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(errConfUnexpect, err)
	}

	if conf.Status != mcat.AvailabilityStatusRejected {
		t.Errorf(
			errConfMismatch,
			mcat.AvailabilityStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_Scheduled(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Scheduled"})
	if err != nil {
		t.Errorf(errConfUnexpect, err)
	}

	if conf.Status != mcat.AvailabilityStatusScheduled {
		t.Errorf(
			errConfMismatch,
			mcat.AvailabilityStatusScheduled,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(errConfWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(errConfWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(errConfWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(errConfWantNil, "Pending (invalid for ChangeAvailability)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(errConfContains, err, errStatus)
	}
}
