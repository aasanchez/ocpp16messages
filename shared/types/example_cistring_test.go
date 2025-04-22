package sharedtypes_test

import (
	"fmt"
	"strings"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleCiString() {
	input := "Hello, OCPP!"
	maxLen := 20

	cs, err := sharedtypes.CiString(input, maxLen)
	if err != nil {
		fmt.Printf("Failed to create ciString: %v", err)
	}

	fmt.Println("ciString created:", cs.String())

	// Output:
	// ciString created: Hello, OCPP!
}

func ExampleCiString20() {
	input := strings.Repeat("A", 20)

	cistr, err := sharedtypes.CiString20(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Length: 20
}

func ExampleCiString20_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := sharedtypes.CiString20(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: ocpp16messages/shared/types/cistring.validate: value exceeds maximum allowed length: actual length 21, max 20
}

func ExampleCiString25() {
	input := strings.Repeat("A", 25)

	cistr, err := sharedtypes.CiString25(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Length: 25
}

func ExampleCiString25_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := sharedtypes.CiString25(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: ocpp16messages/shared/types/cistring.validate: value exceeds maximum allowed length: actual length 26, max 25
}

func ExampleCiString50() {
	input := strings.Repeat("A", 50)

	cistr, err := sharedtypes.CiString50(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Length: 50
}

func ExampleCiString50_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := sharedtypes.CiString50(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: ocpp16messages/shared/types/cistring.validate: value exceeds maximum allowed length: actual length 51, max 50
}

func ExampleCiString255() {
	input := strings.Repeat("A", 255)

	cistr, err := sharedtypes.CiString255(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Length: 255
}

func ExampleCiString255_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := sharedtypes.CiString255(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: ocpp16messages/shared/types/cistring.validate: value exceeds maximum allowed length: actual length 256, max 255
}

func ExampleCiString500() {
	input := strings.Repeat("A", 499)

	cistr, err := sharedtypes.CiString500(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Length: 499
}

func ExampleCiString500_invalid() {
	input := strings.Repeat("A", 501)

	cistr, err := sharedtypes.CiString500(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: ocpp16messages/shared/types/cistring.validate: value exceeds maximum allowed length: actual length 501, max 500
}
