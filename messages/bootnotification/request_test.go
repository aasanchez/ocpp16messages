package bootnotification

import (
	"errors"
	"strings"
	"testing"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func strPtr(s string) *string {
	return &s
}

func TestRequest_validPayload(t *testing.T) {
	t.Parallel()

	input := bootnotificationtypes.RequestPayload{
		ChargeBoxSerialNumber:   strPtr("CBOX-001"),
		ChargePointModel:        "ModelX",
		ChargePointSerialNumber: strPtr("CPSN-001"),
		ChargePointVendor:       "VendorX",
		FirmwareVersion:         strPtr("FW-1.0.0"),
		Iccid:                   strPtr("8944501234567890000"),
		Imsi:                    strPtr("262017123456789"),
		MeterSerialNumber:       strPtr("MTR-001"),
		MeterType:               strPtr("SmartMeter"),
	}

	_, err := Request(input)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestRequest_missingChargePointModel(t *testing.T) {
	t.Parallel()

	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "",
		ChargePointVendor: "VendorX",
	}

	_, err := Request(input)
	if err == nil || !errors.Is(err, ErrEmptyValueNotAllowed) {
		t.Fatalf("expected ErrEmptyValueNotAllowed for ChargePointModel, got: %v", err)
	}
}

func TestRequest_invalidChargePointModel(t *testing.T) {
	t.Parallel()

	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  strings.Repeat("A", 21),
		ChargePointVendor: "VendorX",
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid ChargePointModel, got nil")
	}
}

func TestRequest_missingChargePointVendor(t *testing.T) {
	t.Parallel()

	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "",
	}

	_, err := Request(input)
	if err == nil || !errors.Is(err, ErrEmptyValueNotAllowed) {
		t.Fatalf("expected ErrEmptyValueNotAllowed for ChargePointVendor, got: %v", err)
	}
}

func TestRequest_invalidChargePointVendor(t *testing.T) {
	t.Parallel()

	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: strings.Repeat("V", 21),
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid ChargePointVendor, got nil")
	}
}

func TestRequest_invalidOptionalField_ChargeBoxSerialNumber(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("X", 26)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:        "ModelX",
		ChargePointVendor:       "VendorX",
		ChargeBoxSerialNumber:   &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid ChargeBoxSerialNumber, got nil")
	}
}

func TestRequest_invalidOptionalField_ChargePointSerialNumber(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("S", 26)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:        "ModelX",
		ChargePointVendor:       "VendorX",
		ChargePointSerialNumber: &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid ChargePointSerialNumber, got nil")
	}
}

func TestRequest_invalidOptionalField_FirmwareVersion(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("F", 51)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorX",
		FirmwareVersion:   &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid FirmwareVersion, got nil")
	}
}

func TestRequest_invalidOptionalField_Iccid(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("1", 21)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorX",
		Iccid:             &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid Iccid, got nil")
	}
}

func TestRequest_invalidOptionalField_Imsi(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("2", 21)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorX",
		Imsi:              &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid Imsi, got nil")
	}
}

func TestRequest_invalidOptionalField_MeterSerialNumber(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("M", 26)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:    "ModelX",
		ChargePointVendor:   "VendorX",
		MeterSerialNumber:   &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid MeterSerialNumber, got nil")
	}
}

func TestRequest_invalidOptionalField_MeterType(t *testing.T) {
	t.Parallel()

	invalid := strings.Repeat("T", 26)
	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorX",
		MeterType:         &invalid,
	}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid MeterType, got nil")
	}
}

func TestRequest_allOptionalFieldsNil(t *testing.T) {
	t.Parallel()

	input := bootnotificationtypes.RequestPayload{
		ChargePointModel:  "ModelX",
		ChargePointVendor: "VendorX",
	}

	_, err := Request(input)
	if err != nil {
		t.Fatalf("expected no error with all optional fields nil, got: %v", err)
	}
}
