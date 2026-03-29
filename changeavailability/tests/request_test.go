package changeavailability_test

import (
	"strings"
	"testing"

	ca "github.com/aasanchez/ocpp16messages/changeavailability"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errConnectorId = "connectorId"
	errType        = "type"

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
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.ConnectorId.Value())
	}

	if req.Type != types.AvailabilityTypeInoperative {
		t.Errorf(
			types.ErrorMismatchValue,
			types.AvailabilityTypeInoperative,
			req.Type,
		)
	}
}

func TestReq_Valid_ConnectorOne_Operative(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valuePositive, Type: typeOperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valuePositive {
		t.Errorf(types.ErrorMismatchValue, valuePositive, req.ConnectorId.Value())
	}

	if req.Type != types.AvailabilityTypeOperative {
		t.Errorf(
			types.ErrorMismatchValue,
			types.AvailabilityTypeOperative,
			req.Type,
		)
	}
}

func TestReq_Valid_MaxConnectorId(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueMaxUint16, Type: typeOperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != valueMaxUint16 {
		t.Errorf(types.ErrorMismatchValue, valueMaxUint16, req.ConnectorId.Value())
	}
}

func TestReq_NegativeConnectorId(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueNegative, Type: typeOperative}

	_, err := ca.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connector ID")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_ExceedsMaxConnectorId(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorId: valueExceedsMax, Type: typeOperative}

	_, err := ca.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connector ID exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_InvalidType_Empty(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valuePositive, Type: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valuePositive, Type: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valuePositive, Type: "operative"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_MultipleErrors_NegativeConnectorAndInvalidType(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorId: valueNegative, Type: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}
