package sharedtypes

const (
	ErrCiStringEmpty    = "%s can't be empty, %w"
	ErrorInvalid        = "error: invalid %s -> %w"
	ErrorMissing        = "error: missing %s"
	ErrorWrapper        = "%w: %q"
	ErrorExpectedError  = "Expected Error: %v"
	ErrorStringMismatch = "Expected: %s, got: %s"
	ErrorValueMismatch  = "expected Value() = %q, got %q"
)
