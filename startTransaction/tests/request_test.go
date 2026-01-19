package startTransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/startTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidIdTag        = "RFID-TAG-12345"
	testValidTimestamp    = "2025-01-15T10:30:00Z"
	testConnectorIdOne    = 1
	testMeterStart1000    = 1000
	testReservationId42   = 42
	testValueZero         = 0
	testValueNegativeOne  = -1
	errFieldConnectorId   = "connectorId"
	errFieldMeterStart    = "meterStart"
	errFieldTimestamp     = "timestamp"
	errFieldReservationId = "reservationId"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != testConnectorIdOne {
		t.Errorf(st.ErrorMismatch, testConnectorIdOne, req.ConnectorId.Value())
	}

	if req.IdTag.String() != testValidIdTag {
		t.Errorf(st.ErrorMismatch, testValidIdTag, req.IdTag.String())
	}

	if req.MeterStart.Value() != testMeterStart1000 {
		t.Errorf(st.ErrorMismatch, testMeterStart1000, req.MeterStart.Value())
	}
}

func TestReq_ValidWithReservation(t *testing.T) {
	t.Parallel()

	reservationId := testReservationId42

	req, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: &reservationId,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ReservationId == nil {
		t.Error("Req() ReservationId = nil, want non-nil")
	}

	expectedReservationId := uint16(testReservationId42)
	if req.ReservationId.Value() != expectedReservationId {
		t.Errorf(
			st.ErrorMismatch,
			expectedReservationId,
			req.ReservationId.Value(),
		)
	}
}

func TestReq_ConnectorIdZero(t *testing.T) {
	t.Parallel()

	req, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testValueZero,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != testValueZero {
		t.Errorf(st.ErrorMismatch, testValueZero, req.ConnectorId.Value())
	}
}

func TestReq_ConnectorIdNegative(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testValueNegativeOne,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative connectorId")
	}

	if !strings.Contains(err.Error(), errFieldConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errFieldConnectorId)
	}
}

func TestReq_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         "",
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty idTag")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestReq_IdTagTooLong(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         "RFID-ABC123456789012345", // 23 chars, max is 20
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for IdTag too long")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(
			"Req() error = %v, want 'exceeds maximum length'",
			err,
		)
	}
}

func TestReq_IdTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         "RFID\x00ABC",
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for non-printable chars")
	}

	if !strings.Contains(err.Error(), "non-printable ASCII") {
		t.Errorf(
			"Req() error = %v, want 'non-printable ASCII'",
			err,
		)
	}
}

func TestReq_MeterStartZero(t *testing.T) {
	t.Parallel()

	req, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testValueZero,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.MeterStart.Value() != testValueZero {
		t.Errorf(st.ErrorMismatch, testValueZero, req.MeterStart.Value())
	}
}

func TestReq_MeterStartNegative(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testValueNegativeOne,
		Timestamp:     testValidTimestamp,
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative meterStart")
	}

	if !strings.Contains(err.Error(), errFieldMeterStart) {
		t.Errorf(st.ErrorWantContains, err, errFieldMeterStart)
	}
}

func TestReq_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     "not-a-timestamp",
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(st.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     "",
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(st.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_ReservationIdNegative(t *testing.T) {
	t.Parallel()

	reservationId := testValueNegativeOne

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationId: &reservationId,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative reservationId")
	}

	if !strings.Contains(err.Error(), errFieldReservationId) {
		t.Errorf(st.ErrorWantContains, err, errFieldReservationId)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testValueNegativeOne,
		IdTag:         "",
		MeterStart:    testValueNegativeOne,
		Timestamp:     "invalid",
		ReservationId: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errFieldConnectorId)
	}

	if !strings.Contains(errStr, "idTag") {
		t.Errorf(st.ErrorWantContains, err, "idTag")
	}

	if !strings.Contains(errStr, errFieldMeterStart) {
		t.Errorf(st.ErrorWantContains, err, errFieldMeterStart)
	}

	if !strings.Contains(errStr, errFieldTimestamp) {
		t.Errorf(st.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_MultipleErrorsWithReservation(t *testing.T) {
	t.Parallel()

	reservationId := testValueNegativeOne

	_, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   testValueNegativeOne,
		IdTag:         "",
		MeterStart:    testValueNegativeOne,
		Timestamp:     "invalid",
		ReservationId: &reservationId,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errFieldConnectorId)
	}

	if !strings.Contains(errStr, errFieldReservationId) {
		t.Errorf(st.ErrorWantContains, err, errFieldReservationId)
	}
}
