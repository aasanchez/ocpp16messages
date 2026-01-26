package types_test

import (
	"fmt"
	"strings"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	ciString20Len  = 20
	ciString20Over = 21
)

func ExampleNewCiString20Type() {
	input := strings.Repeat("A", ciString20Len)

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
	input := strings.Repeat("B", ciString20Over)

	_, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// invalid value: exceeds maximum length (len=21, max=20)
}
