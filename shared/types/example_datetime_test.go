package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleSetDateTime() {
	dt, err := st.SetDateTime("2025-08-30T14:34:56Z")
	if err != nil {
		fmt.Println("unexpected error:", err)
		return
	}

	fmt.Println("Time value:", dt.Value().UTC())
	fmt.Println("String:", dt.String())

	// Output:
	// Time value: 2025-08-30 14:34:56 +0000 UTC
	// String: 2025-08-30T14:34:56Z
}

func ExampleSetDateTime_invalid() {
	_, err := st.SetDateTime("not-a-time")
	if err != nil {
		fmt.Println("error occurred")
	}

	// Output:
	// error occurred
}
