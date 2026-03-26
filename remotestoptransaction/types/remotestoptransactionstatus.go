package types

// Accepted is the status value for an accepted remote stop transaction.
const Accepted = "Accepted"

// Rejected is the status value for a rejected remote stop transaction.
const Rejected = "Rejected"

// RemoteStopTransactionStatus represents the result of a
// RemoteStopTransaction request.
type RemoteStopTransactionStatus string

// RemoteStopTransactionStatus enumeration values as defined in OCPP 1.6.
const (
	RemoteStopTransactionStatusAccepted RemoteStopTransactionStatus = Accepted
	RemoteStopTransactionStatusRejected RemoteStopTransactionStatus = Rejected
)

// IsValid checks if the RemoteStopTransactionStatus value is valid
// per OCPP 1.6.
func (r RemoteStopTransactionStatus) IsValid() bool {
	switch r {
	case RemoteStopTransactionStatusAccepted,
		RemoteStopTransactionStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RemoteStopTransactionStatus.
func (r RemoteStopTransactionStatus) String() string {
	return string(r)
}
