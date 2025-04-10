package types

// AuthorizationStatus represents the authorization status of an idTag.
type AuthorizationStatus string

const (
	// Identifier is allowed for charging.
	Accepted AuthorizationStatus = "Accepted"

	// Identifier has been blocked. Not allowed for charging.
	Blocked AuthorizationStatus = "Blocked"

	// Identifier has expired. Not allowed for charging.
	Expired AuthorizationStatus = "Expired"

	// Identifier is unknown. Not allowed for charging.
	Invalid AuthorizationStatus = "Invalid"

	// Identifier is already involved in another transaction and multiple transactions are not allowed. (Only relevant for a StartTransaction.req.)
	ConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid checks if the AuthorizationStatus is valid.
func (s AuthorizationStatus) IsValid() bool {
	switch s {
	case Accepted, Blocked, Expired:
		return true
	}
	return false
}
