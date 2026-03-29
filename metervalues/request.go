package metervalues

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a MeterValues.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The ID of the connector for which meter values are reported.
	// Use 0 for the entire Charge Point.
	ConnectorId int
	// Optional: The transaction ID for which meter values are reported.
	TransactionId *int
	// Required: One or more meter value sets.
	MeterValue []types.MeterValueInput
}

// ReqMessage represents an OCPP 1.6 MeterValues.req message.
type ReqMessage struct {
	ConnectorId   types.Integer
	TransactionId *types.Integer
	MeterValue    []types.MeterValue
}

// reqValidation holds validated fields during construction.
type reqValidation struct {
	connectorId   types.Integer
	transactionId *types.Integer
	meterValue    []types.MeterValue
}

// Req creates a MeterValues.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative
//   - MeterValue is empty or contains invalid entries
//   - TransactionId is provided but invalid
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if errs != nil {
		return ReqMessage{
			ConnectorId:   types.Integer{},
			TransactionId: nil,
			MeterValue:    nil,
		}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorId:   validated.connectorId,
		TransactionId: validated.transactionId,
		MeterValue:    validated.meterValue,
	}, nil
}

func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	validated.connectorId, errs = validateReqConnectorId(
		input.ConnectorId,
		errs,
	)

	if input.TransactionId != nil {
		validated.transactionId, errs = validateReqTransactionId(
			*input.TransactionId,
			errs,
		)
	}

	validated.meterValue, errs = validateReqMeterValues(input.MeterValue, errs)

	return validated, errs
}

func validateReqConnectorId(
	connectorId int,
	errs []error,
) (types.Integer, []error) {
	intVal, err := types.NewInteger(connectorId)
	if err != nil {
		return types.Integer{}, append(
			errs,
			fmt.Errorf(types.ErrorFieldFormat, "ConnectorId", err),
		)
	}

	return intVal, errs
}

func validateReqTransactionId(
	transactionId int,
	errs []error,
) (*types.Integer, []error) {
	intVal, err := types.NewInteger(transactionId)
	if err != nil {
		return nil, append(
			errs,
			fmt.Errorf(types.ErrorFieldFormat, "TransactionId", err),
		)
	}

	return &intVal, errs
}

const metervaluesLenZero = 0

func validateReqMeterValues(
	metervalues []types.MeterValueInput,
	errs []error,
) ([]types.MeterValue, []error) {
	if len(metervalues) == metervaluesLenZero {
		return nil, append(
			errs,
			fmt.Errorf(
				types.ErrorFieldFormat, "MeterValue", types.ErrEmptyValue,
			),
		)
	}

	var validValues []types.MeterValue

	for i, mvInput := range metervalues {
		meterValue, err := types.NewMeterValue(mvInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("meterValue[%d]: %w", i, err))
		} else {
			validValues = append(validValues, meterValue)
		}
	}

	return validValues, errs
}
