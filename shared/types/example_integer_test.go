package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleSetInteger() {
	intVal, err := st.SetInteger("73")
	if err != nil {
		fmt.Println("unexpected error while parsing a valid integer:", err)
		return
	}

	fmt.Println("Value is:", intVal.Value())

	// Output:
	// Value is: 73
}

func ExampleSetInteger_invalid() {
	_, err := st.SetInteger("abc")
	if err != nil {
		fmt.Println("failed to parse integer from input 'abc':", err)
	}

	// Output:
	// failed to parse integer from input 'abc': invalid Integer: strconv.ParseUint: parsing "abc": invalid syntax
}
