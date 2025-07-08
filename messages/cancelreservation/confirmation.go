package cancelreservation

import (
	"fmt"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationMessage struct {
	Status cancelreservationtypes.CancelReservationStatusType
}

func Confirmation(input cancelreservationtypes.ConfirmationPayload) (ConfirmationMessage, error) {
	err := input.Validate()
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf("bootnotificationtypes.Confirmation: invalid payload: %w", err)
	}

	status, err := cancelreservationtypes.CancelReservationStatus(input.Status)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse Status", err)
	}

	return ConfirmationMessage{
		Status: status,
	}, nil
}
