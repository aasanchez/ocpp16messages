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

var statusSet = map[string]struct{}{
	Accepted:     {},
	Blocked:      {},
	Expired:      {},
	Invalid:      {},
	ConcurrentTx: {},
}

type AuthorizationStatusType struct {
	value string
}

func AuthorizationStatus(input string) (AuthorizationStatusType, error) {
	if _, exists := statusSet[input]; !exists {
		return AuthorizationStatusType{}, fmt.Errorf("%w: %q", ErrInvalidAuthorizationStatus, input)
	}
	return AuthorizationStatusType{value: input}, nil
}

func (s AuthorizationStatusType) Value() string {
	return s.value
}
