package changeAvailability_test

import (
	"strings"
	"testing"

	ca "github.com/aasanchez/ocpp16messages/changeAvailability"
	mcat "github.com/aasanchez/ocpp16messages/changeAvailability/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errStatus = "status"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != mcat.AvailabilityStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			mcat.AvailabilityStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Rejected"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != mcat.AvailabilityStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			mcat.AvailabilityStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_Scheduled(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Scheduled"})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != mcat.AvailabilityStatusScheduled {
		t.Errorf(
			st.ErrorMismatch,
			mcat.AvailabilityStatusScheduled,
			conf.Status,
		)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: ""})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Unknown(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: "Unknown"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "unknown status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: "accepted"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}

func TestConf_InvalidStatus_Pending(t *testing.T) {
	t.Parallel()

	_, err := ca.Conf(ca.ConfInput{Status: "Pending"})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Pending (invalid for ChangeAvailability)")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(st.ErrorWantContains, err, errStatus)
	}
}
