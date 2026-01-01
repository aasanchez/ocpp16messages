package sharedtypes_test

//revive:disable line-length-limit
import (
	"fmt"
	"strings"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	msgError                 = "Error:"
	msgLength                = "Length:"
	validCiString500Length   = 499
	invalidCiString500Length = 501
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
	// Error: CiString Error on Construc (len=21, max=20): exceeds maximum length
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
	// Error: CiString Error on Construc (len=26, max=25): exceeds maximum length
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
	// Error: CiString Error on Construc (len=51, max=50): exceeds maximum length
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
	// Error: CiString Error on Construc (len=256, max=255): exceeds maximum length
}

func ExampleSetCiString500Type() {
	input := strings.Repeat("A", validCiString500Length)

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
	input := strings.Repeat("A", invalidCiString500Length)

	cistr, err := st.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construc (len=501, max=500): exceeds maximum length
}
