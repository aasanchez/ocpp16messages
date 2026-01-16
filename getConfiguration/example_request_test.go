package getConfiguration_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getConfiguration"
)

// ExampleReq demonstrates creating a valid GetConfiguration.req message
// without specifying any keys (requests all configuration settings).
func ExampleReq() {
	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("HasKeys:", req.Key != nil)
	// Output:
	// HasKeys: false
}

// ExampleReq_withEmptyKeys demonstrates creating a GetConfiguration.req
// message with an empty key list (equivalent to requesting all settings).
func ExampleReq_withEmptyKeys() {
	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("HasKeys:", req.Key != nil)
	// Output:
	// HasKeys: false
}

// ExampleReq_withSingleKey demonstrates creating a GetConfiguration.req
// message requesting a single configuration key.
func ExampleReq_withSingleKey() {
	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"HeartbeatInterval"},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("KeyCount:", len(req.Key))
	fmt.Println("Key[0]:", req.Key[0].Value())
	// Output:
	// KeyCount: 1
	// Key[0]: HeartbeatInterval
}

// ExampleReq_withMultipleKeys demonstrates creating a GetConfiguration.req
// message requesting multiple configuration keys.
func ExampleReq_withMultipleKeys() {
	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{
			"HeartbeatInterval",
			"ConnectionTimeOut",
			"MeterValueSampleInterval",
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("KeyCount:", len(req.Key))
	// Output:
	// KeyCount: 3
}

// ExampleReq_emptyKey demonstrates the error returned when
// a key in the list is empty.
func ExampleReq_emptyKey() {
	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"HeartbeatInterval", ""},
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// key[1]: value cannot be empty
}

// ExampleReq_invalidKeyChars demonstrates the error returned when
// a key contains non-printable ASCII characters.
func ExampleReq_invalidKeyChars() {
	_, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"Invalid\x00Key"},
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// key[0]: CiString: value contains non-printable ASCII characters
}
