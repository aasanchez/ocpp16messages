package getlocallistversion_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getlocallistversion"
)

// ExampleReq demonstrates creating a valid GetLocalListVersion.req message.
// GetLocalListVersion.req has no fields per OCPP 1.6 specification.
func ExampleReq() {
	_, err := getlocallistversion.Req(getlocallistversion.ReqInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("GetLocalListVersion.req created successfully")
	// Output:
	// GetLocalListVersion.req created successfully
}
