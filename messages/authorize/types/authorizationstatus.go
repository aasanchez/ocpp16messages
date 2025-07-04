package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

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
		return AuthorizationStatus{}, fmt.Errorf("%w: %q", st.ErrInvalidAuthorizationStatus, input)
	}

	return AuthorizationStatus{value: input}, nil
}

func (s AuthorizationStatus) Value() string {
	return s.value
}
