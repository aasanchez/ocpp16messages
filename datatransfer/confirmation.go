package datatransfer

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ConfInput represents the raw input data for creating a DataTransfer.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	Status string  // Required: DataTransferStatus value
	Data   *string // Optional: Data payload (unbounded per OCPP spec)
}

// ConfMessage represents an OCPP 1.6 DataTransfer.conf message.
type ConfMessage struct {
	Status types.DataTransferStatus
	Data   *string
}

// Conf creates a DataTransfer.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid DataTransferStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := types.DataTransferStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	msg := ConfMessage{
		Status: status,
		Data:   nil,
	}

	if input.Data != nil {
		copiedData := *input.Data
		msg.Data = &copiedData
	}

	return msg, nil
}
