package authorizetypes_test

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func ExampleIdTagInfo_statusOnly() {
	payload := mat.IdTagInfoPayload{
		Status:      mat.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	info, err := mat.IdTagInfo(payload)
	if err != nil {
		fmt.Println("Failed statusOnly:", err)

		return
	}

	val := info.Value()
	fmt.Printf("Status: %s\n", val.Status)

	// Output:
	// Status: Accepted
}

func ExampleIdTagInfo_withExpiryDate() {
	exp := "2027-04-12T10:03:04Z"

	payload := mat.IdTagInfoPayload{
		Status:      mat.Accepted,
		ExpiryDate:  &exp,
		ParentIdTag: nil,
	}

	info, err := mat.IdTagInfo(payload)
	if err != nil {
		fmt.Println("Failed withExpiryDate:", err)

		return
	}

	val := info.Value()
	fmt.Printf("Status: %s, ExpiryDate: %s\n", val.Status, *val.ExpiryDate)

	// Output:
	// Status: Accepted, ExpiryDate: 2027-04-12T10:03:04Z
}

func ExampleIdTagInfo_withParentIdTag() {
	parent := "ABC123456"

	payload := mat.IdTagInfoPayload{
		Status:      mat.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &parent,
	}

	info, err := mat.IdTagInfo(payload)
	if err != nil {
		fmt.Println("Failed withParentIdTag:", err)

		return
	}

	val := info.Value()
	fmt.Printf("Status: %s, ParentIdTag: %s\n", val.Status, *val.ParentIdTag)

	// Output:
	// Status: Accepted, ParentIdTag: ABC123456
}

func ExampleIdTagInfo_withAllFields() {
	exp := "2026-04-12T10:03:04Z"
	parent := "XYZ789"

	payload := mat.IdTagInfoPayload{
		Status:      mat.Accepted,
		ExpiryDate:  &exp,
		ParentIdTag: &parent,
	}

	info, err := mat.IdTagInfo(payload)
	if err != nil {
		fmt.Println("Failed withAllFields:", err)

		return
	}

	val := info.Value()
	fmt.Printf("Status: %s, ExpiryDate: %s, ParentIdTag: %s\n", val.Status, *val.ExpiryDate, *val.ParentIdTag)

	// Output:
	// Status: Accepted, ExpiryDate: 2026-04-12T10:03:04Z, ParentIdTag: XYZ789
}
