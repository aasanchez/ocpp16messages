package sharedtypes

const (
	ErrCiStringEmpty     = "%s can't be empty, %w"
	ErrorInvalid         = "Error: invalid %s -> %w"
	ErrorMissing         = "Error: missing %s"
	ErrorWrapper         = "%w: %q"
	ErrorExpectedError   = "Expected Error: %v"
	ErrorUnexpectedError = "Unexpected Error: %v"
	ErrorStringMismatch  = "Expected String: %s, got: %s"
	ErrorValueMismatch   = "Expected Value = %q, got %q"
)
