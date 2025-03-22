// Package core defines shared types and utilities for OCPP 1.6J messages.
//
// This file provides the CallError structure used for MessageTypeId 4 responses.
package core

// CallError represents the structure of an OCPP CALLERROR message (MessageTypeId = 4).
// This message is sent from a server to a client when an error occurs while processing
// a CALL (MessageTypeId = 2).
type CallError struct {
	MessageTypeId int                    `json:"-"`                      // Always 4 for CALLERROR
	UniqueID      string                 `json:"uniqueId"`               // ID of the failed CALL message
	ErrorCode     string                 `json:"errorCode"`              // Error code as defined by OCPP
	ErrorDesc     string                 `json:"errorDescription"`       // Optional human-readable description
	ErrorDetails  map[string]interface{} `json:"errorDetails,omitempty"` // Optional additional error context
}

// NewCallError creates a new CallError message.
//
// Parameters:
//   - uniqueID: the ID of the original CALL message
//   - code: error code (e.g. "ProtocolError")
//   - description: optional error description
//   - details: optional map of extended error data
func NewCallError(uniqueID, code, description string, details map[string]interface{}) CallError {
	return CallError{
		MessageTypeId: 4,
		UniqueID:      uniqueID,
		ErrorCode:     code,
		ErrorDesc:     description,
		ErrorDetails:  details,
	}
}
