package stoptransaction_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/stoptransaction"
)

const (
	exampleStatusAccepted = "Accepted"
	exampleStatusLabel    = "Status:"
	exampleErrorLabel     = "Error:"
)

// ExampleConf demonstrates creating a StopTransaction.conf with no IdTagInfo.
func ExampleConf() {
	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      nil,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	if conf.IdTagInfo == nil {
		fmt.Println("IdTagInfo: nil (no authorization info)")
	}
	// Output:
	// IdTagInfo: nil (no authorization info)
}

// ExampleConf_accepted demonstrates creating an accepted response.
func ExampleConf_accepted() {
	status := exampleStatusAccepted

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	fmt.Println(exampleStatusLabel, conf.IdTagInfo.Status().String())
	// Output:
	// Status: Accepted
}

// ExampleConf_withExpiryDate demonstrates including an expiry date.
func ExampleConf_withExpiryDate() {
	status := exampleStatusAccepted
	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiryDate,
		ParentIdTag: nil,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	fmt.Println(exampleStatusLabel, conf.IdTagInfo.Status().String())
	fmt.Println("Has expiry date:", conf.IdTagInfo.ExpiryDate() != nil)
	// Output:
	// Status: Accepted
	// Has expiry date: true
}

// ExampleConf_withParentIdTag demonstrates including a parent ID tag.
func ExampleConf_withParentIdTag() {
	status := exampleStatusAccepted
	parentIdTag := "PARENT-123"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: &parentIdTag,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	fmt.Println(exampleStatusLabel, conf.IdTagInfo.Status().String())
	fmt.Println("ParentIdTag:", conf.IdTagInfo.ParentIdTag().String())
	// Output:
	// Status: Accepted
	// ParentIdTag: PARENT-123
}

// ExampleConf_invalidStatus demonstrates validation error for invalid status.
func ExampleConf_invalidStatus() {
	status := "InvalidStatus"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		fmt.Println("Validation failed as expected")
	}
	// Output:
	// Validation failed as expected
}
