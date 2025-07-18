package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

type ConfirmationMessage struct {
	IdTagInfo mat.IdTagInfo
}

func Confirmation(input mat.ConfirmationPayload) (ConfirmationMessage, error) {
	err := input.Validate()
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(
			"authorize.Confirmation: invalid payload: %w",
			err,
		)
	}

	info, err := mat.SetIdTagInfo(input.IdTagInfo)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(
			"authorize.Confirmation: invalid IdTagInfo: %w",
			err,
		)
	}

	return ConfirmationMessage{IdTagInfo: info}, nil
}
