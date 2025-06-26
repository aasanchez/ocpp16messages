package bootnotification_test

import (
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/bootnotification"
	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func BenchmarkRequest_validPayload(b *testing.B) {
	input := bootnotificationtypes.RequestPayload{
		ChargeBoxSerialNumber:   strings.Repeat("A", 25),
		ChargePointModel:        strings.Repeat("M", 20),
		ChargePointSerialNumber: strings.Repeat("S", 25),
		ChargePointVendor:       strings.Repeat("V", 20),
		FirmwareVersion:         strings.Repeat("F", 50),
		Iccid:                   strings.Repeat("I", 20),
		Imsi:                    strings.Repeat("Z", 20),
		MeterSerialNumber:       strings.Repeat("X", 25),
		MeterType:               strings.Repeat("T", 25),
	}

	b.ReportAllocs()

	for range b.N {
		_, err := bootnotification.Request(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
