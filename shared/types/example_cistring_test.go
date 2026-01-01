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

func ExampleNewCiString20() {
	input := strings.Repeat("A", 20)

	cistr, err := st.NewCiString20(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))

	// Output:
	// Length: 20
}

func ExampleNewCiString20_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := st.NewCiString20(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construc (len=21, max=20): exceeds maximum length
}

func ExampleNewCiString25() {
	input := strings.Repeat("A", 25)

	cistr, err := st.NewCiString25(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 25
}

func ExampleNewCiString25_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := st.NewCiString25(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construc (len=26, max=25): exceeds maximum length
}

func ExampleNewCiString50() {
	input := strings.Repeat("A", 50)

	cistr, err := st.NewCiString50(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 50
}

func ExampleNewCiString50_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := st.NewCiString50(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construc (len=51, max=50): exceeds maximum length
}

func ExampleNewCiString255() {
	input := strings.Repeat("A", 255)

	cistr, err := st.NewCiString255(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 255
}

func ExampleNewCiString255_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := st.NewCiString255(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construc (len=256, max=255): exceeds maximum length
}

func ExampleNewCiString500() {
	input := strings.Repeat("A", validCiString500Length)

	cistr, err := st.NewCiString500(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 499
}

func ExampleNewCiString500_invalid() {
	input := strings.Repeat("A", invalidCiString500Length)

	cistr, err := st.NewCiString500(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: CiString Error on Construc (len=501, max=500): exceeds maximum length
}
