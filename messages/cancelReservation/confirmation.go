package cancelReservation

import (
	"fmt"

	mct "github.com/aasanchez/ocpp16messages/messages/cancelReservation/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfInput represents the raw input data for creating a CancelReservation.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: CancelReservationStatus value ("Accepted" or "Rejected")
	Status string
}

// ConfMessage represents an OCPP 1.6 CancelReservation.conf message.
type ConfMessage struct {
	Status mct.CancelReservationStatus
}

// Conf creates a CancelReservation.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid CancelReservationStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := mct.CancelReservationStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
