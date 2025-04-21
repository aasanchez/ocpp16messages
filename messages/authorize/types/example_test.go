package authorizetypes_test

import (
	"fmt"
	"time"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleAuthorizationStatus_isValid() {
	status := authorizetypes.AuthorizationStatus("Expired")

	if status.IsValid() {
		fmt.Printf("'%s' is a valid status.\n", status)
	} else {
		fmt.Printf("'%s' is not a valid status.\n", status)
	}

	// Output:
	// 'Expired' is a valid status.
}

func ExampleAuthorizationStatus_invalid() {
	status := authorizetypes.AuthorizationStatus("InProgress")

	if status.IsValid() {
		fmt.Printf("'%s' is a valid status.\n", status)
	} else {
		fmt.Printf("'%s' is not a valid status.\n", status)
	}

	// Output:
	// 'InProgress' is not a valid status.
}

func ExampleIdTagInfoType() {
	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
	if err != nil {
		fmt.Println("Failed to create IdTagInfo:", err)

		return
	}

	location, _ := time.LoadLocation("Europe/Madrid")
	expiry := time.Date(2027, 4, 12, 14, 3, 4, 0, time.UTC).In(location)
	info.ExpiryDate = &expiry

	parentId, _ := authorizetypes.IdToken("ABC123")
	info.ParentIdTag = &parentId

	if err := info.Validate(); err != nil {
		fmt.Println("Validation failed:", err)

		return
	}

	fmt.Println("IdTagInfo:", info.String())

	// Output:
	// IdTagInfo: {status:Accepted, expiryDate:2027-04-12T16:03:04+02:00, parentIdTag:ABC123}
}

func ExampleIdTagInfoType_onlystatus() {
	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
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
	// IdTagInfo: {status:Accepted}
}

func ExampleIdTagInfoType_withParentIdTag() {
	info, err := authorizetypes.IdTagInfo(authorizetypes.Accepted)
	if err != nil {
		fmt.Println("Failed to create IdTagInfo:", err)

		return
	}

	parentId, _ := authorizetypes.IdToken("ABC123")
	info.ParentIdTag = &parentId

	info.ExpiryDate = nil

	if err := info.Validate(); err != nil {
		fmt.Println("Validation failed:", err)

		return
	}

	fmt.Println("IdTagInfo:", info.String())

	// Output:
	// IdTagInfo: {status:Accepted, parentIdTag:ABC123}
}

func ExampleIdToken() {
	input := "ABC1234567890XYZ7890" // valid 20-char ASCII input
	idToken, _ := authorizetypes.IdToken(input)
	fmt.Printf("Valid IdToken: %s\n", idToken.String())

	// Output:
	// Valid IdToken: ABC1234567890XYZ7890
}

func ExampleIdTokenType_invalid() {
	input := "ABC1234567890XYZ7890123123" // invalid 26-char ASCII input

	idToken, err := authorizetypes.IdToken(input)

	if err != nil {
		fmt.Printf("Failed to create IdToken: %v", err)

		return
	}

	fmt.Printf("Valid IdToken: %s\n", idToken.String())

	// Output:
	// Failed to create IdToken: invalid IdToken: value exceeds maximum allowed length: actual length 26, max 20
}
