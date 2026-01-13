package types

// ConfigurationStatus represents the result of a ChangeConfiguration request
// as defined in OCPP 1.6.
type ConfigurationStatus string

// ConfigurationStatus enumeration values as defined in OCPP 1.6.
const (
	ConfigurationStatusAccepted       ConfigurationStatus = "Accepted"
	ConfigurationStatusRejected       ConfigurationStatus = "Rejected"
	ConfigurationStatusRebootRequired ConfigurationStatus = "RebootRequired"
	ConfigurationStatusNotSupported   ConfigurationStatus = "NotSupported"
)

// IsValid checks if the ConfigurationStatus value is valid per OCPP 1.6.
func (c ConfigurationStatus) IsValid() bool {
	switch c {
	case ConfigurationStatusAccepted,
		ConfigurationStatusRejected,
		ConfigurationStatusRebootRequired,
		ConfigurationStatusNotSupported:
		return true
	default:
		return false
	}
}

// String returns the string representation of ConfigurationStatus.
func (c ConfigurationStatus) String() string {
	return string(c)
}
