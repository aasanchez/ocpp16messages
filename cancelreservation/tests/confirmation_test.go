package cancelreservation_test

import (
	"strings"
	"testing"

	cr "github.com/aasanchez/ocpp16messages/cancelreservation"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.CancelReservationStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.CancelReservationStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.CancelReservationStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.CancelReservationStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Pending (invalid for CancelReservation)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}
