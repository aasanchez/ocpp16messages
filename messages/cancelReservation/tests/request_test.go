package cancelReservation_test

import (
	"strings"
	"testing"

	cr "github.com/aasanchez/ocpp16messages/messages/cancelReservation"
)

const (
	errReservationId = "reservationId"
	errUnexpected    = "Unexpected Error: %v"
	errMismatch      = "Expected %v, got %v"
	errWantNil       = "Req() error = nil, want error for %s"
	errWantContains  = "Req() error = %v, want error containing '%s'"

	valueZero         = 0
	valuePositive     = 123
	valueMaxUint16    = 65535
	valueExceedsMax   = 65536
	valueNegative     = -1
	valueLargeNegativ = -65536
)

func TestReq_Valid_Zero(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationId: valueZero})
	if err != nil {
		t.Errorf(errUnexpected, err)
	}

	if req.ReservationId.Value() != valueZero {
		t.Errorf(errMismatch, valueZero, req.ReservationId.Value())
	}
}

func TestReq_Valid_PositiveValue(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationId: valuePositive})
	if err != nil {
		t.Errorf(errUnexpected, err)
	}

	if req.ReservationId.Value() != valuePositive {
		t.Errorf(errMismatch, valuePositive, req.ReservationId.Value())
	}
}

func TestReq_Valid_MaxValue(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationId: valueMaxUint16})
	if err != nil {
		t.Errorf(errUnexpected, err)
	}

	if req.ReservationId.Value() != valueMaxUint16 {
		t.Errorf(errMismatch, valueMaxUint16, req.ReservationId.Value())
	}
}

func TestReq_NegativeValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationId: valueNegative})
	if err == nil {
		t.Errorf(errWantNil, "negative reservation ID")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(errWantContains, err, errReservationId)
	}
}

func TestReq_ExceedsMaxValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationId: valueExceedsMax})
	if err == nil {
		t.Errorf(errWantNil, "reservation ID exceeds max")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(errWantContains, err, errReservationId)
	}
}

func TestReq_LargeNegativeValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationId: valueLargeNegativ})
	if err == nil {
		t.Errorf(errWantNil, "large negative reservation ID")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(errWantContains, err, errReservationId)
	}
}
