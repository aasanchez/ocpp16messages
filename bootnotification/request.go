package bootnotification

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a BootNotification.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required fields
	ChargePointVendor string // Vendor of the ChargePoint (max 20 chars)
	ChargePointModel  string // Model of the ChargePoint (max 20 chars)

	// Optional fields
	ChargePointSerialNumber *string // Serial number of the ChargePoint (max 25)
	ChargeBoxSerialNumber   *string // Serial number of the Charge Box (max 25)
	FirmwareVersion         *string // Firmware version (max 50 chars)
	Iccid                   *string // ICCID of the modem's SIM card (max 20)
	Imsi                    *string // IMSI of the modem's SIM card (max 20)
	MeterType               *string // Type of the main meter (max 25 chars)
	MeterSerialNumber       *string // Serial number of the main meter (max 25)
}

// ReqMessage represents an OCPP 1.6 BootNotification.req message.
type ReqMessage struct {
	ChargePointVendor       types.CiString20Type
	ChargePointModel        types.CiString20Type
	ChargePointSerialNumber *types.CiString25Type
	ChargeBoxSerialNumber   *types.CiString25Type
	FirmwareVersion         *types.CiString50Type
	Iccid                   *types.CiString20Type
	Imsi                    *types.CiString20Type
	MeterType               *types.CiString25Type
	MeterSerialNumber       *types.CiString25Type
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	chargePointVendor       types.CiString20Type
	chargePointModel        types.CiString20Type
	chargePointSerialNumber types.CiString25Type
	chargeBoxSerialNumber   types.CiString25Type
	firmwareVersion         types.CiString50Type
	iccid                   types.CiString20Type
	imsi                    types.CiString20Type
	meterType               types.CiString25Type
	meterSerialNumber       types.CiString25Type
}

// Req creates a BootNotification.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - ChargePointVendor is empty or exceeds 20 characters
//   - ChargePointModel is empty or exceeds 20 characters
//   - Any optional field exceeds its maximum length or contains invalid chars
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if errs != nil {
		return ReqMessage{
			ChargePointVendor:       types.CiString20Type{},
			ChargePointModel:        types.CiString20Type{},
			ChargePointSerialNumber: nil,
			ChargeBoxSerialNumber:   nil,
			FirmwareVersion:         nil,
			Iccid:                   nil,
			Imsi:                    nil,
			MeterType:               nil,
			MeterSerialNumber:       nil,
		}, errors.Join(errs...)
	}

	return buildReqMessage(input, validated), nil
}

// validateReqInput validates all fields in ReqInput and returns validated
// values along with any errors.
func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	// Validate required fields
	validated.chargePointVendor, errs = validateChargePointVendor(
		input.ChargePointVendor,
		errs,
	)
	validated.chargePointModel, errs = validateChargePointModel(
		input.ChargePointModel,
		errs,
	)

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
	if input.ChargePointSerialNumber != nil {
		validated.chargePointSerialNumber, errs = validateCiString25(
			*input.ChargePointSerialNumber,
			"chargePointSerialNumber",
			errs,
		)
	}

	if input.ChargeBoxSerialNumber != nil {
		validated.chargeBoxSerialNumber, errs = validateCiString25(
			*input.ChargeBoxSerialNumber,
			"chargeBoxSerialNumber",
			errs,
		)
	}

	if input.FirmwareVersion != nil {
		validated.firmwareVersion, errs = validateFirmwareVersion(
			*input.FirmwareVersion,
			errs,
		)
	}

	if input.Iccid != nil {
		validated.iccid, errs = validateCiString20(
			*input.Iccid,
			"iccid",
			errs,
		)
	}

	if input.Imsi != nil {
		validated.imsi, errs = validateCiString20(
			*input.Imsi,
			"imsi",
			errs,
		)
	}

	if input.MeterType != nil {
		validated.meterType, errs = validateCiString25(
			*input.MeterType,
			"meterType",
			errs,
		)
	}

	if input.MeterSerialNumber != nil {
		validated.meterSerialNumber, errs = validateCiString25(
			*input.MeterSerialNumber,
			"meterSerialNumber",
			errs,
		)
	}

	return validated, errs
}

// validateChargePointVendor validates the chargePointVendor field.
func validateChargePointVendor(
	vendor string,
	errs []error,
) (types.CiString20Type, []error) {
	ciStr, err := types.NewCiString20Type(vendor)
	if err != nil {
		return types.CiString20Type{}, append(
			errs,
			fmt.Errorf("chargePointVendor: %w", err),
		)
	}

	return ciStr, errs
}

// validateChargePointModel validates the chargePointModel field.
func validateChargePointModel(
	model string,
	errs []error,
) (types.CiString20Type, []error) {
	ciStr, err := types.NewCiString20Type(model)
	if err != nil {
		return types.CiString20Type{}, append(
			errs,
			fmt.Errorf("chargePointModel: %w", err),
		)
	}

	return ciStr, errs
}

// validateCiString20 validates a CiString20 field with the given field name.
func validateCiString20(
	value string,
	fieldName string,
	errs []error,
) (types.CiString20Type, []error) {
	ciStr, err := types.NewCiString20Type(value)
	if err != nil {
		return types.CiString20Type{}, append(
			errs,
			fmt.Errorf("%s: %w", fieldName, err),
		)
	}

	return ciStr, errs
}

// validateCiString25 validates a CiString25 field with the given field name.
func validateCiString25(
	value string,
	fieldName string,
	errs []error,
) (types.CiString25Type, []error) {
	ciStr, err := types.NewCiString25Type(value)
	if err != nil {
		return types.CiString25Type{}, append(
			errs,
			fmt.Errorf("%s: %w", fieldName, err),
		)
	}

	return ciStr, errs
}

// validateFirmwareVersion validates the firmwareVersion field.
func validateFirmwareVersion(
	version string,
	errs []error,
) (types.CiString50Type, []error) {
	ciStr, err := types.NewCiString50Type(version)
	if err != nil {
		return types.CiString50Type{}, append(
			errs,
			fmt.Errorf("firmwareVersion: %w", err),
		)
	}

	return ciStr, errs
}

// buildReqMessage constructs the final ReqMessage with validated fields.
func buildReqMessage(input ReqInput, validated reqValidation) ReqMessage {
	msg := ReqMessage{
		ChargePointVendor:       validated.chargePointVendor,
		ChargePointModel:        validated.chargePointModel,
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	}

	if input.ChargePointSerialNumber != nil {
		msg.ChargePointSerialNumber = &validated.chargePointSerialNumber
	}

	if input.ChargeBoxSerialNumber != nil {
		msg.ChargeBoxSerialNumber = &validated.chargeBoxSerialNumber
	}

	if input.FirmwareVersion != nil {
		msg.FirmwareVersion = &validated.firmwareVersion
	}

	if input.Iccid != nil {
		msg.Iccid = &validated.iccid
	}

	if input.Imsi != nil {
		msg.Imsi = &validated.imsi
	}

	if input.MeterType != nil {
		msg.MeterType = &validated.meterType
	}

	if input.MeterSerialNumber != nil {
		msg.MeterSerialNumber = &validated.meterSerialNumber
	}

	return msg
}
