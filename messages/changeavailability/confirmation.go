package changeavailability

import (
	"fmt"

	cat "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationMessage struct {
	Status cat.ChangeAvailabilityStatusType
}

func Confirmation(input cat.ConfirmationPayload) (ConfirmationMessage, error) {
	if err := input.Validate(); err != nil {
		return ConfirmationMessage{}, fmt.Errorf("changeavailability.Confirmation: invalid payload: %w", err)
	}
	status, err := cat.ChangeAvailabilityStatus(input.Status)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse Status", err)
	}
	return ConfirmationMessage{Status: status}, nil
}
