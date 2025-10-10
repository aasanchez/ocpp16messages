package sharedtypes

import (
	"fmt"
	"strconv"
)

// Integer is a compact unsigned 16-bit number
// tailored for OCPP 1.6 payloads.
//
// OCPP 1.6
//   - Use for fields that naturally fit 0..65535:
//     connector indices, counters, compact sizes,
//     small identifiers exchanged by CP and CSMS.
//   - Keeps messages small and semantics clear.
//     For wider ranges, define a new type.
//
// Behavior
// - Immutable after construction
// - Zero value is valid (0)
// - Concurrency-safe (read-only)
// - No panics; errors are returned
//
// Interop
//   - Parses base-10 strings only (e.g. "73")
//   - Rejects negatives, hex, floats, overflow
//   - Encode with Value(); validate inputs with
//     SetInteger when reading from the wire
//
// Errors
//   - Wraps strconv.ParseUint so callers can use
//     errors.Is/As to inspect the cause
//
// Stability
// - Minimal API, stable behavior across minors
// - Simple shape aids agents and tooling
//
// See
// - ExampleSetInteger, ExampleSetInteger_invalid
// - shared/types/errors.go for error text shapes
type Integer struct {
	value uint16
}

// SetInteger parses a base-10 string into Integer.
//
// Validation
// - Accepts 0..65535 (uint16)
// - Errors on empty, non-decimal, negative, overflow
// - Error wraps parse failure for context
//
// Usage
//
//	v, err := SetInteger("73")
//	if err != nil { /* handle */ }
//	n := v.Value() // 73
func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid Integer: %w", err)
	}

	return Integer{value: uint16(parsedValue)}, nil
}

// Value returns the stored number as uint16.
//
// Guarantees
// - O(1), no allocation
// - Never fails
func (integer Integer) Value() uint16 {
	return integer.value
}
