package firmwarestatusnotification

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a
// FirmwareStatusNotification.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	Status string // Required: FirmwareStatus value
}

// ReqMessage represents an OCPP 1.6 FirmwareStatusNotification.req message.
type ReqMessage struct {
	Status types.FirmwareStatus
}

// Req creates a FirmwareStatusNotification.req message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid FirmwareStatus value
func Req(input ReqInput) (ReqMessage, error) {
	status := types.FirmwareStatus(input.Status)

	if !status.IsValid() {
		return ReqMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ReqMessage{Status: status}, nil
}
