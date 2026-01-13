package bootNotification_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/bootNotification"
)

// ExampleReq demonstrates creating a valid BootNotification.req message
// with only the required fields.
func ExampleReq() {
	req, err := bootNotification.Req(bootNotification.ReqInput{
		ChargePointVendor:       "VendorABC",
		ChargePointModel:        "ModelXYZ",
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Vendor:", req.ChargePointVendor.Value())
	fmt.Println("Model:", req.ChargePointModel.Value())
	// Output:
	// Vendor: VendorABC
	// Model: ModelXYZ
}

// ExampleReq_allFields demonstrates creating a BootNotification.req message
// with all optional fields populated.
func ExampleReq_allFields() {
	serial := "SN12345"
	chargeBox := "CB12345"
	firmware := "1.0.0"
	iccid := "8901234567890123456"
	imsi := "310150123456789"
	meterType := "ABB"
	meterSerial := "MS12345"

	req, err := bootNotification.Req(bootNotification.ReqInput{
		ChargePointVendor:       "VendorABC",
		ChargePointModel:        "ModelXYZ",
		ChargePointSerialNumber: &serial,
		ChargeBoxSerialNumber:   &chargeBox,
		FirmwareVersion:         &firmware,
		Iccid:                   &iccid,
		Imsi:                    &imsi,
		MeterType:               &meterType,
		MeterSerialNumber:       &meterSerial,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Vendor:", req.ChargePointVendor.Value())
	fmt.Println("Model:", req.ChargePointModel.Value())
	fmt.Println("HasSerialNumber:", req.ChargePointSerialNumber != nil)
	fmt.Println("HasFirmware:", req.FirmwareVersion != nil)
	// Output:
	// Vendor: VendorABC
	// Model: ModelXYZ
	// HasSerialNumber: true
	// HasFirmware: true
}

// ExampleReq_emptyVendor demonstrates the error returned when
// the charge point vendor is empty.
func ExampleReq_emptyVendor() {
	_, err := bootNotification.Req(bootNotification.ReqInput{
		ChargePointVendor:       "",
		ChargePointModel:        "ModelXYZ",
		ChargePointSerialNumber: nil,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err != nil {
		fmt.Println("Error: vendor is required")
	}
	// Output:
	// Error: vendor is required
}

// ExampleReq_multipleErrors demonstrates that all validation errors
// are returned at once, not just the first one encountered.
func ExampleReq_multipleErrors() {
	longSerial := "12345678901234567890123456" // 26 chars, max is 25

	_, err := bootNotification.Req(bootNotification.ReqInput{
		ChargePointVendor:       "",
		ChargePointModel:        "",
		ChargePointSerialNumber: &longSerial,
		ChargeBoxSerialNumber:   nil,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterType:               nil,
		MeterSerialNumber:       nil,
	})
	if err != nil {
		fmt.Println("Multiple errors")
	}
	// Output:
	// Multiple errors
}
