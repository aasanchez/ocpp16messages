package remotestoptransaction

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a
// RemoteStopTransaction.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: RemoteStopTransactionStatus value ("Accepted" or "Rejected")
	Status string
}

// ConfMessage represents an OCPP 1.6 RemoteStopTransaction.conf message.
type ConfMessage struct {
	Status types.RemoteStopTransactionStatus
}

// Conf creates a RemoteStopTransaction.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid RemoteStopTransactionStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.RemoteStopTransactionStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
