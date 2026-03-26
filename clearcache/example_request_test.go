package clearcache_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/clearcache"
)

// ExampleReq demonstrates creating a valid ClearCache.req message.
// ClearCache.req has no fields per OCPP 1.6 specification.
func ExampleReq() {
	_, err := clearcache.Req(clearcache.ReqInput{})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ClearCache.req created successfully")
	// Output:
	// ClearCache.req created successfully
}
