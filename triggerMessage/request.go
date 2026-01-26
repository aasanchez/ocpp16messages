package triggerMessage

import (
	"errors"
	"fmt"

	tt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a TriggerMessage.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The type of message to trigger.
	RequestedMessage string
	// Optional: The id of the connector for which the message applies.
	// If absent, applies to the Charge Point as a whole.
	ConnectorId *int
}

// ReqMessage represents an OCPP 1.6 TriggerMessage.req message.
type ReqMessage struct {
	RequestedMessage tt.MessageTrigger
	ConnectorId      *st.Integer
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	requestedMessage tt.MessageTrigger
	connectorId      st.Integer
}

// Req creates a TriggerMessage.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - RequestedMessage is not a valid MessageTrigger value
//   - ConnectorId (if provided) is negative or exceeds uint16 max value (65535)
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return buildReqMessage(input, validated), nil
}

// validateReqInput validates all fields in ReqInput and returns validated
// values along with any errors.
func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	// Validate required field
	validated.requestedMessage, errs = validateRequestedMessage(
		input.RequestedMessage,
		errs,
	)

	// Validate optional field
	if input.ConnectorId != nil {
		validated.connectorId, errs = validateConnectorId(
			*input.ConnectorId,
			errs,
		)
	}

	return validated, errs
}

// validateRequestedMessage validates the requestedMessage field.
func validateRequestedMessage(
	requestedMessage string,
	errs []error,
) (tt.MessageTrigger, []error) {
	messageTrigger := tt.MessageTrigger(requestedMessage)

	if !messageTrigger.IsValid() {
		return "", append(
			errs,
			fmt.Errorf("requestedMessage: %w", st.ErrInvalidValue),
		)
	}

	return messageTrigger, errs
}

// validateConnectorId validates the connectorId field.
func validateConnectorId(connectorId int, errs []error) (st.Integer, []error) {
	val, err := st.NewIntegerFromInt(connectorId)
	if err != nil {
		return st.Integer{}, append(errs, fmt.Errorf("connectorId: %w", err))
	}

	return val, errs
}

// buildReqMessage constructs the final ReqMessage with validated fields.
func buildReqMessage(input ReqInput, validated reqValidation) ReqMessage {
	msg := ReqMessage{
		RequestedMessage: validated.requestedMessage,
		ConnectorId:      nil,
	}

	if input.ConnectorId != nil {
		msg.ConnectorId = &validated.connectorId
	}

	return msg
}
