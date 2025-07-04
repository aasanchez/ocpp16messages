package cancelreservationtypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	CancelReservationAccepted = "Accepted"
	CancelReservationRejected = "Rejected"
)

type CancelReservationStatusType struct {
	value string
}

func getCancelReservationStatusSet() map[string]struct{} {
	return map[string]struct{}{
		CancelReservationAccepted: {},
		CancelReservationRejected: {},
	}
}

func CancelReservationStatus(input string) (CancelReservationStatusType, error) {
	if _, exists := getCancelReservationStatusSet()[input]; !exists {
		return CancelReservationStatusType{}, fmt.Errorf("%w: %q", sharedtypes.ErrInvalidCancelReservationStatus, input)
	}

	return CancelReservationStatusType{value: input}, nil
}

func (s CancelReservationStatusType) Value() string {
	return s.value
}
