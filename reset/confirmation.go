package reset

import (
	"fmt"

	rt "github.com/aasanchez/ocpp16messages/reset/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a Reset.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: ResetStatus value ("Accepted" or "Rejected")
	Status string
}

// ConfMessage represents an OCPP 1.6 Reset.conf message.
type ConfMessage struct {
	Status rt.ResetStatus
}

// Conf creates a Reset.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ResetStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := rt.ResetStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
