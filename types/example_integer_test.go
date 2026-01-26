package types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

func ExampleNewInteger() {
	intVal, err := st.NewInteger("73")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(intVal.Value())

	// Output:
	// 73
}

func ExampleNewInteger_invalid() {
	_, err := st.NewInteger("abc")
	if err != nil {
		fmt.Println("invalid integer input")
	}

	// Output:
	// invalid integer input
}
