package reserveNow_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/reserveNow"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidIdTag            = "RFID-TAG-12345"
	testValidParentIdTag      = "PARENT-12345"
	testValidExpiryDate       = "2025-01-15T10:00:00Z"
	testReservationIdOne      = 1
	testReservationIdZero     = 0
	testReservationIdMax      = 65535
	testReservationIdOver     = 65536
	testReservationIdNeg      = -1
	testConnectorIdOne        = 1
	testConnectorIdZero       = 0
	testConnectorIdMax        = 65535
	testConnectorIdOver       = 65536
	testConnectorIdNeg        = -1
	errReservationId          = "reservationId"
	errConnectorId            = "connectorId"
	errIdTag                  = "idTag"
	errExpiryDate             = "expiryDate"
	errParentIdTag            = "parentIdTag"
	errExceedsMaxLength       = "exceeds maximum length"
	errNonPrintableASCII      = "non-printable ASCII"
	fieldNameParentIdTag      = "ParentIdTag"
	wantParentIdTagNilMsg     = "ParentIdTag should be nil"
	testIdTagTooLong          = "RFID-ABC123456789012345"
	testParentIdTagTooLong    = "PARENT-1234567890123456"
	testInvalidExpiryDate     = "invalid-date"
	testIdTagWithNullByte     = "RFID\x00ABC"
	testParentTagWithNullByte = "PARENT\x00ABC"
)

func TestReq_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ReservationId.Value() != uint16(testReservationIdOne) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testReservationIdOne),
			req.ReservationId.Value(),
		)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdOne) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testConnectorIdOne),
			req.ConnectorId.Value(),
		)
	}

	if req.IdTag.Value() != testValidIdTag {
		t.Errorf(st.ErrorMismatch, testValidIdTag, req.IdTag.Value())
	}

	if req.ParentIdTag != nil {
		t.Errorf(st.ErrorWantNonNil, wantParentIdTagNilMsg)
	}
}

func TestReq_Valid_WithParentIdTag(t *testing.T) {
	t.Parallel()

	parentIdTag := testValidParentIdTag

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   &parentIdTag,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ParentIdTag == nil {
		t.Errorf(st.ErrorWantNonNil, fieldNameParentIdTag)
	}

	if req.ParentIdTag.Value() != testValidParentIdTag {
		t.Errorf(st.ErrorMismatch, testValidParentIdTag, req.ParentIdTag.Value())
	}
}

func TestReq_Valid_ReservationIdZero(t *testing.T) {
	t.Parallel()

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdZero,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ReservationId.Value() != uint16(testReservationIdZero) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testReservationIdZero),
			req.ReservationId.Value(),
		)
	}
}

func TestReq_Valid_ReservationIdMax(t *testing.T) {
	t.Parallel()

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdMax,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ReservationId.Value() != uint16(testReservationIdMax) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testReservationIdMax),
			req.ReservationId.Value(),
		)
	}
}

func TestReq_Valid_ConnectorIdZero(t *testing.T) {
	t.Parallel()

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdZero,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdZero) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testConnectorIdZero),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_ConnectorIdMax(t *testing.T) {
	t.Parallel()

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdMax,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdMax) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testConnectorIdMax),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_ReservationIdNegative(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdNeg,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative reservationId")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(st.ErrorWantContains, err, errReservationId)
	}
}

func TestReq_ReservationIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOver,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "reservationId exceeds max")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(st.ErrorWantContains, err, errReservationId)
	}
}

func TestReq_ConnectorIdNegative(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdNeg,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative connectorId")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_ConnectorIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOver,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "connectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         "",
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty idTag")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestReq_IdTagTooLong(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testIdTagTooLong,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "IdTag too long")
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_IdTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testIdTagWithNullByte,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "non-printable chars in idTag")
	}

	if !strings.Contains(err.Error(), errNonPrintableASCII) {
		t.Errorf(st.ErrorWantContains, err, errNonPrintableASCII)
	}
}

func TestReq_InvalidExpiryDate(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testInvalidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid expiryDate")
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errExpiryDate)
	}
}

func TestReq_ParentIdTagTooLong(t *testing.T) {
	t.Parallel()

	parentIdTag := testParentIdTagTooLong

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   &parentIdTag,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "ParentIdTag too long")
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(st.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_ParentIdTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	parentIdTag := testParentTagWithNullByte

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdOne,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   &parentIdTag,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "non-printable chars in parentIdTag")
	}

	if !strings.Contains(err.Error(), errNonPrintableASCII) {
		t.Errorf(st.ErrorWantContains, err, errNonPrintableASCII)
	}
}

func TestReq_MultipleErrors_AllFieldsInvalid(t *testing.T) {
	t.Parallel()

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdNeg,
		ConnectorId:   testConnectorIdNeg,
		IdTag:         "",
		ExpiryDate:    testInvalidExpiryDate,
		ParentIdTag:   nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "all fields invalid")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(st.ErrorWantContains, err, errReservationId)
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}

	if !strings.Contains(err.Error(), errIdTag) {
		t.Errorf(st.ErrorWantContains, err, errIdTag)
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(st.ErrorWantContains, err, errExpiryDate)
	}
}

func TestReq_MultipleErrors_WithInvalidParentIdTag(t *testing.T) {
	t.Parallel()

	parentIdTag := testParentIdTagTooLong

	_, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: testReservationIdNeg,
		ConnectorId:   testConnectorIdOne,
		IdTag:         testValidIdTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIdTag:   &parentIdTag,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "reservationId and parentIdTag invalid")
	}

	if !strings.Contains(err.Error(), errReservationId) {
		t.Errorf(st.ErrorWantContains, err, errReservationId)
	}

	if !strings.Contains(err.Error(), errParentIdTag) {
		t.Errorf(st.ErrorWantContains, err, errParentIdTag)
	}
}
