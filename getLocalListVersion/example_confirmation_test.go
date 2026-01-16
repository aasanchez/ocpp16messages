package getLocalListVersion_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion"
)

// ExampleConf demonstrates creating a valid GetLocalListVersion.conf message
// with a positive version number.
func ExampleConf() {
	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: 5,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ListVersion:", conf.ListVersion.Value())
	// Output:
	// ListVersion: 5
}

// ExampleConf_unsupported demonstrates creating a GetLocalListVersion.conf
// message indicating the Charge Point does not support Local Authorization
// Lists.
func ExampleConf_unsupported() {
	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: -1,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IsUnsupported:", conf.ListVersion.IsUnsupported())
	// Output:
	// IsUnsupported: true
}

// ExampleConf_emptyList demonstrates creating a GetLocalListVersion.conf
// message indicating the local authorization list is empty.
func ExampleConf_emptyList() {
	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: 0,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IsEmpty:", conf.ListVersion.IsEmpty())
	// Output:
	// IsEmpty: true
}
