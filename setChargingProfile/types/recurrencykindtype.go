package types

// RecurrencyKindType represents the kind of recurrency for a recurring
// charging profile as defined in OCPP 1.6.
type RecurrencyKindType string

// RecurrencyKindType enumeration values as defined in OCPP 1.6.
const (
	// RecurrencyKindDaily indicates the schedule repeats every 24 hours.
	RecurrencyKindDaily RecurrencyKindType = "Daily"
	// RecurrencyKindWeekly indicates the schedule repeats every 7 days
	// starting from startSchedule.
	RecurrencyKindWeekly RecurrencyKindType = "Weekly"
)

// IsValid checks if the RecurrencyKindType value is valid per OCPP 1.6.
func (r RecurrencyKindType) IsValid() bool {
	switch r {
	case RecurrencyKindDaily,
		RecurrencyKindWeekly:
		return true
	default:
		return false
	}
}

// String returns the string representation of RecurrencyKindType.
func (r RecurrencyKindType) String() string {
	return string(r)
}
