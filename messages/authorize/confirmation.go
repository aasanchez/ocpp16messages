package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfInput represents the raw input data for creating an Authorize.conf message.
// The constructor Conf validates all fields automatically.
type ConfInput struct {
	Status      string  // Required: AuthorizationStatus value
	ExpiryDate  *string // Optional: RFC3339 formatted expiry date
	ParentIdTag *string // Optional: Parent ID tag (max 20 chars)
}

// ConfMessage represents an OCPP 1.6 Authorize.conf message.
type ConfMessage struct {
	IdTagInfo mat.IdTagInfo
}

// Conf creates an Authorize.conf message from the given input.
// It validates all fields automatically and returns an error if:
//   - Status is not a valid AuthorizationStatus value
//   - ExpiryDate (if provided) is not a valid RFC3339 date
//   - ParentIdTag (if provided) exceeds 20 characters or contains invalid chars
func Conf(input ConfInput) (ConfMessage, error) {
	status := mat.AuthorizationStatus(input.Status)

	idTagInfo, err := mat.NewIdTagInfo(status)
	if err != nil {
		return ConfMessage{}, fmt.Errorf("idTagInfo: %w", err)
	}

	if input.ExpiryDate != nil {
		expiryDate, err := st.NewDateTime(*input.ExpiryDate)
		if err != nil {
			return ConfMessage{}, fmt.Errorf("expiryDate: %w", err)
		}

		idTagInfo = idTagInfo.WithExpiryDate(expiryDate)
	}

	if input.ParentIdTag != nil {
		ciStr, err := st.NewCiString20Type(*input.ParentIdTag)
		if err != nil {
			return ConfMessage{}, fmt.Errorf("parentIdTag: %w", err)
		}

		parentIdToken, err := mat.NewIdToken(ciStr)
		if err != nil {
			return ConfMessage{}, fmt.Errorf("parentIdTag: %w", err)
		}

		idTagInfo = idTagInfo.WithParentIdTag(parentIdToken)
	}

	return ConfMessage{IdTagInfo: idTagInfo}, nil
}
