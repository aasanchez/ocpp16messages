package bootnotification

import (
	"fmt"
	"strconv"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type ConfirmationMessage struct {
	CurrentTime sharedtypes.DateTimeType
	Interval    uint32
	Status      bootnotificationtypes.RegistrationStatusType
}

func Confirmation(input bootnotificationtypes.ConfirmationPayload) (ConfirmationMessage, error) {
	if err := input.Validate(); err != nil {
		return ConfirmationMessage{}, fmt.Errorf("bootnotificationtypes.Confirmation: invalid payload: %w", err)
	}

	currentTime, err := sharedtypes.DateTime(input.CurrentTime)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse CurrentTime", err)
	}

	intervalTemp, err := strconv.ParseUint(input.Interval, 10, 32)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse Interval", err)
	}
	var interval = uint32(intervalTemp)

	status, err := bootnotificationtypes.RegistrationStatus(input.Status)
	if err != nil {
		return ConfirmationMessage{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse Status", err)
	}

	return ConfirmationMessage{
		CurrentTime: currentTime,
		Interval:    interval,
		Status:      status,
	}, nil
}
