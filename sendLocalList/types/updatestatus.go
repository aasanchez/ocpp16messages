package types

// UpdateStatus enumeration values as defined in OCPP 1.6.
const (
	// Accepted indicates the local authorization list update was accepted.
	Accepted = "Accepted"
	// Failed indicates the local authorization list update failed.
	Failed = "Failed"
	// NotSupported indicates the Charge Point does not support this feature.
	NotSupported = "NotSupported"
	// VersionMismatch indicates a version number mismatch.
	VersionMismatch = "VersionMismatch"
)

// UpdateStatus represents the result of a SendLocalList request.
type UpdateStatus string

// UpdateStatus enumeration values as defined in OCPP 1.6.
const (
	UpdateStatusAccepted        UpdateStatus = Accepted
	UpdateStatusFailed          UpdateStatus = Failed
	UpdateStatusNotSupported    UpdateStatus = NotSupported
	UpdateStatusVersionMismatch UpdateStatus = VersionMismatch
)

// IsValid checks if the UpdateStatus value is valid per OCPP 1.6.
func (u UpdateStatus) IsValid() bool {
	switch u {
	case UpdateStatusAccepted,
		UpdateStatusFailed,
		UpdateStatusNotSupported,
		UpdateStatusVersionMismatch:
		return true
	default:
		return false
	}
}

// String returns the string representation of UpdateStatus.
func (u UpdateStatus) String() string {
	return string(u)
}
