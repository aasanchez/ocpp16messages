package types

import (
	"errors"
)

var (
	ErrCiStringNotSet                   = errors.New("optional CiString value not set")
	ErrEmptyValueNotAllowed             = errors.New("value must not be empty")
	ErrExceedsMaxLength                 = errors.New("value exceeds maximum allowed length")
	ErrInvalidAuthorizationStatus       = errors.New("invalid authorization status")
	ErrInvalidCancelReservationStatus   = errors.New("invalid clear Cancel Reservation")
	ErrInvalidChangeAvailabilityStatus  = errors.New("invalid change availability status")
	ErrInvalidChangeAvailabilityType    = errors.New("invalid change availability type")
	ErrInvalidChangeConfigurationStatus = errors.New("invalid change configuration status")
	ErrInvalidExpiryDate                = errors.New("invalid expiryDate")
	ErrInvalidIdTag                     = errors.New("invalid idTag")
	ErrInvalidIdTagInfo                 = errors.New("invalid idTagInfo")
	ErrInvalidIdToken                   = errors.New("invalid IdToken")
	ErrInvalidParentIdTag               = errors.New("invalid parentIdTag")
	ErrInvalidRegistrationStatus        = errors.New("invalid registration status")
	ErrMissingStatus                    = errors.New("missing required field: idTagInfo.Status")
	ErrNonPrintableASCII                = errors.New("value contains non-printable ASCII characters")
)

const (
	ErrContainsFmt            = "expected error to contain:%q got: %q"
	ErrExpectedError          = "expected error: %v"
	ErrExpectedNil            = "expected nil result for nil input, got: %v"
	ErrExpectedNoError        = "unexpected error: %v"
	ErrExpectedNoValidation   = "expected no validation error, got: %v"
	ErrExpectedNonNilError    = "expected error for %s"
	ErrExpectedValidationFail = "expected validation error, got nil"
	ErrExpectedValueMismatch  = "expected Value() = %q, got %q"
	ErrFmtFieldWrapped        = "%s: %w"
)
