package authorize

import (
	"errors"
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
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at a time.
// Returns an error if:
//   - Status is not a valid AuthorizationStatus value
//   - ExpiryDate (if provided) is not a valid RFC3339 date
//   - ParentIdTag (if provided) exceeds 20 characters or contains invalid chars
func Conf(input ConfInput) (ConfMessage, error) {
	var errs []error

	var idTagInfo mat.IdTagInfo

	var expiryDate st.DateTime

	var parentIdToken mat.IdToken

	// Validate status (required)
	status := mat.AuthorizationStatus(input.Status)

	info, err := mat.NewIdTagInfo(status)
	if err != nil {
		errs = append(errs, fmt.Errorf("status: %w", err))
	} else {
		idTagInfo = info
	}

	// Validate expiryDate (optional)
	if input.ExpiryDate != nil {
		date, err := st.NewDateTime(*input.ExpiryDate)
		if err != nil {
			errs = append(errs, fmt.Errorf("expiryDate: %w", err))
		} else {
			expiryDate = date
		}
	}

	// Validate parentIdTag (optional)
	if input.ParentIdTag != nil {
		ciStr, err := st.NewCiString20Type(*input.ParentIdTag)
		if err != nil {
			errs = append(errs, fmt.Errorf("parentIdTag: %w", err))
		} else {
			token, err := mat.NewIdToken(ciStr)
			if err != nil {
				errs = append(errs, fmt.Errorf("parentIdTag: %w", err))
			} else {
				parentIdToken = token
			}
		}
	}

	// Return all accumulated errors
	if len(errs) > 0 {
		return ConfMessage{}, errors.Join(errs...)
	}

	// Build the message with validated fields
	if input.ExpiryDate != nil {
		idTagInfo = idTagInfo.WithExpiryDate(expiryDate)
	}

	if input.ParentIdTag != nil {
		idTagInfo = idTagInfo.WithParentIdTag(parentIdToken)
	}

	return ConfMessage{IdTagInfo: idTagInfo}, nil
}
