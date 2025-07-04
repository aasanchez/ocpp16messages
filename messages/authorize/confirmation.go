package authorize

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

type ConfirmationMessage struct {
	IdTagInfo authorizetypes.IdTagInfoType
}

func Confirmation(input authorizetypes.ConfirmationPayload) (ConfirmationMessage, error) {
	if err := input.Validate(); err != nil {
		return ConfirmationMessage{}, fmt.Errorf("authorize.Confirmation: invalid payload: %w", err)
	}

	info, err := authorizetypes.IdTagInfo(input.IdTagInfo)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf("authorize.Confirmation: invalid IdTagInfo: %w", err)
	}

	return ConfirmationMessage{IdTagInfo: info}, nil
}
