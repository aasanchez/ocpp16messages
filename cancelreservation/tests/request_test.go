package cancelreservation_test

import (
	"strings"
	"testing"

	cr "github.com/aasanchez/ocpp16messages/cancelreservation"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errReservationId = "reservationId"

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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationId.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.ReservationId.Value())
	}
}

func TestReq_Valid_PositiveValue(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationId: valuePositive})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationId.Value() != valuePositive {
		t.Errorf(
			types.ErrorMismatchValue,
			valuePositive,
			req.ReservationId.Value(),
		)
	}
}

func TestReq_Valid_MaxValue(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationId: valueMaxUint16})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationId.Value() != valueMaxUint16 {
		t.Errorf(
			types.ErrorMismatchValue,
			valueMaxUint16,
			req.ReservationId.Value(),
		)
	}
}

func TestReq_NegativeValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationId: valueNegative})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative reservation ID")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(types.ErrorWantContains, err, errReservationId)
	}
}

func TestReq_ExceedsMaxValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationId: valueExceedsMax})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "reservation ID exceeds max")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(types.ErrorWantContains, err, errReservationId)
	}
}

func TestReq_LargeNegativeValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationId: valueLargeNegativ})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "large negative reservation ID")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(types.ErrorWantContains, err, errReservationId)
	}
}
