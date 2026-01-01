package types

// Error message format strings used throughout the shared types package.
const (
	ErrCiStringEmpty         = "%s can't be empty, %w"
	ErrorInvalid             = "error: invalid %s -> %w"
	ErrorMissing             = "error: missing %s"
	ErrorWrapper             = "%w: %q"
	ErrorExpectedError       = "Expected Error: %v"
	ErrorExpectedEquals      = "Expected Equals: expected: %s, got: %s"
	ErrorMismatch            = "Expected: %s, got: %s"
	ErrExpectedValueMismatch = "expected Value() = %q, got %q"
)
