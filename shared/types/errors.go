package types

import (
	"errors"
)

var (
	ErrInvalidIdTag                    = errors.New("invalid idTag")
	ErrInvalidIdToken                  = errors.New("invalid IdToken")
	ErrInvalidIdTagInfo                = errors.New("invalid idTagInfo")
	ErrInvalidAuthorizationStatus      = errors.New("invalid authorization status")
	ErrInvalidCancelReservationStatus  = errors.New("invalid clear Cancel Reservation")
	ErrInvalidRegistrationStatus       = errors.New("invalid registration status")
	ErrInvalidChangeAvailabilityStatus = errors.New("invalid change availability status")
	ErrInvalidChangeAvailabilityType   = errors.New("invalid change availability type")
	// ErrInvalidChangeConfigurationStatus indicates an unrecognized ChangeConfiguration status value.
	ErrInvalidChangeConfigurationStatus = errors.New("invalid change configuration status")
	ErrMissingStatus                    = errors.New("missing required field: idTagInfo.Status")
	ErrInvalidExpiryDate                = errors.New("invalid expiryDate")
	ErrInvalidParentIdTag               = errors.New("invalid parentIdTag")
	ErrEmptyValueNotAllowed             = errors.New("value must not be empty")
	ErrExceedsMaxLength                 = errors.New("value exceeds maximum allowed length")
	ErrNonPrintableASCII                = errors.New("value contains non-printable ASCII characters")
	ErrCiStringNotSet                   = errors.New("optional CiString value not set")
)

const (
	ErrFmtFieldWrapped        = "%s: %w"
	ErrContainsFmt            = "expected error to contain:%q got: %q"
	ErrExpectedNoError        = "unexpected error: %v"
	ErrExpectedError          = "expected error: %v"
	ErrExpectedNil            = "expected nil result for nil input, got: %v"
	ErrExpectedNonNilError    = "expected error for %s"
	ErrExpectedValidationFail = "expected validation error, got nil"
	ErrExpectedNoValidation   = "expected no validation error, got: %v"
	ErrExpectedValueMismatch  = "expected Value() = %q, got %q"
)
