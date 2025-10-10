package sharedtypes

// Error formats for OCPP 1.6 types and helpers.
//
// Goal
// - Provide consistent, short messages across CP/CSMS code.
// - Make wrapping predictable so callers can inspect causes.
// - Keep outputs readable in logs, UIs, and tests.
//
// Usage patterns
//   - Wrap cause with context:
//     fmt.Errorf(ErrorWrapper, err, input)
//   - Validation:
//     fmt.Errorf(ErrorInvalid, "Integer", err)
//     fmt.Errorf(ErrorMissing, "IdTag")
//   - Tests:
//     t.Errorf(ErrorStringMismatch, want, got)
//     t.Errorf(ErrorValueMismatch, want, got)
//
// Placeholders
// - %s: short label (type/field)
// - %q: quoted value for readability
// - %v: printed value or error
// - %w: wrap an error (fmt.Errorf)
//
// Guidance
// - Keep messages concise; avoid leaking secrets.
// - Use short, protocol-aligned labels (e.g. "IdTag").
// - Prefer ErrorWrapper when adding error context.
// - For user-facing strings, map these to localized text.
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
