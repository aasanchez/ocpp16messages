package bootnotification

import (
	"strings"
	"testing"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func validInput() bootnotificationtypes.RequestPayload {
	return bootnotificationtypes.RequestPayload{
		ChargeBoxSerialNumber:   "CBOX-001",
		ChargePointModel:        "ModelX",
		ChargePointSerialNumber: "CPSN-001",
		ChargePointVendor:       "VendorX",
		FirmwareVersion:         "FW-1.0.0",
		Iccid:                   "8944501234567890000",
		Imsi:                    "262017123456789",
		MeterSerialNumber:       "MTR-001",
		MeterType:               "SmartMeter",
	}
}

func TestRequest_validPayload(t *testing.T) {
	t.Parallel()

	_, err := Request(validInput())
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}
}

func Test_requiredCiString20_MissingAndTooLong(t *testing.T) {
	t.Parallel()

	_, err := requiredCiString20("Field", "")
	if err == nil || !strings.Contains(err.Error(), "Field") {
		t.Error("expected error for empty string")
	}

	_, err = requiredCiString20("Field", strings.Repeat("A", 21))
	if err == nil || !strings.Contains(err.Error(), "Field") {
		t.Error("expected error for too long string")
	}
}

func Test_requiredCiString20_valid(t *testing.T) {
	t.Parallel()

	val, err := requiredCiString20("Field", "Valid")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if val.Value() != "Valid" {
		t.Errorf("unexpected value: got %q", val.Value())
	}
}

func Test_wrapErr_IncludesField(t *testing.T) {
	t.Parallel()

	err := wrapErr("Field", sharedtypes.ErrEmptyValueNotAllowed)
	if err == nil || !strings.Contains(err.Error(), "Field") {
		t.Error("expected wrapped error to include field name")
	}
}

func TestRequest_requiredFieldsValidation_missingModel(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointModel = ""

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for missing ChargePointModel, got nil")
	}
}

func TestRequest_requiredFieldsValidation_missingVendor(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointVendor = ""

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for missing ChargePointVendor, got nil")
	}
}

func TestRequest_requiredFieldsValidation_invalidModel(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointModel = strings.Repeat("A", 21)

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for too long ChargePointModel, got nil")
	}
}

func TestRequest_requiredFieldsValidation_invalidVendor(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointVendor = strings.Repeat("V", 21)

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for too long ChargePointVendor, got nil")
	}
}

func TestRequest_requiredFieldsValidation_nonPrintableModel(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointModel = "Model\x01"

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for non-printable ChargePointModel, got nil")
	}
}

func TestRequest_requiredFieldsValidation_nonPrintableVendor(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointVendor = "Vend\x01"

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for non-printable ChargePointVendor, got nil")
	}
}

func TestRequest_invalidChargeBoxSerialNumber(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargeBoxSerialNumber = strings.Repeat("X", 26)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "chargeBoxSerialNumber") {
		t.Fatalf("expected error for invalid ChargeBoxSerialNumber, got: %v", err)
	}
}

func TestRequest_invalidChargePointSerialNumber(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.ChargePointSerialNumber = strings.Repeat("P", 26)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "chargePointSerialNumber") {
		t.Fatalf("expected error for invalid ChargePointSerialNumber, got: %v", err)
	}
}

func TestRequest_invalidFirmwareVersion(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.FirmwareVersion = strings.Repeat("F", 51)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "firmwareVersion") {
		t.Fatalf("expected error for invalid FirmwareVersion, got: %v", err)
	}
}

func TestRequest_invalidIccid(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.Iccid = strings.Repeat("9", 21)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "iccid") {
		t.Fatalf("expected error for invalid Iccid, got: %v", err)
	}
}

func TestRequest_invalidImsi(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.Imsi = strings.Repeat("2", 21)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "imsi") {
		t.Fatalf("expected error for invalid Imsi, got: %v", err)
	}
}

func TestRequest_invalidMeterSerialNumber(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.MeterSerialNumber = strings.Repeat("M", 26)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "meterSerialNumber") {
		t.Fatalf("expected error for invalid MeterSerialNumber, got: %v", err)
	}
}

func TestRequest_invalidMeterType(t *testing.T) {
	t.Parallel()

	input := validInput()
	input.MeterType = strings.Repeat("T", 26)

	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "meterType") {
		t.Fatalf("expected error for invalid MeterType, got: %v", err)
	}
}
