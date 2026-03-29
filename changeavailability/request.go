package changeavailability

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a ChangeAvailability.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	ConnectorId int    // Required: Connector ID (0 = entire Charge Point)
	Type        string // Required: "Inoperative" or "Operative"
}

// ReqMessage represents an OCPP 1.6 ChangeAvailability.req message.
type ReqMessage struct {
	ConnectorId types.Integer
	Type        types.AvailabilityType
}

// Req creates a ChangeAvailability.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - Type is not a valid AvailabilityType value
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	connectorId, err := types.NewInteger(input.ConnectorId)
	if err != nil {
		errs = append(errs, fmt.Errorf("connectorId: %w", err))
	}

	availabilityType := types.AvailabilityType(input.Type)
	if !availabilityType.IsValid() {
		errs = append(errs, fmt.Errorf("type: %w", types.ErrInvalidValue))
	}

	if errs != nil {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorId: connectorId,
		Type:        availabilityType,
	}, nil
}
