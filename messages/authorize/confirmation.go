package authorize

import (
	"fmt"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

type ConfirmationMessage struct {
	IdTagInfo authorizetypes.IdTagInfoType
}

func Confirmation(info authorizetypes.IdTagInfoType) (ConfirmationMessage, error) {
	if err := info.Validate(); err != nil {
		return ConfirmationMessage{}, fmt.Errorf("confirmation: invalid IdTagInfo: %w", err)
	}

	return ConfirmationMessage{
		IdTagInfo: info,
	}, nil
}

func (m ConfirmationMessage) Validate() error {
	if err := m.IdTagInfo.Validate(); err != nil {
		return fmt.Errorf("confirmationMessage.Validate: %w", err)
	}

	return nil
}
