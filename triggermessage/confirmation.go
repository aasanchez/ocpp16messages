package triggermessage

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a TriggerMessage.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: TriggerMessageStatus value (Accepted, Rejected, NotImplemented)
	Status string
}

// ConfMessage represents an OCPP 1.6 TriggerMessage.conf message.
type ConfMessage struct {
	Status types.TriggerMessageStatus
}

// Conf creates a TriggerMessage.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid TriggerMessageStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.TriggerMessageStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
