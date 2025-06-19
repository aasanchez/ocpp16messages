package authorizetypes_test

// import (
// 	"fmt"

// 	types "github.com/aasanchez/ocpp16messages/messages/authorize/types"
// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// func ExampleRequestMessageInput_valid() {
// 	input := types.RequestMessageInput{
// 		IdTag: "ABC1234567890XYZ78",
// 	}

// 	if err := input.Validate(); err != nil {
// 		fmt.Printf("Validation failed: %v\n", err)
// 		return
// 	}

// 	fmt.Println("RequestMessageInput is valid")

// 	// Output:
// 	// RequestMessageInput is valid
// }

// func ExampleRequestMessageInput_missingIdTag() {
// 	input := types.RequestMessageInput{} // empty

// 	if err := input.Validate(); err != nil {
// 		fmt.Printf("Validation failed: %v\n", err)
// 		return
// 	}

// 	// Output:
// 	// Validation failed: missing required field: idTag
// }

// func ExampleConfirmationMessageInput_valid() {
// 	expiry := "2027-04-12T14:03:04Z"
// 	parentId := "TAG-998877"

// 	input := types.ConfirmationMessageInput{
// 		IdTagInfo: types.IdTagInfoInput{
// 			Status:      "Accepted",
// 			ExpiryDate:  &expiry,
// 			ParentIdTag: &parentId,
// 		},
// 	}

// 	if err := input.Validate(); err != nil {
// 		fmt.Printf("Validation failed: %v\n", err)
// 		return
// 	}

// 	fmt.Println("ConfirmationMessageInput is valid")

// 	// Output:
// 	// ConfirmationMessageInput is valid
// }

// func ExampleIdTagInfo_conversionAndValidation() {
// 	info, err := types.IdTagInfo("Blocked")
// 	if err != nil {
// 		fmt.Printf("Conversion failed: %v\n", err)
// 		return
// 	}

// 	expiry, err := sharedtypes.DateTime("2027-04-12T10:03:04-04:00")
// 	if err != nil {
// 		fmt.Printf("Failed to parse DateTime: %v\n", err)
// 		return
// 	}

// 	parent, err := types.IdToken("PARENT-9999")
// 	if err != nil {
// 		fmt.Printf("Failed to parse IdToken: %v\n", err)
// 		return
// 	}

// 	info.ExpiryDate = &expiry
// 	info.ParentIdTag = &parent

// 	if err := info.Validate(); err != nil {
// 		fmt.Printf("Validation failed: %v\n", err)
// 		return
// 	}

// 	fmt.Println("IdTagInfo is valid and ready to use")

// 	// Output:
// 	// IdTagInfo is valid and ready to use
// }
