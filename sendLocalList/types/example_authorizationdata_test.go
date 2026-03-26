package types_test

import (
	"fmt"

	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
)

// ExampleNewAuthorizationData demonstrates creating an AuthorizationData
// entry with only an idTag (for deletion in Differential updates).
func ExampleNewAuthorizationData() {
	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag:     "RFID12345",
		IdTagInfo: nil,
	})
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("IdTag:", authData.IdTag.String())
	fmt.Println("IdTagInfo:", authData.IdTagInfo)
	// Output:
	// IdTag: RFID12345
	// IdTagInfo: <nil>
}

// ExampleNewAuthorizationData_withStatus demonstrates creating an
// AuthorizationData entry with status information.
func ExampleNewAuthorizationData_withStatus() {
	authData, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag: "RFID12345",
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      "Accepted",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	})
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("IdTag:", authData.IdTag.String())
	fmt.Println("Status:", authData.IdTagInfo.Status.String())
	// Output:
	// IdTag: RFID12345
	// Status: Accepted
}

// ExampleNewAuthorizationData_invalidIdTag demonstrates the error returned
// when an invalid idTag is provided.
func ExampleNewAuthorizationData_invalidIdTag() {
	_, err := slt.NewAuthorizationData(slt.AuthorizationDataInput{
		IdTag:     "",
		IdTagInfo: nil,
	})
	if err != nil {
		fmt.Println("Error: invalid idTag")
	}
	// Output:
	// Error: invalid idTag
}
