package remotestarttransaction

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a
// RemoteStartTransaction.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: RemoteStartTransactionStatus value ("Accepted" or "Rejected")
	Status string
}

// ConfMessage represents an OCPP 1.6 RemoteStartTransaction.conf message.
type ConfMessage struct {
	Status types.RemoteStartTransactionStatus
}

// Conf creates a RemoteStartTransaction.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid RemoteStartTransactionStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.RemoteStartTransactionStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
