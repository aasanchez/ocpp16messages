package types_test

import (
	"fmt"
	"strings"

	types "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	msgError  = "Error:"
	msgLength = "Length:"
)

func ExampleCiString20() {
	input := strings.Repeat("A", 20)

	cistr, err := types.SetCiString20(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))

	// Output:
	// Length: 20
}

func ExampleCiString20_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := types.SetCiString20(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: validation error: the value exceeds the maximum allowed character length for this field (got length 21, max 20)
}

func ExampleCiString25() {
	input := strings.Repeat("A", 25)

	cistr, err := types.SetCiString25(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 25
}

func ExampleCiString25_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := types.SetCiString25(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: validation error: the value exceeds the maximum allowed character length for this field (got length 26, max 25)
}

func ExampleCiString50() {
	input := strings.Repeat("A", 50)

	cistr, err := types.SetCiString50(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 50
}

func ExampleCiString50_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := types.SetCiString50(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: validation error: the value exceeds the maximum allowed character length for this field (got length 51, max 50)
}

func ExampleCiString255() {
	input := strings.Repeat("A", 255)

	cistr, err := types.SetCiString255(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 255
}

func ExampleCiString255_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := types.SetCiString255(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: validation error: the value exceeds the maximum allowed character length for this field (got length 256, max 255)
}

func ExampleCiString500() {
	input := strings.Repeat("A", 499)

	cistr, err := types.SetCiString500(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 499
}

func ExampleCiString500_invalid() {
	input := strings.Repeat("A", 501)

	cistr, err := types.SetCiString500(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: validation error: the value exceeds the maximum allowed character length for this field (got length 501, max 500)
}
