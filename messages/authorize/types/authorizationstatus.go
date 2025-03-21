package authorizetypes

// import (
// 	"fmt"

// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// const (
// 	// Accepted indicates that the IdTag is accepted.
// 	Accepted = "Accepted"
// 	// Blocked indicates that the IdTag is blocked.
// 	Blocked = "Blocked"
// 	// Expired indicates that the IdTag has expired.
// 	Expired = "Expired"
// 	// Invalid indicates that the IdTag is invalid.
// 	Invalid = "Invalid"
// 	// ConcurrentTx indicates that the IdTag is valid, but already in use.
// 	ConcurrentTx = "ConcurrentTx"
// )

// // AuthorizationStatusType is a string type that represents the authorization status of an IdTag.
// // It is used in the Authorize.conf message.
// // See OCPP 1.6J Part 2, Section 4.1.2.2, Table 13 for more details.
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

// // AuthorizationStatus creates a new AuthorizationStatusType from a string.
// // It validates the input string against the predefined set of valid authorization statuses.
// //
// // Parameters:
// //
// //	input: The string representation of the authorization status.
// //
// // Returns:
// //
// //	AuthorizationStatusType: A new AuthorizationStatusType if the input is valid.
// //	error: An error if the input string is not a valid authorization status.
// func SetAuthorizationStatus(input string) (AuthorizationStatus, error) {
// 	if _, exists := getStatusSet()[input]; !exists {
// 		return AuthorizationStatus{}, fmt.Errorf(
// 			"%w: %q",
// 			sharedtypes.ErrInvalidAuthorizationStatus,
// 			input,
// 		)
// 	}

// 	return AuthorizationStatus{value: input}, nil
// }

// // Value returns the string representation of the AuthorizationStatusType.
// func (s AuthorizationStatus) Value() string {
// 	return s.value
// }
