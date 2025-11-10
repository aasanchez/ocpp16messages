package sharedtypes

const (
	ErrCiStringEmpty     = "%s can't be empty, %w"
	ErrorInvalid         = "error: invalid %s -> %w"
	MissingValue         = "error: Missing %s in %q"
	ErrorExpectedError   = "expected Error: %v"
	ErrorUnexpectedError = "unexpected Error: %v"
	ErrorStringMismatch  = "expected String: %s, got: %s"
	ErrorValueMismatch   = "expected Value = %q, got %q"
)
