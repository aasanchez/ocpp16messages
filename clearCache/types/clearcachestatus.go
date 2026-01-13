package types

// ClearCacheStatus represents the result of a ClearCache request.
type ClearCacheStatus string

// ClearCacheStatus enumeration values as defined in OCPP 1.6.
const (
	ClearCacheStatusAccepted ClearCacheStatus = "Accepted"
	ClearCacheStatusRejected ClearCacheStatus = "Rejected"
)

// IsValid checks if the ClearCacheStatus value is valid per OCPP 1.6.
func (c ClearCacheStatus) IsValid() bool {
	switch c {
	case ClearCacheStatusAccepted,
		ClearCacheStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of ClearCacheStatus.
func (c ClearCacheStatus) String() string {
	return string(c)
}
