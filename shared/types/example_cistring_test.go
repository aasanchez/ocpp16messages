package sharedtypes_test

import (
	"fmt"
	"strings"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	msgError  = "Error:"
	msgLength = "Length:"
)

func ExampleSetCiString20Type() {
	input := strings.Repeat("A", 20)

	cistr, err := st.SetCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))

	// Output:
	// Length: 20
}

func ExampleSetCiString20Type_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := st.SetCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString len=21 > max=20 (error): exceeds maximum length
}

func ExampleSetCiString25Type() {
	input := strings.Repeat("A", 25)

	cistr, err := st.SetCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 25
}

func ExampleSetCiString25Type_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := st.SetCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString len=26 > max=25 (error): exceeds maximum length
}

func ExampleSetCiString50Type() {
	input := strings.Repeat("A", 50)

	cistr, err := st.SetCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 50
}

func ExampleSetCiString50Type_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := st.SetCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString len=51 > max=50 (error): exceeds maximum length
}

func ExampleSetCiString255Type() {
	input := strings.Repeat("A", 255)

	cistr, err := st.SetCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 255
}

func ExampleSetCiString255Type_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := st.SetCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString len=256 > max=255 (error): exceeds maximum length
}

func ExampleSetCiString500Type() {
	input := strings.Repeat("A", 499)

	cistr, err := st.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 499
}

func ExampleSetCiString500Type_invalid() {
	input := strings.Repeat("A", 501)

	cistr, err := st.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString len=501 > max=500 (error): exceeds maximum length
}
