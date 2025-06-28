package changeavailability

import (
	"fmt"

	changeavailabilitytypes "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationMessage struct {
	Status changeavailabilitytypes.ChangeAvailabilityStatusType
}

func Confirmation(input changeavailabilitytypes.ConfirmationPayload) (ConfirmationMessage, error) {
	if err := input.Validate(); err != nil {
		return ConfirmationMessage{}, fmt.Errorf("changeavailability.Confirmation: invalid payload: %w", err)
	}
	status, err := changeavailabilitytypes.ChangeAvailabilityStatus(input.Status)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse Status", err)
	}
	return ConfirmationMessage{Status: status}, nil
}
