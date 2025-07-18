package authorizetypes

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

var errInvalidAuthorizationStatus = errors.New("invalid AuthorizationStatus")

const (
	// Identifier is allowed for charging.
	Accepted = "Accepted"
	// Identifier has been blocked. Not allowed for charging.
	Blocked = "Blocked"
	// Identifier has expired. Not allowed for charging.
	Expired = "Expired"
	// Identifier is unknown. Not allowed for charging.
	Invalid = "Invalid"
	// Identifier is already involved in another transaction and multiple transactions are not allowed. (Only relevant for a StartTransaction.req.)
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
		return AuthorizationStatus{}, fmt.Errorf(
			st.ErrorWrapper,
			errInvalidAuthorizationStatus,
			input,
		)
	}

	return AuthorizationStatus{value: input}, nil
}

func (s AuthorizationStatus) Value() string {
	return s.value
}
