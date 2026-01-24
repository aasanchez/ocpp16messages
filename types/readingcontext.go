package types

// ReadingContext represents the context of a meter value reading
// as defined in OCPP 1.6.
type ReadingContext string

// Alias for shorter constant names within this package.
type rc = ReadingContext

// ReadingContext enumeration values as defined in OCPP 1.6.
const (
	// ReadingContextInterruptionBegin indicates the value was taken at the
	// beginning of an interruption.
	ReadingContextInterruptionBegin rc = "Interruption.Begin"
	// ReadingContextInterruptionEnd indicates the value was taken at the end
	// of an interruption.
	ReadingContextInterruptionEnd rc = "Interruption.End"
	// ReadingContextOther indicates an unspecified reading context.
	ReadingContextOther rc = "Other"
	// ReadingContextSampleClock indicates the value was taken at a
	// clock-aligned interval.
	ReadingContextSampleClock rc = "Sample.Clock"
	// ReadingContextSamplePeriodic indicates the value was taken at a
	// periodic interval.
	ReadingContextSamplePeriodic rc = "Sample.Periodic"
	// ReadingContextTransactionBegin indicates the value was taken at the
	// start of a transaction.
	ReadingContextTransactionBegin rc = "Transaction.Begin"
	// ReadingContextTransactionEnd indicates the value was taken at the end
	// of a transaction.
	ReadingContextTransactionEnd rc = "Transaction.End"
	// ReadingContextTrigger indicates the value was taken because of a trigger.
	ReadingContextTrigger rc = "Trigger"
)

// IsValid checks if the ReadingContext value is valid per OCPP 1.6.
func (r ReadingContext) IsValid() bool {
	switch r {
	case ReadingContextInterruptionBegin,
		ReadingContextInterruptionEnd,
		ReadingContextOther,
		ReadingContextSampleClock,
		ReadingContextSamplePeriodic,
		ReadingContextTransactionBegin,
		ReadingContextTransactionEnd,
		ReadingContextTrigger:
		return true
	default:
		return false
	}
}

// String returns the string representation of ReadingContext.
func (r ReadingContext) String() string {
	return string(r)
}
