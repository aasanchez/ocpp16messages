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
		return ConfirmationMessage{}, fmt.Errorf("invalid idTagInfo: %w", err)
	}

	return ConfirmationMessage{
		IdTagInfo: info,
	}, nil
}

func (m ConfirmationMessage) Validate() error {
	if err := m.IdTagInfo.Validate(); err != nil {
		return fmt.Errorf("ConfirmationMessage validation failed: %w", err)
	}

	return nil
}

func (m ConfirmationMessage) String() string {
	return m.IdTagInfo.String()
}
