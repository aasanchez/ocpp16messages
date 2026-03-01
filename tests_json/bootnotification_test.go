package tests_json_test

import (
	"testing"

	bn "github.com/aasanchez/ocpp16messages/bootNotification"
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
		t.Fatalf("bootNotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestBootNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := bn.Conf(bn.ConfInput{
		Status:      "Accepted",
		CurrentTime: "2025-01-15T12:00:00Z",
		Interval:    300,
	})
	if err != nil {
		t.Fatalf("bootNotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

