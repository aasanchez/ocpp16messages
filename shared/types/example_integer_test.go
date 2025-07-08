package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ExampleInteger demonstrates parsing a valid integer string.
func ExampleInteger() {
	input := "123"

	val, err := st.SetInteger(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Value:", val.Value())

	// Output:
	// Value: 123
}

// ExampleInteger_invalid demonstrates handling of an invalid integer string.
func ExampleInteger_invalid() {
	input := "abc"

	_, err := st.SetInteger(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output:
	// Error: invalid integer: strconv.ParseUint: parsing "abc": invalid syntax
}
