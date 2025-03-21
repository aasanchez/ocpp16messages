package test

// ptr returns a pointer to the given value.
// Useful for creating optional fields in test structs.
func ptr[T any](v T) *T {
	return &v
}
