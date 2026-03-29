package sendlocallist

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a SendLocalList.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: UpdateStatus value indicating the result of the update.
	Status string
}

// ConfMessage represents an OCPP 1.6 SendLocalList.conf message.
type ConfMessage struct {
	Status types.UpdateStatus
}

// Conf creates a SendLocalList.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid UpdateStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.UpdateStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
