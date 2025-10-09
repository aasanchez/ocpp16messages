package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleSetDateTime() {
	datetime, err := st.SetDateTime("2025-08-30T14:34:56Z")
	if err != nil {
		fmt.Printf(st.ErrorUnexpectedError, err)

		return
	}

	fmt.Println("Time value:", datetime.Value().UTC())
	fmt.Println("String:", datetime.String())

	// Output:
	// Time value: 2025-08-30 14:34:56 +0000 UTC
	// String: 2025-08-30T14:34:56Z
}

func ExampleSetDateTime_invalid() {
	_, err := st.SetDateTime("not-a-time")
	if err != nil {
		fmt.Println("Invalid datetime")
	}

	// Output:
	// Invalid datetime
}
