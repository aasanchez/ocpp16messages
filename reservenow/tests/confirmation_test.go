package reservenow_test

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/reservenow"
	rt "github.com/aasanchez/ocpp16messages/reservenow/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testStatusAccepted    = "Accepted"
	testStatusFaulted     = "Faulted"
	testStatusOccupied    = "Occupied"
	testStatusRejected    = "Rejected"
	testStatusUnavailable = "Unavailable"
	testStatusInvalid     = "Unknown"
	testStatusEmpty       = ""
	testStatusLowercase   = "accepted"
)

func TestConf_Valid_Accepted(t *testing.T) {
	t.Parallel()

	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusAccepted,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.ReservationStatusAccepted {
		t.Errorf(
			st.ErrorMismatch,
			rt.ReservationStatusAccepted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Faulted(t *testing.T) {
	t.Parallel()

	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusFaulted,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.ReservationStatusFaulted {
		t.Errorf(
			st.ErrorMismatch,
			rt.ReservationStatusFaulted,
			conf.Status,
		)
	}
}

func TestConf_Valid_Occupied(t *testing.T) {
	t.Parallel()

	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusOccupied,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.ReservationStatusOccupied {
		t.Errorf(
			st.ErrorMismatch,
			rt.ReservationStatusOccupied,
			conf.Status,
		)
	}
}

func TestConf_Valid_Rejected(t *testing.T) {
	t.Parallel()

	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusRejected,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.ReservationStatusRejected {
		t.Errorf(
			st.ErrorMismatch,
			rt.ReservationStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_Unavailable(t *testing.T) {
	t.Parallel()

	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusUnavailable,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if conf.Status != rt.ReservationStatusUnavailable {
		t.Errorf(
			st.ErrorMismatch,
			rt.ReservationStatusUnavailable,
			conf.Status,
		)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusInvalid,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusEmpty,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}

func TestConf_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusLowercase,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "lowercase status")
	}

	if !errors.Is(err, st.ErrInvalidValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrInvalidValue)
	}
}
