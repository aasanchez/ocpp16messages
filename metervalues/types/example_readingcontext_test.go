package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/metervalues/types"
)

// ExampleReadingContext demonstrates using the ReadingContext enumeration
// constants defined in OCPP 1.6.
func ExampleReadingContext() {
	context := types.SamplePeriodic
	fmt.Println("Context:", context.String())
	fmt.Println("IsValid:", context.IsValid())
	// Output:
	// Context: Sample.Periodic
	// IsValid: true
}

// ExampleReadingContext_allValues demonstrates all valid ReadingContext values
// as defined in OCPP 1.6.
func ExampleReadingContext_allValues() {
	contexts := []types.ReadingContext{
		types.InterruptionBegin,
		types.InterruptionEnd,
		types.Other,
		types.SampleClock,
		types.SamplePeriodic,
		types.TransactionBegin,
		types.TransactionEnd,
		types.Trigger,
	}

	for _, ctx := range contexts {
		fmt.Println(ctx.String())
	}
	// Output:
	// Interruption.Begin
	// Interruption.End
	// Other
	// Sample.Clock
	// Sample.Periodic
	// Transaction.Begin
	// Transaction.End
	// Trigger
}
