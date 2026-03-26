package statusnotification_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/statusnotification"
)

const (
	labelConnectorId = "ConnectorId:"
	labelStatus      = "Status:"
	labelErrorCode   = "ErrorCode:"
)

// ExampleReq demonstrates creating a valid StatusNotification.req message
// with required fields only.
func ExampleReq() {
	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "NoError",
		Status:          "Available",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelConnectorId, req.ConnectorId.Value())
	fmt.Println(labelErrorCode, req.ErrorCode.String())
	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// ConnectorId: 1
	// ErrorCode: NoError
	// Status: Available
}

// ExampleReq_charging demonstrates a Charging status notification.
func ExampleReq_charging() {
	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     2,
		ErrorCode:       "NoError",
		Status:          "Charging",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelConnectorId, req.ConnectorId.Value())
	fmt.Println(labelStatus, req.Status.String())
	// Output:
	// ConnectorId: 2
	// Status: Charging
}

// ExampleReq_faulted demonstrates a Faulted status with error code.
func ExampleReq_faulted() {
	info := "Ground fault detected on connector"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "GroundFailure",
		Status:          "Faulted",
		Info:            &info,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, req.Status.String())
	fmt.Println(labelErrorCode, req.ErrorCode.String())
	fmt.Println("HasInfo:", req.Info != nil)
	// Output:
	// Status: Faulted
	// ErrorCode: GroundFailure
	// HasInfo: true
}

// ExampleReq_withAllFields demonstrates a complete StatusNotification.req.
func ExampleReq_withAllFields() {
	info := "Charging in progress"
	timestamp := "2025-01-15T10:30:00Z"
	vendorId := "VendorX"
	vendorErrorCode := "V001"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "NoError",
		Status:          "Charging",
		Info:            &info,
		Timestamp:       &timestamp,
		VendorId:        &vendorId,
		VendorErrorCode: &vendorErrorCode,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelConnectorId, req.ConnectorId.Value())
	fmt.Println(labelStatus, req.Status.String())
	fmt.Println("HasTimestamp:", req.Timestamp != nil)
	fmt.Println("HasVendorId:", req.VendorId != nil)
	// Output:
	// ConnectorId: 1
	// Status: Charging
	// HasTimestamp: true
	// HasVendorId: true
}

// ExampleReq_invalidStatus demonstrates the error for an invalid status.
func ExampleReq_invalidStatus() {
	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "NoError",
		Status:          "InvalidStatus",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// status: invalid value
}

// ExampleReq_invalidErrorCode demonstrates the error for an invalid error code.
func ExampleReq_invalidErrorCode() {
	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "InvalidCode",
		Status:          "Available",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// errorCode: invalid value
}
