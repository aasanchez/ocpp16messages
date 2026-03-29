package starttransaction

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a StartTransaction.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The connector on which the transaction started.
	// A value of 0 indicates that the transaction started on an unspecified
	// connector.
	ConnectorId int
	// Required: The identifier that was used to authorize this transaction.
	IdTag string
	// Required: Energy meter reading at the start of the transaction in Wh.
	MeterStart int
	// Required: Timestamp of the start of the transaction.
	Timestamp string
	// Optional: If the transaction is started because of a reservation, this
	// contains the reservation ID.
	ReservationId *int
}

// ReqMessage represents an OCPP 1.6 StartTransaction.req message.
type ReqMessage struct {
	ConnectorId   types.Integer
	IdTag         types.IdToken
	MeterStart    types.Integer
	Timestamp     types.DateTime
	ReservationId *types.Integer
}

// Req creates a StartTransaction.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - IdTag is empty, exceeds 20 characters, or contains non-printable ASCII
//   - MeterStart is negative or exceeds uint16 max value (65535)
//   - Timestamp is not a valid RFC3339 formatted date
//   - ReservationId (if provided) is negative or exceeds uint16 max (65535)
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	connectorId, errs := validateConnectorId(input.ConnectorId, errs)
	idTag, errs := validateIdTag(input.IdTag, errs)
	meterStart, errs := validateMeterStart(input.MeterStart, errs)
	timestamp, errs := validateTimestamp(input.Timestamp, errs)

	var reservationId *types.Integer

	if input.ReservationId != nil {
		reservationId, errs = validateReservationId(*input.ReservationId, errs)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorId:   connectorId,
		IdTag:         idTag,
		MeterStart:    meterStart,
		Timestamp:     timestamp,
		ReservationId: reservationId,
	}, nil
}

// validateConnectorId validates the connectorId field.
func validateConnectorId(
	connectorId int,
	errs []error,
) (types.Integer, []error) {
	val, err := types.NewInteger(connectorId)
	if err != nil {
		return types.Integer{}, append(errs, fmt.Errorf("connectorId: %w", err))
	}

	return val, errs
}

// validateIdTag validates the idTag field.
func validateIdTag(idTag string, errs []error) (types.IdToken, []error) {
	ciStr, err := types.NewCiString20Type(idTag)
	if err != nil {
		return types.IdToken{}, append(errs, fmt.Errorf("idTag: %w", err))
	}

	return types.NewIdToken(ciStr), errs
}

// validateMeterStart validates the meterStart field.
func validateMeterStart(meterStart int, errs []error) (types.Integer, []error) {
	val, err := types.NewInteger(meterStart)
	if err != nil {
		return types.Integer{}, append(errs, fmt.Errorf("meterStart: %w", err))
	}

	return val, errs
}

// validateTimestamp validates the timestamp field.
func validateTimestamp(
	timestamp string,
	errs []error,
) (types.DateTime, []error) {
	val, err := types.NewDateTime(timestamp)
	if err != nil {
		return types.DateTime{}, append(errs, fmt.Errorf("timestamp: %w", err))
	}

	return val, errs
}

// validateReservationId validates the reservationId field.
func validateReservationId(
	reservationId int,
	errs []error,
) (*types.Integer, []error) {
	val, err := types.NewInteger(reservationId)
	if err != nil {
		return nil, append(errs, fmt.Errorf("reservationId: %w", err))
	}

	return &val, errs
}
