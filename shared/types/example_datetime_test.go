package types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleNewDateTime() {
	datetime, err := st.NewDateTime("2025-08-30T14:34:56Z")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Time value:", datetime.Value().UTC())

	// Output:
	// Time value: 2025-08-30 14:34:56 +0000 UTC
}

func ExampleNewDateTime_invalid() {
	_, err := st.NewDateTime("not-a-time")
	if err != nil {
		fmt.Println("cannot parse \"not-a-time\" as date")
	}

	// Output:
	// cannot parse "not-a-time" as date
}
