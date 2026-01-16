package getLocalListVersion

import (
	"fmt"

	mt "github.com/aasanchez/ocpp16messages/getLocalListVersion/types"
)

// ConfInput represents the raw input data for creating a
// GetLocalListVersion.conf message. The constructor Conf validates all fields
// automatically.
type ConfInput struct {
	// Required: Version number of the local authorization list.
	// Special values:
	//   -1: Charge Point does not support Local Authorization Lists
	//    0: Local authorization list is empty
	ListVersion int
}

// ConfMessage represents an OCPP 1.6 GetLocalListVersion.conf message.
type ConfMessage struct {
	ListVersion mt.ListVersionNumber
}

// Conf creates a GetLocalListVersion.conf message from the given input.
// It validates all fields and returns an error if:
//   - ListVersion cannot be converted to a valid int32
func Conf(input ConfInput) (ConfMessage, error) {
	listVersion, err := mt.NewListVersionNumber(fmt.Sprintf("%d", input.ListVersion))
	if err != nil {
		return ConfMessage{}, fmt.Errorf("listVersion: %w", err)
	}

	return ConfMessage{ListVersion: listVersion}, nil
}
