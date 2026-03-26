package unlockconnector

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// connectorIdMinValue is the minimum valid connectorId (must be > 0).
	connectorIdMinValue = 0
)

// ReqInput represents the raw input data for creating an UnlockConnector.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The identifier of the connector to unlock. Must be > 0.
	ConnectorId int
}

// ReqMessage represents an OCPP 1.6 UnlockConnector.req message.
type ReqMessage struct {
	ConnectorId st.Integer
}

// Req creates an UnlockConnector.req message from the given input.
// It validates all fields and returns an error if:
//   - ConnectorId is zero, negative, or exceeds uint16 max value (65535)
//
// Note: ConnectorId must be > 0 because connector 0 refers to the Charge Point
// itself, not a physical connector.
func Req(input ReqInput) (ReqMessage, error) {
	if input.ConnectorId <= connectorIdMinValue {
		return ReqMessage{}, fmt.Errorf("connectorId: %w", st.ErrInvalidValue)
	}

	connectorId, err := st.NewInteger(input.ConnectorId)
	if err != nil {
		return ReqMessage{}, fmt.Errorf("connectorId: %w", err)
	}

	return ReqMessage{
		ConnectorId: connectorId,
	}, nil
}
