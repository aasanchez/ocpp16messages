// Package ocpp16messages provides OCPP 1.6-compliant message types and
// validation for EV charging station management.
//
// This library implements strict type validation for OCPP 1.6 protocol
// messages, including CiString types (case-insensitive strings with
// length validation), DateTime types (RFC3339-compliant timestamps
// normalized to UTC), and Integer types (validated uint16 values).
//
// All types use the constructor pattern with validation, returning
// errors for invalid inputs rather than panicking. Types are designed
// to be thread-safe with immutable fields and value receivers.
package ocpp16messages
