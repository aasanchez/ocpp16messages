package types

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
	// stackLevelMin is the minimum valid stack level.
	stackLevelMin = 0
)

// ChargingProfileInput represents the raw input data for creating a
// ChargingProfile. The constructor NewChargingProfile validates all fields
// automatically.
type ChargingProfileInput struct {
	// Required: Unique identifier of this profile
	ChargingProfileId int
	// Optional: Only valid if ChargingProfilePurpose is TxProfile
	TransactionId *int
	// Required: Value determining level in hierarchy stack of profiles
	StackLevel int
	// Required: Defines the purpose of the schedule
	ChargingProfilePurpose string
	// Required: Indicates the kind of schedule
	ChargingProfileKind string
	// Optional: Only for Recurring profiles
	RecurrencyKind *string
	// Optional: Point in time at which the profile starts to be valid
	// (RFC3339 format)
	ValidFrom *string
	// Optional: Point in time at which the profile stops to be valid
	// (RFC3339 format)
	ValidTo *string
	// Required: Contains limits for the available power or current over time
	ChargingSchedule st.ChargingScheduleInput
}

// ChargingProfile represents a charging profile as defined in OCPP 1.6.
type ChargingProfile struct {
	chargingProfileId      st.Integer
	transactionId          *st.Integer
	stackLevel             st.Integer
	chargingProfilePurpose st.ChargingProfilePurposeType
	chargingProfileKind    ChargingProfileKindType
	recurrencyKind         *RecurrencyKindType
	validFrom              *st.DateTime
	validTo                *st.DateTime
	chargingSchedule       st.ChargingSchedule
}

// NewChargingProfile creates a new ChargingProfile from input.
// Returns an error if:
//   - ChargingProfileId is negative or exceeds uint16 max value (65535)
//   - TransactionId (if provided) is negative or exceeds uint16 max (65535)
//   - StackLevel is negative or exceeds uint16 max value (65535)
//   - ChargingProfilePurpose is not a valid value
//   - ChargingProfileKind is not a valid value
//   - RecurrencyKind (if provided) is not a valid value
//   - ValidFrom (if provided) is not a valid RFC3339 timestamp
//   - ValidTo (if provided) is not a valid RFC3339 timestamp
//   - ChargingSchedule is invalid
//
//nolint:cyclop,funlen,revive // validation requires checking many fields
func NewChargingProfile(input ChargingProfileInput) (ChargingProfile, error) {
	var errs []error

	chargingProfileId, err := validateChargingProfileId(input.ChargingProfileId)
	if err != nil {
		errs = append(errs, err)
	}

	transactionId, err := validateTransactionId(input.TransactionId)
	if err != nil {
		errs = append(errs, err)
	}

	stackLevel, err := validateStackLevel(input.StackLevel)
	if err != nil {
		errs = append(errs, err)
	}

	purpose, err := validateChargingProfilePurpose(input.ChargingProfilePurpose)
	if err != nil {
		errs = append(errs, err)
	}

	kind, err := validateChargingProfileKind(input.ChargingProfileKind)
	if err != nil {
		errs = append(errs, err)
	}

	recurrencyKind, err := validateRecurrencyKind(input.RecurrencyKind)
	if err != nil {
		errs = append(errs, err)
	}

	validFrom, err := validateValidFrom(input.ValidFrom)
	if err != nil {
		errs = append(errs, err)
	}

	validTo, err := validateValidTo(input.ValidTo)
	if err != nil {
		errs = append(errs, err)
	}

	chargingSchedule, err := st.NewChargingSchedule(input.ChargingSchedule)
	if err != nil {
		errs = append(errs, fmt.Errorf("chargingSchedule: %w", err))
	}

	if len(errs) > errCountZero {
		return ChargingProfile{}, errors.Join(errs...)
	}

	return ChargingProfile{
		chargingProfileId:      chargingProfileId,
		transactionId:          transactionId,
		stackLevel:             stackLevel,
		chargingProfilePurpose: purpose,
		chargingProfileKind:    kind,
		recurrencyKind:         recurrencyKind,
		validFrom:              validFrom,
		validTo:                validTo,
		chargingSchedule:       chargingSchedule,
	}, nil
}

// validateChargingProfileId validates the charging profile ID.
func validateChargingProfileId(id int) (st.Integer, error) {
	profileId, err := st.NewInteger(id)
	if err != nil {
		return st.Integer{}, fmt.Errorf("chargingProfileId: %w", err)
	}

	return profileId, nil
}

