package types

// UnlockStatus string constants for use in enumeration values.
const (
	unlocked     = "Unlocked"
	unlockFailed = "UnlockFailed"
	notSupported = "NotSupported"
)

// UnlockStatus represents the result of an UnlockConnector request.
type UnlockStatus string

// UnlockStatus enumeration values as defined in OCPP 1.6.
const (
	// UnlockStatusUnlocked indicates the connector has unlocked successfully.
	UnlockStatusUnlocked UnlockStatus = unlocked
	// UnlockStatusUnlockFailed indicates the connector failed to unlock.
	UnlockStatusUnlockFailed UnlockStatus = unlockFailed
	// UnlockStatusNotSupported indicates the Charge Point has no connector
	// lock.
	UnlockStatusNotSupported UnlockStatus = notSupported
)

// IsValid checks if the UnlockStatus value is valid per OCPP 1.6.
func (u UnlockStatus) IsValid() bool {
	switch u {
	case UnlockStatusUnlocked,
		UnlockStatusUnlockFailed,
		UnlockStatusNotSupported:
		return true
	default:
		return false
	}
}

// String returns the string representation of UnlockStatus.
func (u UnlockStatus) String() string {
	return string(u)
}
