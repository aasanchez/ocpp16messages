package setchargingprofile

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a
// SetChargingProfile.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: ChargingProfileStatus value indicating whether the Charge
	// Point has been able to process the message successfully.
	Status string
}

// ConfMessage represents an OCPP 1.6 SetChargingProfile.conf message.
type ConfMessage struct {
	Status types.ChargingProfileStatus
}

// Conf creates a SetChargingProfile.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ChargingProfileStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.ChargingProfileStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