// validateTransactionId validates the optional transaction ID.
func validateTransactionId(id *int) (*st.Integer, error) {
	if id == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	txId, err := st.NewInteger(*id)
	if err != nil {
		return nil, fmt.Errorf("transactionId: %w", err)
	}

	return &txId, nil
}

// validateStackLevel validates the stack level.
func validateStackLevel(level int) (st.Integer, error) {
	if level < stackLevelMin {
		return st.Integer{}, fmt.Errorf("stackLevel: %w", st.ErrInvalidValue)
	}

	stackLevel, err := st.NewInteger(level)
	if err != nil {
		return st.Integer{}, fmt.Errorf("stackLevel: %w", err)
	}

	return stackLevel, nil
}

// validateChargingProfilePurpose validates the charging profile purpose.
func validateChargingProfilePurpose(
	purpose string,
) (st.ChargingProfilePurposeType, error) {
	purposeType := st.ChargingProfilePurposeType(purpose)
	if !purposeType.IsValid() {
		return "", fmt.Errorf("chargingProfilePurpose: %w", st.ErrInvalidValue)
	}

	return purposeType, nil
}

// validateChargingProfileKind validates the charging profile kind.
func validateChargingProfileKind(kind string) (ChargingProfileKindType, error) {
	kindType := ChargingProfileKindType(kind)
	if !kindType.IsValid() {
		return "", fmt.Errorf("chargingProfileKind: %w", st.ErrInvalidValue)
	}

	return kindType, nil
}

// validateRecurrencyKind validates the optional recurrency kind.
func validateRecurrencyKind(kind *string) (*RecurrencyKindType, error) {
	if kind == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	recurrencyKind := RecurrencyKindType(*kind)
	if !recurrencyKind.IsValid() {
		return nil, fmt.Errorf("recurrencyKind: %w", st.ErrInvalidValue)
	}

	return &recurrencyKind, nil
}

// validateValidFrom validates the optional valid from timestamp.
func validateValidFrom(validFrom *string) (*st.DateTime, error) {
	if validFrom == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	vf, err := st.NewDateTime(*validFrom)
	if err != nil {
		return nil, fmt.Errorf("validFrom: %w", err)
	}

	return &vf, nil
}

// validateValidTo validates the optional valid to timestamp.
func validateValidTo(validTo *string) (*st.DateTime, error) {
	if validTo == nil {
		return nil, nil //nolint:nilnil // nil is valid for optional field
	}

	vt, err := st.NewDateTime(*validTo)
	if err != nil {
		return nil, fmt.Errorf("validTo: %w", err)
	}

	return &vt, nil
}

// ChargingProfileId returns the unique identifier of this profile.
func (c ChargingProfile) ChargingProfileId() st.Integer {
	return c.chargingProfileId
}

// TransactionId returns the transaction ID, or nil if not specified.
func (c ChargingProfile) TransactionId() *st.Integer {
	if c.transactionId == nil {
		return nil
	}

	copiedTransactionId := *c.transactionId

	return &copiedTransactionId
}

// StackLevel returns the stack level of this profile.
func (c ChargingProfile) StackLevel() st.Integer {
	return c.stackLevel
}

// ChargingProfilePurpose returns the purpose of this profile.
//
//nolint:revive // line-length-limit: method signature exceeds limit
func (c ChargingProfile) ChargingProfilePurpose() st.ChargingProfilePurposeType {
	return c.chargingProfilePurpose
}

// ChargingProfileKind returns the kind of this profile.
func (c ChargingProfile) ChargingProfileKind() ChargingProfileKindType {
	return c.chargingProfileKind
}

// RecurrencyKind returns the recurrency kind, or nil if not specified.
func (c ChargingProfile) RecurrencyKind() *RecurrencyKindType {
	if c.recurrencyKind == nil {
		return nil
	}

	copiedRecurrencyKind := *c.recurrencyKind

	return &copiedRecurrencyKind
}

// ValidFrom returns the valid from timestamp, or nil if not specified.
func (c ChargingProfile) ValidFrom() *st.DateTime {
	if c.validFrom == nil {
		return nil
	}

	copiedValidFrom := *c.validFrom

	return &copiedValidFrom
}

// ValidTo returns the valid to timestamp, or nil if not specified.
func (c ChargingProfile) ValidTo() *st.DateTime {
	if c.validTo == nil {
		return nil
	}

	copiedValidTo := *c.validTo

	return &copiedValidTo
}

// ChargingSchedule returns the charging schedule for this profile.
func (c ChargingProfile) ChargingSchedule() st.ChargingSchedule {
	return c.chargingSchedule
}
