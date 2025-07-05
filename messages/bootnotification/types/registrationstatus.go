package bootnotificationtypes

import (
	"fmt"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	RegistrationAccepted = "Accepted"
	RegistrationPending  = "Pending"
	RegistrationRejected = "Rejected"
)

type RegistrationStatusType struct {
	value string
}

func getRegistrationStatusSet() map[string]struct{} {
	return map[string]struct{}{
		RegistrationAccepted: {},
		RegistrationPending:  {},
		RegistrationRejected: {},
	}
}

func RegistrationStatus(input string) (RegistrationStatusType, error) {
	if _, exists := getRegistrationStatusSet()[input]; !exists {
		return RegistrationStatusType{}, fmt.Errorf("%w: %q", sharedtypes.ErrInvalidRegistrationStatus, input)
	}

	return RegistrationStatusType{value: input}, nil
}

func (s RegistrationStatusType) Value() string {
	return s.value
}
