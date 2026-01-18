package types

// RemoteStartTransactionStatus represents the result of a
// RemoteStartTransaction request.
type RemoteStartTransactionStatus string

// Accepted is the status value for an accepted remote start transaction.
const Accepted = "Accepted"

// Rejected is the status value for a rejected remote start transaction.
const Rejected = "Rejected"

// RemoteStartTransactionStatus enumeration values as defined in OCPP 1.6.
const (
	RemoteStartTransactionStatusAccepted RemoteStartTransactionStatus = Accepted
	RemoteStartTransactionStatusRejected RemoteStartTransactionStatus = Rejected
)

// IsValid checks if the RemoteStartTransactionStatus value is valid
// per OCPP 1.6.
func (r RemoteStartTransactionStatus) IsValid() bool {
	switch r {
	case RemoteStartTransactionStatusAccepted,
		RemoteStartTransactionStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RemoteStartTransactionStatus.
func (r RemoteStartTransactionStatus) String() string {
	return string(r)
}
