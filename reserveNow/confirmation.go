package reserveNow

import (
	"fmt"

	rt "github.com/aasanchez/ocpp16messages/reserveNow/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a ReserveNow.conf
// message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: ReservationStatus value (Accepted, Faulted, Occupied,
	// Rejected, Unavailable)
	Status string
}

// ConfMessage represents an OCPP 1.6 ReserveNow.conf message.
type ConfMessage struct {
	Status rt.ReservationStatus
}

// Conf creates a ReserveNow.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ReservationStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := rt.ReservationStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
