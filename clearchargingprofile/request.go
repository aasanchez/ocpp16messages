package clearchargingprofile

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a
// ClearChargingProfile.req message. All fields are optional per OCPP 1.6.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Optional: The ID of the charging profile to clear
	Id *int
	// Optional: Specifies the ID of the connector for which to clear profiles
	// (0 = all connectors)
	ConnectorId *int
	// Optional: Specifies to clear profiles with this purpose
	// ("ChargePointMaxProfile", "TxDefaultProfile", "TxProfile")
	ChargingProfilePurpose *string
	// Optional: Specifies the stack level for which to clear profiles
	StackLevel *int
}

// ReqMessage represents an OCPP 1.6 ClearChargingProfile.req message.
type ReqMessage struct {
	Id                     *types.Integer
	ConnectorId            *types.Integer
	ChargingProfilePurpose *types.ChargingProfilePurposeType
	StackLevel             *types.Integer
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	id                     *types.Integer
	connectorId            *types.Integer
	chargingProfilePurpose *types.ChargingProfilePurposeType
	stackLevel             *types.Integer
}

// Req creates a ClearChargingProfile.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - Id (if provided) is negative or exceeds uint16 max value (65535)
//   - ConnectorId (if provided) is negative or exceeds uint16 max (65535)
//   - ChargingProfilePurpose (if provided) is not a valid value
//   - StackLevel (if provided) is negative or exceeds uint16 max (65535)
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if errs != nil {
		return ReqMessage{
			Id:                     nil,
			ConnectorId:            nil,
			ChargingProfilePurpose: nil,
			StackLevel:             nil,
		}, errors.Join(errs...)
	}

	return ReqMessage{
		Id:                     validated.id,
		ConnectorId:            validated.connectorId,
		ChargingProfilePurpose: validated.chargingProfilePurpose,
		StackLevel:             validated.stackLevel,
	}, nil
}

// validateReqInput validates all fields in ReqInput and returns validated
// values along with any errors.
func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	if input.Id != nil {
		validated.id, errs = validateId(*input.Id, errs)
	}

	if input.ConnectorId != nil {
		validated.connectorId, errs = validateConnectorId(
			*input.ConnectorId,
			errs,
		)
	}

	if input.ChargingProfilePurpose != nil {
		validated.chargingProfilePurpose, errs = validatePurpose(
			*input.ChargingProfilePurpose,
			errs,
		)
	}

	if input.StackLevel != nil {
		validated.stackLevel, errs = validateStackLevel(*input.StackLevel, errs)
	}

	return validated, errs
}

// validateId validates the id field.
func validateId(id int, errs []error) (*types.Integer, []error) {
	val, err := types.NewInteger(id)
	if err != nil {
		return nil, append(errs, fmt.Errorf("id: %w", err))
	}

	return &val, errs
}

// validateConnectorId validates the connectorId field.
func validateConnectorId(
	connectorId int,
	errs []error,
) (*types.Integer, []error) {
	val, err := types.NewInteger(connectorId)
	if err != nil {
		return nil, append(errs, fmt.Errorf("connectorId: %w", err))
	}

	return &val, errs
}

// validatePurpose validates the chargingProfilePurpose field.
func validatePurpose(
	purpose string,
	errs []error,
) (*types.ChargingProfilePurposeType, []error) {
	purposeType := types.ChargingProfilePurposeType(purpose)
	if !purposeType.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf("chargingProfilePurpose: %w", types.ErrInvalidValue),
		)
	}

	return &purposeType, errs
}

// validateStackLevel validates the stackLevel field.
func validateStackLevel(
	stackLevel int,
	errs []error,
) (*types.Integer, []error) {
	val, err := types.NewInteger(stackLevel)
	if err != nil {
		return nil, append(errs, fmt.Errorf("stackLevel: %w", err))
	}

	return &val, errs
}
