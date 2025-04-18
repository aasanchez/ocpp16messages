package types_test

import (
	"fmt"
	"time"

	types "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleDateTime_valid() {
	raw := "2027-04-12T10:03:04-04:00"

	datetime, err := types.DateTime(raw)
	if err != nil {
		fmt.Printf("unexpected error: %v\n", err)
		return
	}

	if err := datetime.Validate(); err != nil {
		fmt.Printf("validation failed: %v\n", err)
		return
	}

	utc := datetime.Value().UTC().Format(time.RFC3339)
	fmt.Println("Parsed UTC time:", utc)

	// Output:
	// Parsed UTC time: 2027-04-12T14:03:04Z
}

func ExampleDateTime_invalidFormat() {
	_, err := types.DateTime("not-a-date")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output:
	// Error: parsing time "not-a-date" as "2006-01-02T15:04:05Z07:00": cannot parse "not-a-date" as "2006"
}
