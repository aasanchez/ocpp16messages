package bootnotification

import (
	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// ConfirmationMessage represents a BootNotification confirmation payload.
type ConfirmationMessage struct {
	CurrentTime             sharedtypes.DateTimeType
	ChargePointSerialNumber sharedtypes.CiString25Type
}

// confirmation is not currently used; implementation pending.
func confirmation(input bootnotificationtypes.ConfirmationPayload) {
	// TODO: implement; placeholder only.
}
