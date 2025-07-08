// Package authorize provides the OCPP 1.6J Authorize message.
package authorize

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

// ConfirmationMessage represents the Authorize.conf PDU.
// This message is sent by the Central System to the Charge Point in response to an Authorize.req.
// It contains information about the authorization status of the IdTag.
// See OCPP 1.6J Part 2, Section 4.1.2.2 for details.
type ConfirmationMessage struct {
	IdTagInfo authorizetypes.IdTagInfoType // IdTagInfo is the information about the authorization status of the IdTag. (Required)
}

func Confirmation(input authorizetypes.ConfirmationPayload) (ConfirmationMessage, error) {
	err := input.Validate()
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf("authorize.Confirmation: invalid payload: %w", err)
	}

	info, err := authorizetypes.IdTagInfo(input.IdTagInfo)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf("authorize.Confirmation: invalid IdTagInfo: %w", err)
	}

	return ConfirmationMessage{IdTagInfo: info}, nil
}
