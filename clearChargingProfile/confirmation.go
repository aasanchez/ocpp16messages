package clearChargingProfile

import (
	"fmt"

	ct "github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a
// ClearChargingProfile.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: ClearChargingProfileStatus value ("Accepted" or "Unknown")
	Status string
}

// ConfMessage represents an OCPP 1.6 ClearChargingProfile.conf message.
type ConfMessage struct {
	Status ct.ClearChargingProfileStatus
}

// Conf creates a ClearChargingProfile.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ClearChargingProfileStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := ct.ClearChargingProfileStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
