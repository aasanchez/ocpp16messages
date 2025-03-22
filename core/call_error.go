// Package core provides shared types and utilities used throughout the OCPP 1.6 message validation system.
package core

import (
	"encoding/json"
)

// CALLERRORMessageTypeId represents the message type ID for OCPP CALLERROR messages.
const CALLERRORMessageTypeId = 4

// Library-specific error code prefix to avoid confusion with official OCPP error codes.
const LibraryErrorPrefix = "Library."

// NewCallErrorMessage creates a properly formatted OCPP CALLERROR message as a JSON-encoded byte slice.
// It is primarily intended for use in proxies, test harnesses, or emulators to simulate CSMS behavior.
//
// Parameters:
// - uniqueID: The ID of the message that this error responds to.
// - errorCode: A short string identifying the error type (e.g., "ProtocolError").
// - errorDescription: A human-readable explanation of the error.
// - errorDetails: Optional structured data to provide more context (can be nil).
//
// Example usage:
//
//	errMsg, err := NewCallErrorMessage("123", "Library.ValidationError", "Missing required field", nil)
//
// Returns the JSON-encoded CALLERROR message or an error.
func NewCallErrorMessage(uniqueID string, errorCode string, errorDescription string, errorDetails interface{}) ([]byte, error) {
	callError := []interface{}{
		CALLERRORMessageTypeId,
		uniqueID,
		errorCode,
		errorDescription,
		errorDetails,
	}
	return json.Marshal(callError)
}

// NewLibraryValidationError creates a standardized validation error message with the "Library.ValidationError" code.
func NewLibraryValidationError(uniqueID, message string) ([]byte, error) {
	return NewCallErrorMessage(uniqueID, LibraryErrorPrefix+"ValidationError", message, nil)
}
