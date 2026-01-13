package changeAvailability_test

import (
	"strings"
	"testing"

	ca "github.com/aasanchez/ocpp16messages/messages/changeAvailability"
	mcat "github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
)

const (
	errConnectorId  = "connectorId"
	errType         = "type"
	errUnexpected   = "Unexpected Error: %v"
	errMismatch     = "Expected %v, got %v"
	errWantNil      = "Req() error = nil, want error for %s"
	errWantContains = "Req() error = %v, want error containing '%s'"

	typeOperative   = "Operative"
	typeInoperative = "Inoperative"

	valueZero       = 0
	valuePositive   = 1
	valueMaxUint16  = 65535
	valueExceedsMax = 65536
	valueNegative   = -1
)

func TestReq_Valid_ConnectorZero_Inoperative(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueZero, Type: typeInoperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(errUnexpected, err)
	}

	if req.ConnectorId.Value() != valueZero {
		t.Errorf(errMismatch, valueZero, req.ConnectorId.Value())
	}

	if req.Type != mcat.AvailabilityTypeInoperative {
		t.Errorf(errMismatch, mcat.AvailabilityTypeInoperative, req.Type)
	}
}

func TestReq_Valid_ConnectorOne_Operative(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valuePositive, Type: typeOperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(errUnexpected, err)
	}

	if req.ConnectorId.Value() != valuePositive {
		t.Errorf(errMismatch, valuePositive, req.ConnectorId.Value())
	}

	if req.Type != mcat.AvailabilityTypeOperative {
		t.Errorf(errMismatch, mcat.AvailabilityTypeOperative, req.Type)
	}
}

func TestReq_Valid_MaxConnectorId(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueMaxUint16, Type: typeOperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(errUnexpected, err)
	}

	if req.ConnectorId.Value() != valueMaxUint16 {
		t.Errorf(errMismatch, valueMaxUint16, req.ConnectorId.Value())
	}
}

func TestReq_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueNegative, Type: typeOperative}

	_, err := ca.Req(input)
	if err == nil {
		t.Errorf(errWantNil, "negative connector ID")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(errWantContains, err, errConnectorId)
	}
}

func TestReq_ExceedsMaxConnectorId(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueExceedsMax, Type: typeOperative}

	_, err := ca.Req(input)
	if err == nil {
		t.Errorf(errWantNil, "connector ID exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(errWantContains, err, errConnectorId)
	}
}

func TestReq_InvalidType_Empty(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valuePositive, Type: ""})
	if err == nil {
		t.Errorf(errWantNil, "empty type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(errWantContains, err, errType)
	}
}

func TestReq_InvalidType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valuePositive, Type: "Unknown"})
	if err == nil {
		t.Errorf(errWantNil, "unknown type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(errWantContains, err, errType)
	}
}

func TestReq_InvalidType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valuePositive, Type: "operative"})
	if err == nil {
		t.Errorf(errWantNil, "lowercase type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(errWantContains, err, errType)
	}
}

func TestReq_MultipleErrors_NegativeConnectorAndInvalidType(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valueNegative, Type: "Unknown"})
	if err == nil {
		t.Errorf(errWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(errWantContains, err, errConnectorId)
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(errWantContains, err, errType)
	}
}
