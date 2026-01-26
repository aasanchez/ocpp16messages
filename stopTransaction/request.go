package stopTransaction

import (
	"errors"
	"fmt"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	stt "github.com/aasanchez/ocpp16messages/stopTransaction/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a StopTransaction.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The transaction ID of the transaction to stop.
	TransactionId int
	// Optional: The identifier that was used to stop the transaction.
	// May be omitted when the Charge Point itself stops the transaction.
	IdTag *string
	// Required: Energy meter reading at the end of the transaction in Wh.
	MeterStop int
	// Required: Timestamp of when the transaction was stopped.
	Timestamp string
	// Optional: The reason for stopping the transaction.
	// May be omitted if the transaction ended normally (Local).
	Reason *string
	// Optional: Transaction-related meter values.
	TransactionData []mt.MeterValueInput
}

// ReqMessage represents an OCPP 1.6 StopTransaction.req message.
type ReqMessage struct {
	TransactionId   st.Integer
	IdTag           *st.IdToken
	MeterStop       st.Integer
	Timestamp       st.DateTime
	Reason          *stt.Reason
	TransactionData []mt.MeterValue
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	transactionId   st.Integer
	idTag           *st.IdToken
	meterStop       st.Integer
	timestamp       st.DateTime
	reason          *stt.Reason
	transactionData []mt.MeterValue
}

// Req creates a StopTransaction.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - TransactionId is negative or exceeds uint16 max value (65535)
//   - IdTag (if provided) exceeds 20 characters or contains non-printable ASCII
//   - MeterStop is negative or exceeds uint16 max value (65535)
//   - Timestamp is not a valid RFC3339 formatted date
//   - Reason (if provided) is not a valid Reason value
//   - TransactionData (if provided) contains invalid meter values
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if len(errs) > errCountZero {
		return ReqMessage{
			TransactionId:   st.Integer{},
			IdTag:           nil,
			MeterStop:       st.Integer{},
			Timestamp:       st.DateTime{},
			Reason:          nil,
			TransactionData: nil,
		}, errors.Join(errs...)
	}

	return ReqMessage{
		TransactionId:   validated.transactionId,
		IdTag:           validated.idTag,
		MeterStop:       validated.meterStop,
		Timestamp:       validated.timestamp,
		Reason:          validated.reason,
		TransactionData: validated.transactionData,
	}, nil
}

// validateReqInput validates all fields in ReqInput.
func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	validated.transactionId, errs = validateTransactionId(
		input.TransactionId,
		errs,
	)
	validated.meterStop, errs = validateMeterStop(input.MeterStop, errs)
	validated.timestamp, errs = validateTimestamp(input.Timestamp, errs)

	if input.IdTag != nil {
		validated.idTag, errs = validateIdTag(*input.IdTag, errs)
	}

	if input.Reason != nil {
		validated.reason, errs = validateReason(*input.Reason, errs)
	}

	if len(input.TransactionData) > errCountZero {
		validated.transactionData, errs = validateTransactionData(
			input.TransactionData,
			errs,
		)
	}

	return validated, errs
}

// validateTransactionId validates the transactionId field.
func validateTransactionId(
	transactionId int,
	errs []error,
) (st.Integer, []error) {
	val, err := st.NewInteger(transactionId)
	if err != nil {
		return st.Integer{}, append(errs, fmt.Errorf("transactionId: %w", err))
	}

	return val, errs
}

// validateIdTag validates the idTag field.
func validateIdTag(idTag string, errs []error) (*st.IdToken, []error) {
	ciStr, err := st.NewCiString20Type(idTag)
	if err != nil {
		return nil, append(errs, fmt.Errorf("idTag: %w", err))
	}

	token := st.NewIdToken(ciStr)

	return &token, errs
}

// validateMeterStop validates the meterStop field.
func validateMeterStop(meterStop int, errs []error) (st.Integer, []error) {
	val, err := st.NewInteger(meterStop)
	if err != nil {
		return st.Integer{}, append(errs, fmt.Errorf("meterStop: %w", err))
	}

	return val, errs
}

// validateTimestamp validates the timestamp field.
func validateTimestamp(timestamp string, errs []error) (st.DateTime, []error) {
	val, err := st.NewDateTime(timestamp)
	if err != nil {
		return st.DateTime{}, append(errs, fmt.Errorf("timestamp: %w", err))
	}

	return val, errs
}

// validateReason validates the reason field.
func validateReason(reason string, errs []error) (*stt.Reason, []error) {
	reasonVal := stt.Reason(reason)
	if !reasonVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf("reason: %w: %s", st.ErrInvalidValue, reason),
		)
	}

	return &reasonVal, errs
}

// validateTransactionData validates the transactionData field.
func validateTransactionData(
	transactionData []mt.MeterValueInput,
	errs []error,
) ([]mt.MeterValue, []error) {
	const transactionDataEmpty = 0

	if transactionData == nil {
		return nil, errs
	}

	var validValues []mt.MeterValue

	for i, mvInput := range transactionData {
		meterValue, err := mt.NewMeterValue(mvInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("transactionData[%d]: %w", i, err))
		} else {
			validValues = append(validValues, meterValue)
		}
	}

	if len(transactionData) == transactionDataEmpty {
		return []mt.MeterValue{}, errs
	}

	return validValues, errs
}
