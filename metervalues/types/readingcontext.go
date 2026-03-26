package types

// ReadingContext represents the context of a meter value reading
// as defined in OCPP 1.6.
type ReadingContext string

// Alias for shorter constant names within this package.
type rc = ReadingContext

// ReadingContext enumeration values as defined in OCPP 1.6.
const (
	// InterruptionBegin indicates the value was taken at the beginning
	// of an interruption.
	InterruptionBegin rc = "Interruption.Begin"
	// InterruptionEnd indicates the value was taken at the end
	// of an interruption.
	InterruptionEnd rc = "Interruption.End"
	// Other indicates an unspecified reading context.
	Other rc = "Other"
	// SampleClock indicates the value was taken at a clock-aligned interval.
	SampleClock rc = "Sample.Clock"
	// SamplePeriodic indicates the value was taken at a periodic interval.
	SamplePeriodic rc = "Sample.Periodic"
	// TransactionBegin indicates the value was taken at the start
	// of a transaction.
	TransactionBegin rc = "Transaction.Begin"
	// TransactionEnd indicates the value was taken at the end of a transaction.
	TransactionEnd rc = "Transaction.End"
	// Trigger indicates the value was taken because of a trigger.
	Trigger rc = "Trigger"
)

// IsValid checks if the ReadingContext value is valid per OCPP 1.6.
func (r ReadingContext) IsValid() bool {
	switch r {
	case InterruptionBegin,
		InterruptionEnd,
		Other,
		SampleClock,
		SamplePeriodic,
		TransactionBegin,
		TransactionEnd,
		Trigger:
		return true
	default:
		return false
	}
}

// String returns the string representation of ReadingContext.
func (r ReadingContext) String() string {
	return string(r)
}
