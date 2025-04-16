package types_test

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aasanchez/ocpp16messages/types"
)

func ExampleAuthorizationStatus_isValid() {
	status := types.AuthorizationStatus("Expired")

	if status.IsValid() {
		fmt.Printf("'%s' is a valid status.\n", status)
	} else {
		fmt.Printf("'%s' is not a valid status.\n", status)
	}

	// Output:
	// 'Expired' is a valid status.
}

func ExampleAuthorizationStatus_invalid() {
	status := types.AuthorizationStatus("InProgress")

	if status.IsValid() {
		fmt.Printf("'%s' is a valid status.\n", status)
	} else {
		fmt.Printf("'%s' is not a valid status.\n", status)
	}

	// Output:
	// 'InProgress' is not a valid status.
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

func ExampleCiString25() {
	input := strings.Repeat("A", 25)

	cistr, err := types.CiString25(input)
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

	cistr, err := types.CiString25(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: value exceeds maximum allowed length: actual length 26, max 25
}

func ExampleCiString50() {
	input := strings.Repeat("A", 50)

	cistr, err := types.CiString50(input)
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

	cistr, err := types.CiString50(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: value exceeds maximum allowed length: actual length 51, max 50
}

func ExampleCiString255() {
	input := strings.Repeat("A", 255)

	cistr, err := types.CiString255(input)
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

	cistr, err := types.CiString255(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Length:", len(cistr.String()))
	// Output:
	// Error: value exceeds maximum allowed length: actual length 256, max 255
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

func ExampleIdTagInfoType() {
	info, err := types.IdTagInfo(types.Accepted)
	if err != nil {
		fmt.Println("Failed to create IdTagInfo:", err)

		return
	}

	location, _ := time.LoadLocation("Europe/Madrid")
	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)
	info.ExpiryDate = &expiry

	parentId, _ := types.IdToken("ABC123")
	info.ParentIdTag = &parentId

	if err := info.Validate(); err != nil {
		fmt.Println("Validation failed:", err)

		return
	}

	fmt.Println("IdTagInfo:", info.String())

	// Output:
	// IdTagInfo: {status=Accepted, expiryDate=2027-04-12T16:03:04+02:00, parentIdTag=ABC123}
}

func ExampleIdTagInfoType_onlystatus() {
	info, err := types.IdTagInfo(types.Accepted)
	if err != nil {
		fmt.Println("Failed to create IdTagInfo:", err)

		return
	}

	info.ExpiryDate = nil
	info.ParentIdTag = nil

	if err := info.Validate(); err != nil {
		fmt.Println("Validation failed:", err)

		return
	}

	fmt.Println("IdTagInfo:", info.String())

	// Output:
	// IdTagInfo: {status=Accepted}
}

func ExampleIdTagInfoType_withParentIdTag() {
	info, err := types.IdTagInfo(types.Accepted)
	if err != nil {
		fmt.Println("Failed to create IdTagInfo:", err)

		return
	}

	parentId, _ := types.IdToken("ABC123")
	info.ParentIdTag = &parentId

	info.ExpiryDate = nil

	if err := info.Validate(); err != nil {
		fmt.Println("Validation failed:", err)

		return
	}

	fmt.Println("IdTagInfo:", info.String())

	// Output:
	// IdTagInfo: {status=Accepted, parentIdTag=ABC123}
}

func ExampleIdToken() {
	input := "ABC1234567890XYZ7890" // valid 20-char ASCII input
	idToken, _ := types.IdToken(input)
	fmt.Printf("Valid IdToken: %s\n", idToken.String())

	// Output:
	// Valid IdToken: ABC1234567890XYZ7890
}

func ExampleIdTokenType_invalid() {
	input := "ABC1234567890XYZ7890123123" // invalid 26-char ASCII input

	idToken, err := types.IdToken(input)

	if err != nil {
		fmt.Printf("Failed to create IdToken: %v", err)

		return
	}

	fmt.Printf("Valid IdToken: %s\n", idToken.String())

	// Output:
	// Failed to create IdToken: value exceeds maximum allowed length: actual length 26, max 20
}
