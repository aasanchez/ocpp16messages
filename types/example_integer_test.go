package types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	exampleIntegerValid   = 73
	exampleIntegerInvalid = -1
)

func ExampleNewInteger() {
	intVal, err := st.NewInteger(exampleIntegerValid)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(intVal.Value())

	// Output:
	// 73
}

func ExampleNewInteger_invalid() {
	_, err := st.NewInteger(exampleIntegerInvalid)
	if err != nil {
		fmt.Println("invalid integer input")
	}

	// Output:
	// invalid integer input
}
