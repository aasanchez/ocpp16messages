package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/statusNotification"
)

func TestStatusNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := statusNotification.Req(statusNotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "NoError",
		Status:          "Available",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Fatalf("statusNotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestStatusNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := statusNotification.Conf(statusNotification.ConfInput{})
	if err != nil {
		t.Fatalf("statusNotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

