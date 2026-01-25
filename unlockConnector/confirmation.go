package unlockConnector

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
	ut "github.com/aasanchez/ocpp16messages/unlockConnector/types"
)

// ConfInput represents the raw input data for creating an UnlockConnector.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: UnlockStatus value ("Unlocked", "UnlockFailed", "NotSupported")
	Status string
}

// ConfMessage represents an OCPP 1.6 UnlockConnector.conf message.
type ConfMessage struct {
	Status ut.UnlockStatus
}

// Conf creates an UnlockConnector.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid UnlockStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := ut.UnlockStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
