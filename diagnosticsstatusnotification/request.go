package diagnosticsstatusnotification

import (
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a
// DiagnosticsStatusNotification.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	Status string // Required: DiagnosticsStatus value
}

// ReqMessage represents an OCPP 1.6 DiagnosticsStatusNotification.req message.
type ReqMessage struct {
	Status types.DiagnosticsStatus
}

// Req creates a DiagnosticsStatusNotification.req message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid DiagnosticsStatus value
func Req(input ReqInput) (ReqMessage, error) {
	status := types.DiagnosticsStatus(input.Status)

	if !status.IsValid() {
		return ReqMessage{}, fmt.Errorf("status: %w", types.ErrInvalidValue)
	}

	return ReqMessage{Status: status}, nil
}
