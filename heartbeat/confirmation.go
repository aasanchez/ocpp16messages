package heartbeat

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// ConfInput represents the raw input data for creating a Heartbeat.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	CurrentTime string // Required: RFC3339 formatted current time
}

// ConfMessage represents an OCPP 1.6 Heartbeat.conf message.
type ConfMessage struct {
	CurrentTime st.DateTime
}

// Conf creates a Heartbeat.conf message from the given input.
// It validates all fields automatically and returns an error if:
//   - CurrentTime is empty
//   - CurrentTime is not a valid RFC3339 date
func Conf(input ConfInput) (ConfMessage, error) {
	currentTime, err := st.NewDateTime(input.CurrentTime)
	if err != nil {
		return ConfMessage{}, fmt.Errorf("currentTime: %w", err)
	}

	return ConfMessage{CurrentTime: currentTime}, nil
}
