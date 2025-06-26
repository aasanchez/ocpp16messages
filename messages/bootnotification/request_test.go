package bootnotification

import (
	"errors"
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
	err := wrapErr("Field", errors.New("msg"))
	if err == nil || !strings.Contains(err.Error(), "Field") {
		t.Errorf("expected wrapped error to include field name")
	}
}

func TestRequest_requiredFieldsValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input bootnotificationtypes.RequestPayload
	}{
		{
			"missing model",
			bootnotificationtypes.RequestPayload{
				ChargePointVendor: "VendorX",
			},
		},
		{
			"missing vendor",
			bootnotificationtypes.RequestPayload{
				ChargePointModel: "ModelX",
			},
		},
		{
			"too long model",
			bootnotificationtypes.RequestPayload{
				ChargePointModel:  strings.Repeat("A", 21),
				ChargePointVendor: "VendorX",
			},
		},
		{
			"too long vendor",
			bootnotificationtypes.RequestPayload{
				ChargePointModel:  "ModelX",
				ChargePointVendor: strings.Repeat("V", 21),
			},
		},
		{
			"non-printable model",
			bootnotificationtypes.RequestPayload{
				ChargePointModel:  "Model\x01",
				ChargePointVendor: "VendorX",
			},
		},
		{
			"non-printable vendor",
			bootnotificationtypes.RequestPayload{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "Vend\x01",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Request(tt.input)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}

func TestRequest_requiredCiString20_TriggersErrorForModel(t *testing.T) {
	t.Parallel()
	input := validInput()
	input.ChargePointModel = strings.Repeat("M", 21)
	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "ChargePointModel") {
		t.Fatal("expected error on ChargePointModel, got nil or wrong message")
	}
}

func TestRequest_requiredCiString20_TriggersErrorForVendor(t *testing.T) {
	t.Parallel()
	input := validInput()
	input.ChargePointVendor = strings.Repeat("V", 21)
	_, err := Request(input)
	if err == nil || !strings.Contains(err.Error(), "ChargePointVendor") {
		t.Fatal("expected error on ChargePointVendor, got nil or wrong message")
	}
}

func TestRequest_optionalFieldsTooLong(t *testing.T) {
	t.Parallel()

	base := validInput()
	fields := map[string]string{
		"ChargeBoxSerialNumber":   strings.Repeat("X", 26),
		"ChargePointSerialNumber": strings.Repeat("Y", 26),
		"FirmwareVersion":         strings.Repeat("Z", 51),
		"Iccid":                   strings.Repeat("1", 21),
		"Imsi":                    strings.Repeat("2", 21),
		"MeterSerialNumber":       strings.Repeat("M", 26),
		"MeterType":               strings.Repeat("T", 26),
	}

	for field, value := range fields {
		t.Run(field, func(t *testing.T) {
			input := base
			switch field {
			case "ChargeBoxSerialNumber":
				input.ChargeBoxSerialNumber = value
			case "ChargePointSerialNumber":
				input.ChargePointSerialNumber = value
			case "FirmwareVersion":
				input.FirmwareVersion = value
			case "Iccid":
				input.Iccid = value
			case "Imsi":
				input.Imsi = value
			case "MeterSerialNumber":
				input.MeterSerialNumber = value
			case "MeterType":
				input.MeterType = value
			}

			_, err := Request(input)
			if err == nil {
				t.Fatalf("expected error for field %s, got nil", field)
			}
		})
	}
}

func TestRequest_eachFieldTriggersValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		field  string
		mutate func(bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload
	}{
		{
			"ChargeBoxSerialNumber",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.ChargeBoxSerialNumber = "\x01"
				return p
			},
		},
		{
			"ChargePointSerialNumber",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.ChargePointSerialNumber = "\x01"
				return p
			},
		},
		{
			"FirmwareVersion",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.FirmwareVersion = "\x01"
				return p
			},
		},
		{
			"Iccid",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.Iccid = "\x01"
				return p
			},
		},
		{
			"Imsi",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.Imsi = "\x01"
				return p
			},
		},
		{
			"MeterSerialNumber",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.MeterSerialNumber = "\x01"
				return p
			},
		},
		{
			"MeterType",
			func(p bootnotificationtypes.RequestPayload) bootnotificationtypes.RequestPayload {
				p.MeterType = "\x01"
				return p
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.field+"_nonPrintable", func(t *testing.T) {
			input := tt.mutate(validInput())
			_, err := Request(input)
			if err == nil {
				t.Errorf("expected error for %s, got nil", tt.field)
			}
		})
	}
}
