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

func ExampleCiString20Type() {
	input := strings.Repeat("A", 20)

	cistr, err := types.SetCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))

	// Output:
	// Length: 20
}

func ExampleCiString20Type_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := types.SetCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 21, max 20)
}

func ExampleCiString25Type() {
	input := strings.Repeat("A", 25)

	cistr, err := types.SetCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 25
}

func ExampleCiString25Type_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := types.SetCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 26, max 25)
}

func ExampleCiString50Type() {
	input := strings.Repeat("A", 50)

	cistr, err := types.SetCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 50
}

func ExampleCiString50Type_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := types.SetCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 51, max 50)
}

func ExampleCiString255Type() {
	input := strings.Repeat("A", 255)

	cistr, err := types.SetCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 255
}

func ExampleCiString255Type_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := types.SetCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 256, max 255)
}

func ExampleCiString500Type() {
	input := strings.Repeat("A", 499)

	cistr, err := types.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 499
}

func ExampleCiString500Type_invalid() {
	input := strings.Repeat("A", 501)

	cistr, err := types.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 501, max 500)
}
