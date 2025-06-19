package authorize_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleConfirmation() {
	expiry := "2027-04-12T10:03:04Z"
	parent := "A632-E2BB0231072C"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	msg, err := authorize.Confirmation(payload)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	value := msg.IdTagInfo.Value()

	fmt.Println("Status:", value.Status)
	fmt.Println("ExpiryDate:", *value.ExpiryDate)
	fmt.Println("ParentIdTag:", *value.ParentIdTag)

	// Output:
	// Status: Accepted
	// ExpiryDate: 2027-04-12T10:03:04Z
	// ParentIdTag: A632-E2BB0231072C
}
