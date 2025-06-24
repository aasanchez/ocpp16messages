package bootnotificationtypes

import (
	"testing"
)

const (
	errExpectedValidationFail = "expected validation error, got nil"
	errExpectedNoValidation   = "expected no validation error, got: %v"
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
		t.Errorf(errExpectedNoValidation, err)
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
		t.Fatal(errExpectedValidationFail)
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
		t.Fatal(errExpectedValidationFail)
	}
}
