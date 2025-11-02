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
	ErrorInvalid         = "error: invalid %s -> %w"
	MissingValue         = "error: Missing %s in %q"
	ErrorExpectedError   = "expected Error: %v"
	ErrorUnexpectedError = "unexpected Error: %v"
	ErrorStringMismatch  = "expected String: %s, got: %s"
	ErrorValueMismatch   = "expected Value = %q, got %q"
)
