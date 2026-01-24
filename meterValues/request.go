package meterValues

import (
	"errors"
	"fmt"
	"strconv"

	st "github.com/aasanchez/ocpp16messages/types"
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
	MeterValue []st.MeterValueInput
}

// ReqMessage represents an OCPP 1.6 MeterValues.req message.
type ReqMessage struct {
	ConnectorId   st.Integer
	TransactionId *st.Integer
	MeterValue    []st.MeterValue
}

// reqValidation holds validated fields during construction.
type reqValidation struct {
	connectorId   st.Integer
	transactionId *st.Integer
	meterValue    []st.MeterValue
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
			ConnectorId:   st.Integer{},
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
) (st.Integer, []error) {
	intVal, err := st.NewInteger(strconv.Itoa(connectorId))
	if err != nil {
		return st.Integer{}, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "ConnectorId", err),
		)
	}

	return intVal, errs
}

func validateReqTransactionId(
	transactionId int,
	errs []error,
) (*st.Integer, []error) {
	intVal, err := st.NewInteger(strconv.Itoa(transactionId))
	if err != nil {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "TransactionId", err),
		)
	}

	return &intVal, errs
}

const meterValuesLenZero = 0

func validateReqMeterValues(
	meterValues []st.MeterValueInput,
	errs []error,
) ([]st.MeterValue, []error) {
	if len(meterValues) == meterValuesLenZero {
		return nil, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "MeterValue", st.ErrEmptyValue),
		)
	}

	var validValues []st.MeterValue

	for i, mvInput := range meterValues {
		meterValue, err := st.NewMeterValue(mvInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("meterValue[%d]: %w", i, err))
		} else {
			validValues = append(validValues, meterValue)
		}
	}

	return validValues, errs
}
