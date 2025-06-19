package authorizetypes

import (
	"fmt"
)

const (
	Accepted     = "Accepted"
	Blocked      = "Blocked"
	Expired      = "Expired"
	Invalid      = "Invalid"
	ConcurrentTx = "ConcurrentTx"
)

type AuthorizationStatusType struct {
	value string
}

func getStatusSet() map[string]struct{} {
	return map[string]struct{}{
		Accepted:     {},
		Blocked:      {},
		Expired:      {},
		Invalid:      {},
		ConcurrentTx: {},
	}
}

func AuthorizationStatus(input string) (AuthorizationStatusType, error) {
	if _, exists := getStatusSet()[input]; !exists {
		return AuthorizationStatusType{}, fmt.Errorf("%w: %q", ErrInvalidAuthorizationStatus, input)
	}

	return AuthorizationStatusType{value: input}, nil
}

func (s AuthorizationStatusType) Value() string {
	return s.value
}
