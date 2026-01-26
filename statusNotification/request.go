package statusNotification

import (
	"errors"
	"fmt"

	sn "github.com/aasanchez/ocpp16messages/statusNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a StatusNotification.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The id of the connector for which the status is reported.
	// Id 0 (zero) is used for the Charge Point main controller.
	ConnectorId int
	// Required: The error code reported by the Charge Point.
	ErrorCode string
	// Required: The current status of the Charge Point.
	Status string
	// Optional: Additional free format information related to the error.
	Info *string
	// Optional: The time for which the status is reported (RFC3339 format).
	Timestamp *string
	// Optional: Identifies the vendor-specific implementation.
	VendorId *string
	// Optional: Vendor-specific error code.
	VendorErrorCode *string
}

// ReqMessage represents an OCPP 1.6 StatusNotification.req message.
type ReqMessage struct {
	ConnectorId     st.Integer
	ErrorCode       sn.ChargePointErrorCode
	Status          sn.ChargePointStatus
	Info            *st.CiString50Type
	Timestamp       *st.DateTime
	VendorId        *st.CiString255Type
	VendorErrorCode *st.CiString50Type
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	connectorId     st.Integer
	errorCode       sn.ChargePointErrorCode
	status          sn.ChargePointStatus
	info            st.CiString50Type
	timestamp       st.DateTime
	vendorId        st.CiString255Type
	vendorErrorCode st.CiString50Type
}

// Req creates a StatusNotification.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorId is negative or exceeds uint16 max value (65535)
//   - ErrorCode is not a valid ChargePointErrorCode value
//   - Status is not a valid ChargePointStatus value
//   - Info (if provided) exceeds 50 characters or contains invalid chars
//   - Timestamp (if provided) is not a valid RFC3339 date
//   - VendorId (if provided) exceeds 255 characters or contains invalid chars
//   - VendorErrorCode (if provided) exceeds 50 chars or contains invalid chars
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

	// Validate required fields
	validated.connectorId, errs = validateConnectorId(input.ConnectorId, errs)
	validated.errorCode, errs = validateErrorCode(input.ErrorCode, errs)
	validated.status, errs = validateStatus(input.Status, errs)

	// Validate optional fields
	validated, errs = validateOptionalFields(input, validated, errs)

	return validated, errs
}

// validateOptionalFields validates all optional fields in ReqInput.
func validateOptionalFields(
	input ReqInput,
	validated reqValidation,
	errs []error,
) (reqValidation, []error) {
	if input.Info != nil {
		validated.info, errs = validateInfo(*input.Info, errs)
	}

	if input.Timestamp != nil {
		validated.timestamp, errs = validateTimestamp(*input.Timestamp, errs)
	}

	if input.VendorId != nil {
		validated.vendorId, errs = validateVendorId(*input.VendorId, errs)
	}

	if input.VendorErrorCode != nil {
		validated.vendorErrorCode, errs = validateVendorErrorCode(
			*input.VendorErrorCode,
			errs,
		)
	}

	return validated, errs
}

// validateConnectorId validates the connectorId field.
func validateConnectorId(connectorId int, errs []error) (st.Integer, []error) {
	val, err := st.NewIntegerFromInt(connectorId)
	if err != nil {
		return st.Integer{}, append(errs, fmt.Errorf("connectorId: %w", err))
	}

	return val, errs
}

// validateErrorCode validates the errorCode field.
func validateErrorCode(
	errorCode string,
	errs []error,
) (sn.ChargePointErrorCode, []error) {
	code := sn.ChargePointErrorCode(errorCode)

	if !code.IsValid() {
		return "", append(errs, fmt.Errorf("errorCode: %w", st.ErrInvalidValue))
	}

	return code, errs
}

// validateStatus validates the status field.
func validateStatus(
	status string,
	errs []error,
) (sn.ChargePointStatus, []error) {
	chargePointStatus := sn.ChargePointStatus(status)

	if !chargePointStatus.IsValid() {
		return "", append(errs, fmt.Errorf("status: %w", st.ErrInvalidValue))
	}

	return chargePointStatus, errs
}

// validateInfo validates the info field.
func validateInfo(info string, errs []error) (st.CiString50Type, []error) {
	val, err := st.NewCiString50Type(info)
	if err != nil {
		return st.CiString50Type{}, append(errs, fmt.Errorf("info: %w", err))
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

// validateVendorId validates the vendorId field.
func validateVendorId(
	vendorId string,
	errs []error,
) (st.CiString255Type, []error) {
	val, err := st.NewCiString255Type(vendorId)
	if err != nil {
		return st.CiString255Type{}, append(
			errs,
			fmt.Errorf("vendorId: %w", err),
		)
	}

	return val, errs
}

// validateVendorErrorCode validates the vendorErrorCode field.
func validateVendorErrorCode(
	vendorErrorCode string,
	errs []error,
) (st.CiString50Type, []error) {
	val, err := st.NewCiString50Type(vendorErrorCode)
	if err != nil {
		return st.CiString50Type{}, append(
			errs,
			fmt.Errorf("vendorErrorCode: %w", err),
		)
	}

	return val, errs
}

// buildReqMessage constructs the final ReqMessage with validated fields.
func buildReqMessage(input ReqInput, validated reqValidation) ReqMessage {
	msg := ReqMessage{
		ConnectorId:     validated.connectorId,
		ErrorCode:       validated.errorCode,
		Status:          validated.status,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	}

	if input.Info != nil {
		msg.Info = &validated.info
	}

	if input.Timestamp != nil {
		msg.Timestamp = &validated.timestamp
	}

	if input.VendorId != nil {
		msg.VendorId = &validated.vendorId
	}

	if input.VendorErrorCode != nil {
		msg.VendorErrorCode = &validated.vendorErrorCode
	}

	return msg
}
