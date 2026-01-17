package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/meterValues/types"
)

// ExamplePhase demonstrates using the Phase enumeration constants
// defined in OCPP 1.6.
func ExamplePhase() {
	phase := types.L1
	fmt.Println("Phase:", phase.String())
	fmt.Println("IsValid:", phase.IsValid())
	// Output:
	// Phase: L1
	// IsValid: true
}

// ExamplePhase_allValues demonstrates all valid Phase values
// as defined in OCPP 1.6.
func ExamplePhase_allValues() {
	phases := []types.Phase{
		types.L1,
		types.L2,
		types.L3,
		types.N,
		types.L1N,
		types.L2N,
		types.L3N,
		types.L1L2,
		types.L2L3,
		types.L3L1,
	}

	for _, ph := range phases {
		fmt.Println(ph.String())
	}
	// Output:
	// L1
	// L2
	// L3
	// N
	// L1-N
	// L2-N
	// L3-N
	// L1-L2
	// L2-L3
	// L3-L1
}
