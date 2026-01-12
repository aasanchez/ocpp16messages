package authorize_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
)

// ExampleConf demonstrates creating a valid Authorize.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := authorize.Conf(authorize.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.IdTagInfo.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_blocked demonstrates creating an Authorize.conf message
// with a Blocked status.
func ExampleConf_blocked() {
	conf, err := authorize.Conf(authorize.ConfInput{
		Status: "Blocked",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.IdTagInfo.Status.String())
	// Output:
	// Status: Blocked
}

// ExampleConf_withExpiryDate demonstrates creating an Authorize.conf message
// with an expiry date.
func ExampleConf_withExpiryDate() {
	expiryDate := "2025-12-31T23:59:59Z"
	conf, err := authorize.Conf(authorize.ConfInput{
		Status:     "Accepted",
		ExpiryDate: &expiryDate,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.IdTagInfo.Status.String())
	fmt.Println("HasExpiryDate:", conf.IdTagInfo.ExpiryDate != nil)
	// Output:
	// Status: Accepted
	// HasExpiryDate: true
}

// ExampleConf_withParentIdTag demonstrates creating an Authorize.conf message
// with a parent ID tag.
func ExampleConf_withParentIdTag() {
	parentTag := "PARENT-123"
	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ParentIdTag: &parentTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.IdTagInfo.Status.String())
	fmt.Println("ParentIdTag:", conf.IdTagInfo.ParentIdTag.String())
	// Output:
	// Status: Accepted
	// ParentIdTag: PARENT-123
}

// ExampleConf_complete demonstrates creating a complete Authorize.conf message
// with all optional fields populated.
func ExampleConf_complete() {
	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"
	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIdTag: &parentTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Status:", conf.IdTagInfo.Status.String())
	fmt.Println("HasExpiryDate:", conf.IdTagInfo.ExpiryDate != nil)
	fmt.Println("ParentIdTag:", conf.IdTagInfo.ParentIdTag.String())
	// Output:
	// Status: Accepted
	// HasExpiryDate: true
	// ParentIdTag: PARENT-123
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := authorize.Conf(authorize.ConfInput{
		Status: "Unknown",
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTagInfo: NewIdTagInfo: AuthorizationStatus: invalid value
}
