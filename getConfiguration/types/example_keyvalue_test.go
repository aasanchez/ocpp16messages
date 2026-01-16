package types_test

import (
	"fmt"

	gt "github.com/aasanchez/ocpp16messages/getConfiguration/types"
)

const labelKey = "Key:"

// ExampleNewKeyValue demonstrates creating a valid KeyValue with all fields.
func ExampleNewKeyValue() {
	value := "300"

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelKey, keyValue.Key().Value())
	fmt.Println("Readonly:", keyValue.Readonly())
	fmt.Println("Value:", keyValue.Value().Value())
	// Output:
	// Key: HeartbeatInterval
	// Readonly: false
	// Value: 300
}

// ExampleNewKeyValue_readonlyKey demonstrates creating a read-only KeyValue.
func ExampleNewKeyValue_readonlyKey() {
	value := "Core,FirmwareManagement"

	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "SupportedFeatureProfiles",
		Readonly: true,
		Value:    &value,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelKey, keyValue.Key().Value())
	fmt.Println("Readonly:", keyValue.Readonly())
	// Output:
	// Key: SupportedFeatureProfiles
	// Readonly: true
}

// ExampleNewKeyValue_noValue demonstrates creating a KeyValue without a value
// (for keys that are known but not set).
func ExampleNewKeyValue_noValue() {
	keyValue, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "AuthorizationKey",
		Readonly: false,
		Value:    nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelKey, keyValue.Key().Value())
	fmt.Println("HasValue:", keyValue.Value() != nil)
	// Output:
	// Key: AuthorizationKey
	// HasValue: false
}

// ExampleNewKeyValue_emptyKey demonstrates the error returned when
// an empty key is provided.
func ExampleNewKeyValue_emptyKey() {
	value := "test"

	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "",
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// key: value cannot be empty
}

// ExampleNewKeyValue_invalidKeyChars demonstrates the error returned when
// the key contains non-printable ASCII characters.
func ExampleNewKeyValue_invalidKeyChars() {
	_, err := gt.NewKeyValue(gt.KeyValueInput{
		Key:      "Invalid\x00Key",
		Readonly: false,
		Value:    nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// key: CiString: value contains non-printable ASCII characters
}
