package changeConfiguration_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeConfiguration"
)

const (
	keyLabel   = "Key:"
	valueLabel = "Value:"
)

// ExampleReq demonstrates creating a valid ChangeConfiguration.req message
// to change the HeartbeatInterval configuration.
func ExampleReq() {
	req, err := changeConfiguration.Req(changeConfiguration.ReqInput{
		Key:   "HeartbeatInterval",
		Value: "900",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(keyLabel, req.Key.Value())
	fmt.Println(valueLabel, req.Value.Value())
	// Output:
	// Key: HeartbeatInterval
	// Value: 900
}

// ExampleReq_emptyValue demonstrates the error returned when
// an empty value is provided.
func ExampleReq_emptyValue() {
	_, err := changeConfiguration.Req(changeConfiguration.ReqInput{
		Key:   "AuthorizationKey",
		Value: "",
	})
	if err != nil {
		fmt.Println("Error: invalid value")
	}
	// Output:
	// Error: invalid value
}

// ExampleReq_emptyKey demonstrates the error returned when
// an empty key is provided.
func ExampleReq_emptyKey() {
	_, err := changeConfiguration.Req(changeConfiguration.ReqInput{
		Key:   "",
		Value: "900",
	})
	if err != nil {
		fmt.Println("Error: invalid key")
	}
	// Output:
	// Error: invalid key
}

// ExampleReq_keyTooLong demonstrates the error returned when
// the key exceeds 50 characters.
func ExampleReq_keyTooLong() {
	_, err := changeConfiguration.Req(changeConfiguration.ReqInput{
		Key:   "ThisKeyIsWayTooLongAndExceedsTheFiftyCharacterLimit",
		Value: "900",
	})
	if err != nil {
		fmt.Println("Error: key too long")
	}
	// Output:
	// Error: key too long
}
