package authorize

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// ConfirmationMessage represents the OCPP 1.6J Authorize.conf message.
//
// This message is returned by the Central System (CSMS) in response to an Authorize.req,
// indicating whether the provided idTag is authorized to begin a charging session.
//
// The confirmation includes the authorization status and optionally an expiry date
// or a parent idTag reference.
//
// Specification Reference:
// - OCPP 1.6J, Section 5.2: Authorize.conf
type ConfirmationMessage struct {
	// IdTagInfo contains the authorization result and metadata.
	//
	// This field is mandatory in the Authorize.conf response. It encapsulates the authorization
	// status (Accepted, Blocked, etc.) and may include additional fields like expiryDate
	// or a reference to a parentIdTag.
	IdTagInfo types.IdTagInfoType
}

// Confirmation constructs a new ConfirmationMessage.
//
// It accepts a validated IdTagInfoType. If the input is invalid, an error is returned.
// This function ensures correctness before serializing or dispatching the message.
func Confirmation(info types.IdTagInfoType) (ConfirmationMessage, error) {
	if err := info.Validate(); err != nil {
		return ConfirmationMessage{}, fmt.Errorf("invalid idTagInfo: %w", err)
	}

	return ConfirmationMessage{
		IdTagInfo: info,
	}, nil
}

// Validate performs internal validation on the ConfirmationMessage.
//
// It re-validates the contents of the message to ensure compliance with OCPP 1.6J,
// and is typically used after decoding messages from a transport layer (e.g., JSON or SOAP).
func (m ConfirmationMessage) Validate() error {
	if err := m.IdTagInfo.Validate(); err != nil {
		return fmt.Errorf("ConfirmationMessage validation failed: %w", err)
	}

	return nil
}

// String returns a human-readable representation of the ConfirmationMessage.
//
// Useful for debugging, logging, or visual inspection during development and testing.
func (m ConfirmationMessage) String() string {
	return fmt.Sprintf("Authorize.conf{%s}", m.IdTagInfo.String())
}
