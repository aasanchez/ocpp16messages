package remotestarttransaction_test

import (
	"errors"
	"strings"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remotestarttransaction"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	testValidIdTag        = "RFID-TAG-12345"
	testConnectorIdOne    = 1
	testConnectorIdZero   = 0
	testConnectorIdMax    = 65535
	testConnectorIdOver   = 65536
	testConnectorIdNeg    = -1
	errIdTag              = "idTag"
	errConnectorId        = "connectorId"
	errExceedsMaxLength   = "exceeds maximum length"
	errNonPrintableASCII  = "non-printable ASCII"
	fieldNameConnectorId  = "ConnectorId"
	wantConnectorIdNilMsg = "ConnectorId should be nil"
)

func TestReq_Valid_IdTagOnly(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{
		IdTag:       testValidIdTag,
		ConnectorId: nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.IdTag.Value() != testValidIdTag {
		t.Errorf(st.ErrorMismatch, testValidIdTag, req.IdTag.Value())
	}

	if req.ConnectorId != nil {
		t.Errorf(st.ErrorWantNonNil, wantConnectorIdNilMsg)
	}
}

func TestReq_Valid_WithConnectorId(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIdOne

	req, err := rst.Req(rst.ReqInput{
		IdTag:       testValidIdTag,
		ConnectorId: &connectorId,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.IdTag.Value() != testValidIdTag {
		t.Errorf(st.ErrorMismatch, testValidIdTag, req.IdTag.Value())
	}

	if req.ConnectorId == nil {
		t.Errorf(st.ErrorWantNonNil, fieldNameConnectorId)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdOne) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testConnectorIdOne),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_ConnectorIdZero(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIdZero

	req, err := rst.Req(rst.ReqInput{
		IdTag:       testValidIdTag,
		ConnectorId: &connectorId,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Errorf(st.ErrorWantNonNil, fieldNameConnectorId)
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

	connectorId := testConnectorIdMax

	req, err := rst.Req(rst.ReqInput{
		IdTag:       testValidIdTag,
		ConnectorId: &connectorId,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Errorf(st.ErrorWantNonNil, fieldNameConnectorId)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdMax) {
		t.Errorf(
			st.ErrorMismatchValue,
			uint16(testConnectorIdMax),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_EmptyIdTag(t *testing.T) {
	t.Parallel()

	_, err := rst.Req(rst.ReqInput{IdTag: "", ConnectorId: nil})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty idTag")
	}

	if !errors.Is(err, st.ErrEmptyValue) {
		t.Errorf(st.ErrorWrapping, err, st.ErrEmptyValue)
	}
}

func TestReq_IdTagTooLong(t *testing.T) {
	t.Parallel()

	// 23 chars, max is 20
	_, err := rst.Req(rst.ReqInput{
		IdTag:       "RFID-ABC123456789012345",
		ConnectorId: nil,
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

	// Contains null byte
	_, err := rst.Req(rst.ReqInput{IdTag: "RFID\x00ABC", ConnectorId: nil})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "non-printable chars in idTag")
	}

	if !strings.Contains(err.Error(), errNonPrintableASCII) {
		t.Errorf(st.ErrorWantContains, err, errNonPrintableASCII)
	}
}

func TestReq_ConnectorIdNegative(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIdNeg

	_, err := rst.Req(rst.ReqInput{
		IdTag:       testValidIdTag,
		ConnectorId: &connectorId,
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

	connectorId := testConnectorIdOver

	_, err := rst.Req(rst.ReqInput{
		IdTag:       testValidIdTag,
		ConnectorId: &connectorId,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "connectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_MultipleErrors_EmptyIdTagAndInvalidConnectorId(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIdNeg

	_, err := rst.Req(rst.ReqInput{
		IdTag:       "",
		ConnectorId: &connectorId,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty idTag and negative connectorId")
	}

	// Should contain both errors
	if !strings.Contains(err.Error(), errIdTag) {
		t.Errorf(st.ErrorWantContains, err, errIdTag)
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(st.ErrorWantContains, err, errConnectorId)
	}
}
