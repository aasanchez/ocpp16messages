package changeconfiguration

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a
// ChangeConfiguration.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: ConfigurationStatus value
	// ("Accepted", "Rejected", "RebootRequired", or "NotSupported")
	Status string
}

// ConfMessage represents an OCPP 1.6 ChangeConfiguration.conf message.
type ConfMessage struct {
	Status types.ConfigurationStatus
}

// Conf creates a ChangeConfiguration.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ConfigurationStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.ConfigurationStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
