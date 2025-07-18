package authorizetypes

import (
	"errors"
	"fmt"
)

var errInvalidAuthorizationStatus = errors.New("invalid AuthorizationStatus")

const (
	Accepted     = "Accepted"
	Blocked      = "Blocked"
	Expired      = "Expired"
	Invalid      = "Invalid"
	ConcurrentTx = "ConcurrentTx"
)

type AuthorizationStatus struct {
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

func SetAuthorizationStatus(input string) (AuthorizationStatus, error) {
	if _, exists := getStatusSet()[input]; !exists {
		return AuthorizationStatus{}, fmt.Errorf("%w: %q", errInvalidAuthorizationStatus, input)
	}

	return AuthorizationStatus{value: input}, nil
}

func (s AuthorizationStatus) Value() string {
	return s.value
}
