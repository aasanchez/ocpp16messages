package getLocalListVersion_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion"
)

// ExampleReq demonstrates creating a valid GetLocalListVersion.req message.
// GetLocalListVersion.req has no fields per OCPP 1.6 specification.
func ExampleReq() {
	_, err := getLocalListVersion.Req(getLocalListVersion.ReqInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("GetLocalListVersion.req created successfully")
	// Output:
	// GetLocalListVersion.req created successfully
}
