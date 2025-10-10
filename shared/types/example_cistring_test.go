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

// Example: OCPP 1.6 CiString20.
//
// Context
//   - Printable ASCII only (32..126)
//   - Max length 20, empty allowed
//   - Common for short IDs (e.g. IdTag)
//
// What this shows
//   - Construct a valid CiString20
//   - Read back its length
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

// Example: CiString20 invalid length.
//
// What this shows
//   - Input longer than 20 fails
//   - Error message includes max length
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

// Example: OCPP 1.6 CiString25.
//
// Context
//   - Printable ASCII (32..126)
//   - Max length 25, empty allowed
//   - Use for vendor/model fields
//
// What this shows
//   - Construct a valid CiString25
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

// Example: CiString25 invalid length.
//
// What this shows
//   - Input longer than 25 fails
//   - Error text shows the limit
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

// Example: OCPP 1.6 CiString50.
//
// Context
//   - Printable ASCII (32..126)
//   - Max length 50, empty allowed
//   - Use for longer labels
//
// What this shows
//   - Construct a valid CiString50
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

// Example: CiString50 invalid length.
//
// What this shows
//   - Input longer than 50 fails
//   - Error includes the max bound
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

// Example: OCPP 1.6 CiString255.
//
// Context
//   - Printable ASCII (32..126)
//   - Max length 255, empty allowed
//   - Use for descriptions/notes
//
// What this shows
//   - Construct a valid CiString255
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

// Example: CiString255 invalid length.
//
// What this shows
//   - Input longer than 255 fails
//   - Error prints actual and max
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

// Example: OCPP 1.6 CiString500.
//
// Context
//   - Printable ASCII (32..126)
//   - Max length 500, empty allowed
//   - For extended notes or URLs
//
// What this shows
//   - Construct a valid CiString500
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

// Example: CiString500 invalid length.
//
// What this shows
//   - Input longer than 500 fails
//   - Error shows limit and cause
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
