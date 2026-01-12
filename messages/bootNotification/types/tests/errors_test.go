package types_test

// Test error message format strings for public API tests.
const (
	// ErrorMethodMismatch is used in tests to format method return value
	// comparison failures. The first parameter is the actual value, and
	// the second parameter is the expected value.
	ErrorMethodMismatch = "RegistrationStatus.String() = %v, want %v"
)
