package changeAvailability

import (
	"errors"
	"fmt"
	"strconv"

	mcat "github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ReqInput represents the raw input data for creating a ChangeAvailability.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	ConnectorId int    // Required: Connector ID (0 = entire Charge Point)
	Type        string // Required: "Inoperative" or "Operative"
}

// ReqMessage represents an OCPP 1.6 ChangeAvailability.req message.
type ReqMessage struct {
	ConnectorId st.Integer
	Type        mcat.AvailabilityType
}

// Req creates a ChangeAvailability.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - Type is not a valid AvailabilityType value
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	connectorId, err := st.NewInteger(strconv.Itoa(input.ConnectorId))
	if err != nil {
		errs = append(errs, fmt.Errorf("connectorId: %w", err))
	}

	availabilityType := mcat.AvailabilityType(input.Type)
	if !availabilityType.IsValid() {
		errs = append(errs, fmt.Errorf("type: %w", st.ErrInvalidValue))
	}

	if errs != nil {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorId: connectorId,
		Type:        availabilityType,
	}, nil
}
