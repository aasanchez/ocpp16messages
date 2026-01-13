package changeAvailability

import (
	"fmt"

	mcat "github.com/aasanchez/ocpp16messages/messages/changeAvailability/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
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
	Status mcat.AvailabilityStatus
}

// Conf creates a ChangeAvailability.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid AvailabilityStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := mcat.AvailabilityStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
