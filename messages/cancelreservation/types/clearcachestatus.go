package cancelreservationtypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	ClearCacheAccepted = "Accepted"
	ClearCacheRejected = "Rejected"
)

type ClearCacheStatusType struct {
	value string
}

func getClearCacheStatusSet() map[string]struct{} {
	return map[string]struct{}{
		ClearCacheAccepted: {},
		ClearCacheRejected: {},
	}
}

func ClearCacheStatus(input string) (ClearCacheStatusType, error) {
	if _, exists := getClearCacheStatusSet()[input]; !exists {
		return ClearCacheStatusType{}, fmt.Errorf("%w: %q", sharedtypes.ErrInvalidClearCacheStatus, input)
	}

	return ClearCacheStatusType{value: input}, nil
}

func (s ClearCacheStatusType) Value() string {
	return s.value
}
