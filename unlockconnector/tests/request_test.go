package unlockconnector_test

import (
	"strings"
	"testing"

	uc "github.com/aasanchez/ocpp16messages/unlockconnector"
	types "github.com/aasanchez/ocpp16types"
)

const (
	testConnectorIdValid = 1
	testConnectorIdTwo   = 2
	testConnectorIdMax   = 65535
	testConnectorIdZero  = 0
	testConnectorIdOver  = 65536
	testConnectorIdNeg   = -1
	errConnectorId       = "connectorId"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorId: testConnectorIdValid})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdValid) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIdValid),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_ConnectorTwo(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorId: testConnectorIdTwo})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdTwo) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIdTwo),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_Max(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorId: testConnectorIdMax})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != uint16(testConnectorIdMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIdMax),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Invalid_ConnectorIdZero(t *testing.T) {
	t.Parallel()

	_, err := uc.Req(uc.ReqInput{ConnectorId: testConnectorIdZero})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId zero")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_ConnectorIdNegative(t *testing.T) {
	t.Parallel()

	_, err := uc.Req(uc.ReqInput{ConnectorId: testConnectorIdNeg})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connectorId")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_Invalid_ConnectorIdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := uc.Req(uc.ReqInput{ConnectorId: testConnectorIdOver})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}
