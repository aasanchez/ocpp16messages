package types_test

//revive:disable line-length-limit
import (
	"fmt"
	"strings"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func ExampleNewCiString20Type() {
	input := strings.Repeat("A", 20)

	cistr, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(cistr.Value())

	// Output:
	// AAAAAAAAAAAAAAAAAAAA
}

func ExampleNewCiString20Type_invalid() {
	input := strings.Repeat("B", 21)

	_, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// CiString Error on Construct (len=21, max=20): exceeds maximum length
}
