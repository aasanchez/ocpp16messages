package bootnotification_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/messages/bootnotification"
	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

// ExampleRequest demonstrates how to construct a BootNotification request message.
func ExampleRequest() {
	input := bootnotificationtypes.RequestPayload{
		ChargeBoxSerialNumber:   "BOX-123456",
		ChargePointModel:        "Model-X",
		ChargePointSerialNumber: "SN-987654",
		ChargePointVendor:       "VendorY",
		FirmwareVersion:         "FW-1.2.3",
		Iccid:                   "ICCId-998877",
		Imsi:                    "IMSI-776655",
		MeterSerialNumber:       "MSN-112233",
		MeterType:               "MeterType-A1",
	}

	req, err := bootnotification.Request(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("ChargePointModel:", req.ChargePointModel.Value())
	fmt.Println("ChargePointVendor:", req.ChargePointVendor.Value())

	// Output:
	// ChargePointModel: Model-X
	// ChargePointVendor: VendorY
}

// ExampleRequest_invalid demonstrates how the BootNotification request fails with missing required fields.
func ExampleRequest_invalid() {
	input := bootnotificationtypes.RequestPayload{
		ChargeBoxSerialNumber:   "BOX-123456",
		ChargePointModel:        "",
		ChargePointSerialNumber: "",
		ChargePointVendor:       "",
		FirmwareVersion:         "",
		Iccid:                   "",
		Imsi:                    "",
		MeterSerialNumber:       "",
		MeterType:               "",
	}

	_, err := bootnotification.Request(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output:
	// Error: ChargePointModel: value must not be empty
}
