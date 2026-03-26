package types

// GetCompositeScheduleStatus represents the result of a GetCompositeSchedule
// request as defined in OCPP 1.6.
type GetCompositeScheduleStatus string

// GetCompositeScheduleStatus enumeration values as defined in OCPP 1.6.
const (
	GetCompositeScheduleStatusAccepted GetCompositeScheduleStatus = "Accepted"
	GetCompositeScheduleStatusRejected GetCompositeScheduleStatus = "Rejected"
)

// IsValid checks if the GetCompositeScheduleStatus value is valid per OCPP 1.6.
func (g GetCompositeScheduleStatus) IsValid() bool {
	switch g {
	case GetCompositeScheduleStatusAccepted,
		GetCompositeScheduleStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of GetCompositeScheduleStatus.
func (g GetCompositeScheduleStatus) String() string {
	return string(g)
}
