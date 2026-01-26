// Package types provides Open Charge Point Protocol (OCPP) 1.6
// core data types with strict validation for EV charging.
// validation.
//
// This package implements validated types for the OCPP 1.6 protocol including:
//   - CiString types: case-insensitive strings with maximum length constraints
//   - DateTime: RFC3339-compliant timestamps that must already be UTC
//   - Integer: validated uint16 values
//
// All types enforce validation at construction time and are designed for
// thread-safe concurrent use with immutable fields and value receivers.
package types
