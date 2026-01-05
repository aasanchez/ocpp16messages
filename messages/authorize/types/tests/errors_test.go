package types_test

// Test error message format strings for public API tests.
const (
	// ErrorExpectedError is used in tests to format error assertion
	// messages.
	ErrorExpectedError = "Expected Error: %v"
	// ErrorMismatch is used in tests to format expected vs actual value
	// comparison failures.
	ErrorMismatch = "expected %q, got %q"
	// ErrorMethodMismatch is used in tests to format method return value
	// comparison failures. The first parameter is the actual value, and
	// the second parameter is the expected value.
	ErrorMethodMismatch = "AuthorizationStatus.String() = %v, want %v"
	// ErrorStatusMismatch is used in tests to format IdTagInfo.Status
	// comparison failures. The first parameter is the actual value, and
	// the second parameter is the expected value.
	ErrorStatusMismatch = "IdTagInfo.Status = %v, want %v"
)
