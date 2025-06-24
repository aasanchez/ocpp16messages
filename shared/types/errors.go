package types

import (
	"errors"
)

var (
	ErrEmptyValueNotAllowed = errors.New("value must not be empty")
	ErrExceedsMaxLength     = errors.New("value exceeds maximum allowed length")
	ErrNonPrintableASCII    = errors.New("value contains non-printable ASCII characters")
	ErrInvalidIdTag         = errors.New("invalid idTag")
	ErrCiStringNotSet       = errors.New("optional CiString value not set")
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
