package reset

import (
	"fmt"

	rt "github.com/aasanchez/ocpp16messages/reset/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating a Reset.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The type of reset to perform ("Hard" or "Soft")
	Type string
}

// ReqMessage represents an OCPP 1.6 Reset.req message.
type ReqMessage struct {
	Type rt.ResetType
}

// Req creates a Reset.req message from the given input.
// It validates all fields and returns an error if:
//   - Type is not a valid ResetType value
func Req(input ReqInput) (ReqMessage, error) {
	resetType := rt.ResetType(input.Type)

	if !resetType.IsValid() {
		return ReqMessage{}, fmt.Errorf("type: %w", st.ErrInvalidValue)
	}

	return ReqMessage{Type: resetType}, nil
}
