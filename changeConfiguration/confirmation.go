package changeConfiguration

import (
	"fmt"

	cct "github.com/aasanchez/ocpp16messages/changeConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
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
	Status cct.ConfigurationStatus
}

// Conf creates a ChangeConfiguration.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ConfigurationStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := cct.ConfigurationStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
