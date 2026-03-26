package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/statusnotification"
)

func TestStatusNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     1,
		ErrorCode:       "NoError",
		Status:          "Available",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Fatalf("statusnotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestStatusNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := statusnotification.Conf(statusnotification.ConfInput{})
	if err != nil {
		t.Fatalf("statusnotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
