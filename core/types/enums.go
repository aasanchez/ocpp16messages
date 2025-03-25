package types

// MessageType represents the three standard OCPP 1.6 message types.
//
// These correspond to the messageTypeId used in the OCPP JSON protocol:
//
//	2 - CALL (Request)
//	3 - CALLRESULT (Response)
//	4 - CALLERROR (Error response)
type MessageType int

const (
	// CALL is a request sent from client to server.
	CALL MessageType = 2

	// CALLRESULT is a response sent from server to client.
	CALLRESULT MessageType = 3

	// CALLERROR is an error response sent when a request cannot be processed.
	CALLERROR MessageType = 4
)
