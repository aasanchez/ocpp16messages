package firmwareStatusNotification

import (
	"fmt"

	fn "github.com/aasanchez/ocpp16messages/firmwareStatusNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating a
// FirmwareStatusNotification.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	Status string // Required: FirmwareStatus value
}

// ReqMessage represents an OCPP 1.6 FirmwareStatusNotification.req message.
type ReqMessage struct {
	Status fn.FirmwareStatus
}

// Req creates a FirmwareStatusNotification.req message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid FirmwareStatus value
func Req(input ReqInput) (ReqMessage, error) {
	status := fn.FirmwareStatus(input.Status)

	if !status.IsValid() {
		return ReqMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ReqMessage{Status: status}, nil
}
