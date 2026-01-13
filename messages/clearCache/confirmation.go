package clearCache

import (
	"fmt"

	mct "github.com/aasanchez/ocpp16messages/messages/clearCache/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfInput represents the raw input data for creating a ClearCache.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: ClearCacheStatus value ("Accepted" or "Rejected")
	Status string
}

// ConfMessage represents an OCPP 1.6 ClearCache.conf message.
type ConfMessage struct {
	Status mct.ClearCacheStatus
}

// Conf creates a ClearCache.conf message from the given input.
// It validates all fields and returns an error if:
//   - Status is not a valid ClearCacheStatus value
func Conf(input ConfInput) (ConfMessage, error) {
	status := mct.ClearCacheStatus(input.Status)

	if !status.IsValid() {
		return ConfMessage{}, fmt.Errorf("status: %w", st.ErrInvalidValue)
	}

	return ConfMessage{Status: status}, nil
}
