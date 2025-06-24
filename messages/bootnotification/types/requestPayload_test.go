package bootnotificationtypes

import (
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestRequestPayload_ValidateValidPayload(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargeBoxSerialNumber:   nil,
		ChargePointModel:        "ACME Model X",
		ChargePointSerialNumber: nil,
		ChargePointVendor:       "ACME Inc.",
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterSerialNumber:       nil,
		MeterType:               nil,
	}

	if err := payload.Validate(); err != nil {
		t.Errorf(sharedtypes.ErrExpectedNoValidation, err)
	}
}

func TestRequestPayload_ValidateFailsWhenChargePointVendorIsEmpty(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargeBoxSerialNumber:   nil,
		ChargePointModel:        "Model X",
		ChargePointSerialNumber: nil,
		ChargePointVendor:       "",
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterSerialNumber:       nil,
		MeterType:               nil,
	}

	if err := payload.Validate(); err == nil {
		t.Fatal(sharedtypes.ErrExpectedValidationFail)
	}
}

func TestRequestPayload_ValidateFailsWhenChargePointModelIsEmpty(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargeBoxSerialNumber:   nil,
		ChargePointModel:        "",
		ChargePointSerialNumber: nil,
		ChargePointVendor:       "Vendor X",
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterSerialNumber:       nil,
		MeterType:               nil,
	}

	if err := payload.Validate(); err == nil {
		t.Fatal(sharedtypes.ErrExpectedValidationFail)
	}
}
