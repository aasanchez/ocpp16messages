package changeavailability

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a
// ChangeAvailability.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: AvailabilityStatus value
	// ("Accepted", "Rejected", or "Scheduled")
	Status string
}

// ConfMessage represents an OCPP 1.6 ChangeAvailability.conf message.
type ConfMessage struct {
	Status types.AvailabilityStatus
}

// Conf creates a ChangeAvailability.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid AvailabilityStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.AvailabilityStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
