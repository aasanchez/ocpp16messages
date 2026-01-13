package cancelReservation_test

import (
	"strings"
	"testing"

	cr "github.com/aasanchez/ocpp16messages/messages/cancelReservation"
	mct "github.com/aasanchez/ocpp16messages/messages/cancelReservation/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != mct.CancelReservationStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			mct.CancelReservationStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != mct.CancelReservationStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			mct.CancelReservationStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := cr.Conf(cr.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Pending (invalid for CancelReservation)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
