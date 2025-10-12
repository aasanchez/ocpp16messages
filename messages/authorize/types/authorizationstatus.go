package authorizetypes

// import (
// 	"errors"
// 	"fmt"

// 	st "github.com/aasanchez/ocpp16messages/shared/types"
// )

// var errInvalidAuthorizationStatus = errors.New("invalid AuthorizationStatus")

// const (
// 	Accepted     = "Accepted"     // is allowed for charging.
// 	Blocked      = "Blocked"      // has been blocked. Not allowed for charging.
// 	Expired      = "Expired"      // has expired. Not allowed for charging.
// 	Invalid      = "Invalid"      // is unknown. Not allowed for charging.
// 	ConcurrentTx = "ConcurrentTx" // "ID in use, multiple tx not allowed."
// )

// type AuthorizationStatus struct {
// 	value string
// }

// func getStatusSet() map[string]struct{} {
// 	return map[string]struct{}{
// 		Accepted:     {},
// 		Blocked:      {},
// 		Expired:      {},
// 		Invalid:      {},
// 		ConcurrentTx: {},
// 	}
// }

// func SetAuthorizationStatus(input string) (AuthorizationStatus, error) {
// 	if _, exists := getStatusSet()[input]; !exists {
// 		return AuthorizationStatus{}, fmt.Errorf(
// 			st.ErrorWrapper,
// 			errInvalidAuthorizationStatus,
// 			input,
// 		)
// 	}

// 	return AuthorizationStatus{value: input}, nil
// }

// func (s AuthorizationStatus) Value() string {
// 	return s.value
// }
