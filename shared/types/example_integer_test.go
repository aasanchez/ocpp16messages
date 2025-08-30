package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleSetInteger() {
	intVal, err := st.SetInteger("73")
	if err != nil {
		fmt.Println("unexpected error:", err)

		return
	}

	fmt.Println("Value is:", intVal.Value())

	// Output:
	// Value is: 73
}

func ExampleSetInteger_invalid() {
	_, err := st.SetInteger("abc")
	if err != nil {
		fmt.Println(err != nil)
	}

	// Output:
	// true
}
