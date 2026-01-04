package authorize

// Test error message format strings.
const (
	// ErrorResponseStatusMismatch is used in tests to format
	// Response.IdTagInfo.Status comparison failures. The first parameter
	// is the actual value, and the second parameter is the expected value.
	ErrorResponseStatusMismatch = "Response.IdTagInfo.Status = %v, want %v"
)
