package types_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/aasanchez/ocpp16messages/types"
)

func ExampleAuthorizationStatus() {
	status := types.AuthorizationStatus("Expired")

	if status.IsValid() {
		fmt.Printf("'%s' is a valid status.\n", status)
	} else {
		fmt.Printf("'%s' is not a valid status.\n", status)
	}

	// Output:
	// 'Expired' is a valid status.
}

func ExampleAuthorizationStatus_inValid() {
	status := types.AuthorizationStatus("InProgress")

	if status.IsValid() {
		fmt.Printf("'%s' is a valid status.\n", status)
	} else {
		fmt.Printf("'%s' is not a valid status.\n", status)
	}

	// Output:
	// 'InProgress' is not a valid status.
}

func ExampleCiString20() {
	input := strings.Repeat("A", 20)

	cistr, err := types.CiString20(input)
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

	cistr, err := types.CiString20(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: value exceeds maximum allowed length: actual length 21, max 20
}

func ExampleIdTokenType() {
	input := "ABC1234567890XYZ7890" // valid 20-char ASCII input

	idToken, err := types.IdToken(input)
	if err != nil {
		log.Fatalf("Failed to create IdToken: %v", err)

		return
	}

	fmt.Printf("Valid IdToken: %s\n", idToken.String())
	// Output:
	// Valid IdToken: ABC1234567890XYZ7890
}

func ExampleCiString() {
	input := "Hello, OCPP!"
	maxLen := 20

	// Attempt to create a valid ciString
	cs, err := types.CiString(input, maxLen)
	if err != nil {
		log.Fatalf("Failed to create ciString: %v", err)
	}

	fmt.Println("ciString created:", cs.String())

	// Output:
	// ciString created: Hello, OCPP!
}

func ExampleCiString500() {
	input := strings.Repeat("A", 499)

	cistr, err := types.CiString500(input)
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

	cistr, err := types.CiString500(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: value exceeds maximum allowed length: actual length 501, max 500
}
