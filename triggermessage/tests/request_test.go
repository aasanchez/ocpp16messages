package triggermessage_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/triggermessage"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errRequestedMessage = "requestedMessage"
	errConnectorId      = "connectorId"
	fieldConnectorId    = "ConnectorId"
	connectorIdZero     = 0
	connectorIdOne      = 1
	connectorIdNegative = -1
	connectorIdMax      = 65535
	connectorIdOverflow = 65536
)

func TestReq_Valid_BootNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "BootNotification",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerBootNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerBootNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_DiagnosticsStatusNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "DiagnosticsStatusNotification",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage !=
		types.MessageTriggerDiagnosticsStatusNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerDiagnosticsStatusNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_FirmwareStatusNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "FirmwareStatusNotification",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerFirmwareStatusNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerFirmwareStatusNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_Heartbeat(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerHeartbeat {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerHeartbeat,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_MeterValues(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "MeterValues",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerMeterValues {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerMeterValues,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_StatusNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StatusNotification",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerStatusNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerStatusNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_WithConnectorIdZero(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdZero

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StatusNotification",
		ConnectorId:      &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Errorf(types.ErrorWantNonNil, fieldConnectorId)
	}

	if req.ConnectorId.Value() != uint16(connectorIdZero) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(connectorIdZero),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_WithConnectorIdOne(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdOne

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "MeterValues",
		ConnectorId:      &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Errorf(types.ErrorWantNonNil, fieldConnectorId)
	}

	if req.ConnectorId.Value() != uint16(connectorIdOne) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(connectorIdOne),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_WithConnectorIdMax(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdMax

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId == nil {
		t.Errorf(types.ErrorWantNonNil, fieldConnectorId)
	}

	if req.ConnectorId.Value() != uint16(connectorIdMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(connectorIdMax),
			req.ConnectorId.Value(),
		)
	}
}

func TestReq_Valid_WithoutConnectorId(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId != nil {
		t.Errorf("ConnectorId = %v, want nil", req.ConnectorId)
	}
}

func TestReq_EmptyRequestedMessage(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "",
		ConnectorId:      nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty requestedMessage")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidRequestedMessage_Unknown(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Unknown",
		ConnectorId:      nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown requestedMessage")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidRequestedMessage_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "heartbeat",
		ConnectorId:      nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase requestedMessage")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidRequestedMessage_StartTransaction(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StartTransaction",
		ConnectorId:      nil,
	})
	if err == nil {
		t.Errorf(
			types.ErrorWantNil,
			"StartTransaction (not valid per OCPP 1.6)",
		)
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidConnectorId_Negative(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdNegative

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connectorId")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_InvalidConnectorId_Overflow(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdOverflow

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId overflow")
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}

func TestReq_MultipleErrors_InvalidMessageAndConnectorId(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdNegative

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Unknown",
		ConnectorId:      &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid message and connectorId")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}

	if !strings.Contains(err.Error(), errConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errConnectorId)
	}
}
