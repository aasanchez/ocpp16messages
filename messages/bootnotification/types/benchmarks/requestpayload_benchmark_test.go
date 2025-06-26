package bootnotification_test

import (
	"strings"
	"testing"

	types "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func BenchmarkRequestPayload_Validate(b *testing.B) {
	input := types.RequestPayload{
		ChargeBoxSerialNumber:   strings.Repeat("A", 25),
		ChargePointModel:        "Model-X",
		ChargePointSerialNumber: strings.Repeat("S", 25),
		ChargePointVendor:       "VendorY",
		FirmwareVersion:         strings.Repeat("F", 50),
		Iccid:                   strings.Repeat("I", 20),
		Imsi:                    strings.Repeat("Z", 20),
		MeterSerialNumber:       strings.Repeat("X", 25),
		MeterType:               strings.Repeat("T", 25),
	}

	b.ReportAllocs()
	for range b.N {
		if err := input.Validate(); err != nil {
			b.Fatalf("unexpected validation error: %v", err)
		}
	}
}
