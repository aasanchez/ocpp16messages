package unlockconnector

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating an UnlockConnector.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: UnlockStatus value ("Unlocked", "UnlockFailed", "NotSupported")
	Status string
}

// ConfMessage represents an OCPP 1.6 UnlockConnector.conf message.
type ConfMessage struct {
	Status types.UnlockStatus
}

// Conf creates an UnlockConnector.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid UnlockStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.UnlockStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
