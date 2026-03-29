package remotestarttransaction

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a
// RemoteStartTransaction.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The identifier to use for starting the transaction.
	IdTag string
	// Optional: The connector on which to start the transaction.
	// If not provided, the Charge Point will choose an available connector.
	ConnectorId *int
}

// ReqMessage represents an OCPP 1.6 RemoteStartTransaction.req message.
type ReqMessage struct {
	IdTag       types.CiString20Type
	ConnectorId *types.Integer
}

// Req creates a RemoteStartTransaction.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - IdTag is empty
//   - IdTag exceeds 20 characters
//   - IdTag contains non-printable ASCII characters
//   - ConnectorId (if provided) is negative or exceeds uint16 max value (65535)
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	idTag, err := types.NewCiString20Type(input.IdTag)
	if err != nil {
		errs = append(errs, fmt.Errorf("idTag: %w", err))
	}

	var connectorId *types.Integer

	if input.ConnectorId != nil {
		connectorId, errs = validateConnectorId(*input.ConnectorId, errs)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		IdTag:       idTag,
		ConnectorId: connectorId,
	}, nil
}

// validateConnectorId validates the connectorId field.
func validateConnectorId(connectorId int, errs []error) (*types.Integer, []error) {
	val, err := types.NewInteger(connectorId)
	if err != nil {
		return nil, append(errs, fmt.Errorf("connectorId: %w", err))
	}

	return &val, errs
}
