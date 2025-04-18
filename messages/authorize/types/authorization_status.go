package authorizetypes

import (
	"errors"
	"fmt"
)

const (
	Accepted     = "Accepted"
	Blocked      = "Blocked"
	Expired      = "Expired"
	Invalid      = "Invalid"
	ConcurrentTx = "ConcurrentTx"
)

var (
	ErrInvalidAuthorizationStatus = errors.New("invalid authorization status")
)

var authorizationStatusSet = map[string]struct{}{
	Accepted:     {},
	Blocked:      {},
	Expired:      {},
	Invalid:      {},
	ConcurrentTx: {},
}

type AuthorizationStatusType struct {
	value string
}

func AuthorizationStatus(s string) (AuthorizationStatusType, error) {
	if _, ok := authorizationStatusSet[s]; !ok {
		return AuthorizationStatusType{}, fmt.Errorf("%w: %q", ErrInvalidAuthorizationStatus, s)
	}
	return AuthorizationStatusType{value: s}, nil
}

func (s AuthorizationStatusType) Validate() error {
	if _, ok := authorizationStatusSet[s.value]; !ok {
		return fmt.Errorf("%w: %q", ErrInvalidAuthorizationStatus, s.value)
	}
	return nil
}
