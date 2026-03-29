package reservenow_test

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/reservenow"
	types "github.com/aasanchez/ocpp16types"
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ReservationStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.ReservationStatusAccepted,
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ReservationStatusFaulted {
		t.Errorf(
			types.ErrorMismatch,
			types.ReservationStatusFaulted,
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ReservationStatusOccupied {
		t.Errorf(
			types.ErrorMismatch,
			types.ReservationStatusOccupied,
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ReservationStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.ReservationStatusRejected,
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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.ReservationStatusUnavailable {
		t.Errorf(
			types.ErrorMismatch,
			types.ReservationStatusUnavailable,
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
		t.Errorf(types.ErrorWantNil, "invalid status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusEmpty,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestConf_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Conf(reservenow.ConfInput{
		Status: testStatusLowercase,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}
