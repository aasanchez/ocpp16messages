package types_test

//revive:disable line-length-limit
import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
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
		fmt.Println(err)
	}

	// Output:
	// strconv.ParseUint: parsing "abc": invalid syntax
}
