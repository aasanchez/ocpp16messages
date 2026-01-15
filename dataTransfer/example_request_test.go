package dataTransfer_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/dataTransfer"
)

// ExampleReq demonstrates creating a valid DataTransfer.req message
// with only the required vendorId field.
func ExampleReq() {
	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "com.example.vendor",
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("VendorId:", req.VendorId.Value())
	// Output:
	// VendorId: com.example.vendor
}

// ExampleReq_withMessageId demonstrates creating a DataTransfer.req message
// with an optional messageId.
func ExampleReq_withMessageId() {
	messageId := "CustomMessage"

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "com.example.vendor",
		MessageId: &messageId,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("VendorId:", req.VendorId.Value())
	fmt.Println("MessageId:", req.MessageId.Value())
	// Output:
	// VendorId: com.example.vendor
	// MessageId: CustomMessage
}

// ExampleReq_withData demonstrates creating a DataTransfer.req message
// with an optional data payload.
func ExampleReq_withData() {
	data := `{"temperature": 25.5, "humidity": 60}`

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "com.example.vendor",
		MessageId: nil,
		Data:      &data,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("VendorId:", req.VendorId.Value())
	fmt.Println("HasData:", req.Data != nil)
	// Output:
	// VendorId: com.example.vendor
	// HasData: true
}

// ExampleReq_complete demonstrates creating a complete DataTransfer.req
// message with all fields populated.
func ExampleReq_complete() {
	messageId := "GetTemperature"
	data := `{"sensorId": "temp-001"}`

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "com.example.vendor",
		MessageId: &messageId,
		Data:      &data,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("VendorId:", req.VendorId.Value())
	fmt.Println("MessageId:", req.MessageId.Value())
	fmt.Println("Data:", *req.Data)
	// Output:
	// VendorId: com.example.vendor
	// MessageId: GetTemperature
	// Data: {"sensorId": "temp-001"}
}

// ExampleReq_emptyVendorId demonstrates the error returned when
// an empty vendorId is provided.
func ExampleReq_emptyVendorId() {
	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "",
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// vendorId: value cannot be empty
}

// ExampleReq_invalidVendorIdChars demonstrates the error returned when
// the vendorId contains non-printable ASCII characters.
func ExampleReq_invalidVendorIdChars() {
	_, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "vendor\x00id",
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// vendorId: CiString: value contains non-printable ASCII characters
}
