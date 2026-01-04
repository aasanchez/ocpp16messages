package types_test

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

func ExampleNewCiString20Type() {
	input := strings.Repeat("A", 20)

	cistr, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))

	// Output:
	// Length: 20
}

func ExampleNewCiString20Type_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construct (len=21, max=20): exceeds maximum length
}

func ExampleNewCiString25Type() {
	input := strings.Repeat("A", 25)

	cistr, err := st.NewCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 25
}

func ExampleNewCiString25Type_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := st.NewCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construct (len=26, max=25): exceeds maximum length
}

func ExampleNewCiString50Type() {
	input := strings.Repeat("A", 50)

	cistr, err := st.NewCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 50
}

func ExampleNewCiString50Type_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := st.NewCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construct (len=51, max=50): exceeds maximum length
}

func ExampleNewCiString255Type() {
	input := strings.Repeat("A", 255)

	cistr, err := st.NewCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 255
}

func ExampleNewCiString255Type_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := st.NewCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construct (len=256, max=255): exceeds maximum length
}

func ExampleNewCiString500Type() {
	input := strings.Repeat("A", validCiString500Length)

	cistr, err := st.NewCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 499
}

func ExampleNewCiString500Type_invalid() {
	input := strings.Repeat("A", invalidCiString500Length)

	cistr, err := st.NewCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construct (len=501, max=500): exceeds maximum length
}
