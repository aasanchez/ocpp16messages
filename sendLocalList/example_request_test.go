package sendLocalList_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/sendLocalList"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
)

// ExampleReq demonstrates creating a valid SendLocalList.req message
// with a full list replacement.
func ExampleReq() {
	req, err := sendLocalList.Req(sendLocalList.ReqInput{
		ListVersion: 1,
		LocalAuthorizationList: []slt.AuthorizationDataInput{
			{
				IdTag: "RFID12345",
				IdTagInfo: &slt.IdTagInfoInput{
					Status:      "Accepted",
					ExpiryDate:  nil,
					ParentIdTag: nil,
				},
			},
		},
		UpdateType: "Full",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("ListVersion: %d\n", req.ListVersion.Value())
	fmt.Printf("UpdateType: %s\n", req.UpdateType.String())
	fmt.Printf("Entries: %d\n", len(req.LocalAuthorizationList))
	// Output:
	// ListVersion: 1
	// UpdateType: Full
	// Entries: 1
}

// ExampleReq_differential demonstrates creating a SendLocalList.req message
// with a differential update.
func ExampleReq_differential() {
	req, err := sendLocalList.Req(sendLocalList.ReqInput{
		ListVersion: 2,
		LocalAuthorizationList: []slt.AuthorizationDataInput{
			{
				IdTag:     "RFID99999",
				IdTagInfo: nil,
			},
		},
		UpdateType: "Differential",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("ListVersion: %d\n", req.ListVersion.Value())
	fmt.Printf("UpdateType: %s\n", req.UpdateType.String())
	// Output:
	// ListVersion: 2
	// UpdateType: Differential
}

// ExampleReq_clearList demonstrates clearing the local authorization list
// by sending an empty list with Full update type.
func ExampleReq_clearList() {
	req, err := sendLocalList.Req(sendLocalList.ReqInput{
		ListVersion:            3,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf("ListVersion: %d\n", req.ListVersion.Value())
	fmt.Printf("UpdateType: %s\n", req.UpdateType.String())
	fmt.Printf("Entries: %d\n", len(req.LocalAuthorizationList))
	// Output:
	// ListVersion: 3
	// UpdateType: Full
	// Entries: 0
}

// ExampleReq_invalidUpdateType demonstrates the error returned when
// an invalid update type is provided.
func ExampleReq_invalidUpdateType() {
	_, err := sendLocalList.Req(sendLocalList.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: nil,
		UpdateType:             "Invalid",
	})
	if err != nil {
		fmt.Println("Error: invalid update type")
	}
	// Output:
	// Error: invalid update type
}
