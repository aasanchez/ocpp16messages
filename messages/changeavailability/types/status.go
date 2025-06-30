package changeavailabilitytypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	ChangeAvailabilityStatusAccepted = "Accepted"
	ChangeAvailabilityStatusRejected = "Rejected"
)

type ChangeAvailabilityStatusType struct {
	value string
}

func getChangeAvailabilityStatusSet() map[string]struct{} {
	return map[string]struct{}{
		ChangeAvailabilityStatusAccepted: {},
		ChangeAvailabilityStatusRejected: {},
	}
}

// ChangeAvailabilityStatus converts a string to ChangeAvailabilityStatusType enforcing valid values.
func ChangeAvailabilityStatus(input string) (ChangeAvailabilityStatusType, error) {
	if _, exists := getChangeAvailabilityStatusSet()[input]; !exists {
		return ChangeAvailabilityStatusType{}, fmt.Errorf("%w: %q", st.ErrInvalidChangeAvailabilityStatus, input)
	}

	return ChangeAvailabilityStatusType{value: input}, nil
}

// Value returns the raw string value of the ChangeAvailabilityStatusType.
func (t ChangeAvailabilityStatusType) Value() string {
	return t.value
}
