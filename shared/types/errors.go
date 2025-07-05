package types

import (
	"errors"
)

var (
	ErrCiStringNotSet                   = errors.New("optional CiString value was accessed but not set; ensure it has been initialized")
	ErrEmptyValueNotAllowed             = errors.New("validation error: a required value was empty; please provide a non-empty input")
	ErrExceedsMaxLength                 = errors.New("validation error: the value exceeds the maximum allowed character length for this field")
	ErrInvalidAuthorizationStatus       = errors.New("invalid value provided for authorization status; expected one of the allowed enum values")
	ErrInvalidCancelReservationStatus   = errors.New("invalid status for CancelReservation; expected a valid reservation cancellation status")
	ErrInvalidChangeAvailabilityStatus  = errors.New("invalid status for ChangeAvailability; please provide a valid availability status value")
	ErrInvalidChangeAvailabilityType    = errors.New("invalid type for ChangeAvailability; only valid availability types are accepted")
	ErrInvalidChangeConfigurationStatus = errors.New("invalid response for ChangeConfiguration; the status is not recognized")
	ErrInvalidExpiryDate                = errors.New("invalid format or value for expiryDate; expected a valid ISO 8601 datetime string")
	ErrInvalidIdTag                     = errors.New("invalid idTag; must meet the expected format and constraints (e.g., length and characters)")
	ErrInvalidIdTagInfo                 = errors.New("invalid idTagInfo structure; ensure all required fields are valid and correctly set")
	ErrInvalidIdToken                   = errors.New("invalid IdToken value; expected a valid identifier for the token")
	ErrInvalidParentIdTag               = errors.New("invalid parentIdTag; must match expected token format and constraints")
	ErrInvalidRegistrationStatus        = errors.New("invalid registration status; expected a recognized status value from the specification")
	ErrMissingStatus                    = errors.New("missing required field: idTagInfo.Status; this field is mandatory and must be set")
	ErrNonPrintableASCII                = errors.New("value contains characters outside the printable ASCII range (0x20 to 0x7E); only standard text is allowed")
	ErrNotFound                         = errors.New("requested item was not found")
)

// Internal error format templates used for test assertions or internal error wrapping
const (
	ErrContainsFmt            = "expected error to contain: %q, but got: %q"
	ErrExpectedError          = "expected an error, but got: %v"
	ErrExpectedNil            = "expected nil result for nil input, but got: %v"
	ErrExpectedNoError        = "unexpected error encountered: %v"
	ErrExpectedNoValidation   = "expected no validation error, but received: %v"
	ErrExpectedNonNilError    = "expected an error for %s, but none occurred"
	ErrExpectedValidationFail = "expected validation to fail, but it passed"
	ErrExpectedValueMismatch  = "expected value = %q, but got %q"
	ErrFmtFieldWrapped        = "%s: %w"
)
