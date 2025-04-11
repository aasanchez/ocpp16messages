package bootnotificationtypes

import (
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestRequestPayload_ValidateValidPayload(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargeBoxSerialNumber:   "",
		ChargePointModel:        "ACME Model X",
		ChargePointSerialNumber: "",
		ChargePointVendor:       "ACME Inc.",
		FirmwareVersion:         "",
		Iccid:                   "",
		Imsi:                    "",
		MeterSerialNumber:       "",
		MeterType:               "",
	}

	if err := payload.Validate(); err != nil {
		t.Errorf(sharedtypes.ErrExpectedNoValidation, err)
	}
}

func TestRequestPayload_ValidateFailsWhenChargePointVendorIsEmpty(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargeBoxSerialNumber:   "",
		ChargePointModel:        "Model X",
		ChargePointSerialNumber: "",
		ChargePointVendor:       "",
		FirmwareVersion:         "",
		Iccid:                   "",
		Imsi:                    "",
		MeterSerialNumber:       "",
		MeterType:               "",
	}

	if err := payload.Validate(); err == nil {
		t.Fatal(sharedtypes.ErrExpectedValidationFail)
	}
}

func TestRequestPayload_ValidateFailsWhenChargePointModelIsEmpty(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargeBoxSerialNumber:   "",
		ChargePointModel:        "",
		ChargePointSerialNumber: "",
		ChargePointVendor:       "Vendor X",
		FirmwareVersion:         "",
		Iccid:                   "",
		Imsi:                    "",
		MeterSerialNumber:       "",
		MeterType:               "",
	}

	if err := payload.Validate(); err == nil {
		t.Fatal(sharedtypes.ErrExpectedValidationFail)
	}
}
