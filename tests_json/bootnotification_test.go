package testsjson_test

import (
	"testing"

	bn "github.com/aasanchez/ocpp16messages/bootnotification"
)

func TestBootNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := bn.Req(bn.ReqInput{
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
		t.Fatalf("bootnotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestBootNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := bn.Conf(bn.ConfInput{
		Status:      "Accepted",
		CurrentTime: "2025-01-15T12:00:00Z",
		Interval:    300,
	})
	if err != nil {
		t.Fatalf("bootnotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
