package changeavailabilitytypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	AvailabilityTypeOperative   = "Operative"
	AvailabilityTypeInoperative = "Inoperative"
)

type AvailabilityTypeType struct {
	value string
}

func getAvailabilityTypeSet() map[string]struct{} {
	return map[string]struct{}{
		AvailabilityTypeOperative:   {},
		AvailabilityTypeInoperative: {},
	}
}

// AvailabilityType converts a string to AvailabilityTypeType enforcing valid values.
func AvailabilityType(input string) (AvailabilityTypeType, error) {
	if _, exists := getAvailabilityTypeSet()[input]; !exists {
		return AvailabilityTypeType{}, fmt.Errorf("%w: %q", st.ErrInvalidChangeAvailabilityType, input)
	}

	return AvailabilityTypeType{value: input}, nil
}

// Value returns the raw string value of the AvailabilityTypeType.
func (t AvailabilityTypeType) Value() string {
	return t.value
}
